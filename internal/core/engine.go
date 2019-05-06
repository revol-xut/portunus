/*******************************************************************************
*
* Copyright 2019 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package core

import (
	"sync"
)

//LDAPObject describes an object that can be stored in the LDAP directory.
type LDAPObject struct {
	DN         string
	Attributes map[string][]string
}

//Engine is the core engine of portunus-server.
type Engine interface {
	FindGroup(name string) *Group
	FindUser(loginName string) *UserWithPerms
	ListGroups() []Group
	ListUsers() []User
	//The ChangeX() methods are used to create, modify and delete entities.
	//When creating a new entity, the action is invoked with a
	//default-constructed argument. To delete an entity, return nil from the
	//action. If a non-nil error is returned, it's the one returned by the
	//action.
	ChangeUser(loginName string, action func(User) (*User, error)) error
	ChangeGroup(name string, action func(Group) (*Group, error)) error
}

//engine implements the Engine interface.
type engine struct {
	FileStoreAPI *FileStoreAPI
	LDAPUpdates  chan<- []LDAPObject
	Users        map[string]User
	Groups       map[string]Group
	Mutex        *sync.RWMutex
	LDAPSuffix   string
}

//RunEngineAsync runs the main engine of portunus-server. It consumes the
//FileStoreAPI and returns an Engine interface for the HTTP server to use, and
//a stream of events for the LDAP worker.
func RunEngineAsync(fsAPI *FileStoreAPI, ldapSuffix string) (Engine, <-chan []LDAPObject) {
	ldapUpdatesChan := make(chan []LDAPObject, 1)
	e := engine{
		FileStoreAPI: fsAPI,
		LDAPUpdates:  ldapUpdatesChan,
		Mutex:        &sync.RWMutex{},
		LDAPSuffix:   ldapSuffix,
	}

	go func() {
		for db := range e.FileStoreAPI.LoadEvents {
			e.handleLoadEvent(db)
		}
	}()

	return &e, ldapUpdatesChan
}

func (e *engine) handleLoadEvent(db Database) {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()

	e.Groups = make(map[string]Group, len(db.Groups))
	for _, group := range db.Groups {
		e.Groups[group.Name] = group
	}
	e.Users = make(map[string]User, len(db.Users))
	for _, user := range db.Users {
		e.Users[user.LoginName] = user
	}
	e.persistLDAP()
}

//FindUser implements the Engine interface.
func (e *engine) FindGroup(name string) *Group {
	e.Mutex.RLock()
	defer e.Mutex.RUnlock()

	g, exists := e.Groups[name]
	if !exists {
		return nil
	}
	g = g.Cloned()
	return &g
}

//FindUser implements the Engine interface.
func (e *engine) FindUser(loginName string) *UserWithPerms {
	e.Mutex.RLock()
	defer e.Mutex.RUnlock()

	u, exists := e.Users[loginName]
	if !exists {
		return nil
	}

	user := UserWithPerms{User: u.Cloned()}
	for _, group := range e.Groups {
		if group.ContainsUser(u) {
			user.GroupMemberships = append(user.GroupMemberships, group.Cloned())
			user.Perms = user.Perms.Union(group.Permissions)
		}
	}
	return &user
}

//ListGroups implements the Engine interface.
func (e *engine) ListGroups() []Group {
	e.Mutex.RLock()
	defer e.Mutex.RUnlock()

	result := make([]Group, 0, len(e.Groups))
	for _, group := range e.Groups {
		result = append(result, group.Cloned())
	}
	return result
}

//ListUsers implements the Engine interface.
func (e *engine) ListUsers() []User {
	e.Mutex.RLock()
	defer e.Mutex.RUnlock()

	result := make([]User, 0, len(e.Users))
	for _, user := range e.Users {
		result = append(result, user.Cloned())
	}
	return result
}

//ChangeUser implements the Engine interface.
func (e *engine) ChangeUser(loginName string, action func(User) (*User, error)) error {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()

	oldState, exists := e.Users[loginName]
	oldStatePtr := &oldState
	if !exists {
		oldStatePtr = nil
	}
	newState, err := action(oldState.Cloned())
	if err != nil {
		return err
	}

	//only change database if there are actual changes
	if newState == nil {
		if oldStatePtr == nil {
			return nil
		}
		delete(e.Users, loginName)
	} else {
		if oldState == *newState {
			return nil
		}
		e.Users[loginName] = *newState
	}

	e.persistDatabase()
	e.persistLDAP()
	return nil
}

//ChangeGroup implements the Engine interface.
func (e *engine) ChangeGroup(name string, action func(Group) (*Group, error)) error {
	e.Mutex.Lock()
	defer e.Mutex.Unlock()

	oldState, exists := e.Groups[name]
	oldStatePtr := &oldState
	if !exists {
		oldStatePtr = nil
	}
	newState, err := action(oldState.Cloned())
	if err != nil {
		return err
	}

	//only change database if there are actual changes
	if newState == nil {
		if oldStatePtr == nil {
			return nil
		}
		delete(e.Groups, name)
	} else {
		if oldState.IsEqualTo(*newState) {
			return nil
		}
		e.Groups[name] = *newState
	}

	e.persistDatabase()
	e.persistLDAP()
	return nil
}

func (e *engine) persistDatabase() {
	//NOTE: This is always called from functions that have locked e.Mutex, so we
	//don't need to do it ourselves.
	var db Database
	for _, user := range e.Users {
		db.Users = append(db.Users, user.Cloned())
	}
	for _, group := range e.Groups {
		db.Groups = append(db.Groups, group.Cloned())
	}
	e.FileStoreAPI.SaveRequests <- db
}

func (e *engine) persistLDAP() {
	//NOTE: This is always called from functions that have locked e.Mutex, so we
	//don't need to do it ourselves.
	ldapDB := make([]LDAPObject, 0, len(e.Users)+len(e.Groups))
	for _, user := range e.Users {
		ldapDB = append(ldapDB, user.RenderToLDAP(e.LDAPSuffix))
	}
	for _, group := range e.Groups {
		ldapDB = append(ldapDB, group.RenderToLDAP(e.LDAPSuffix)...)
	}
	e.LDAPUpdates <- ldapDB
}
