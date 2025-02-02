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
	"fmt"
)

// User represents a single user account.
type User struct {
	LoginName     string   `json:"login_name"`
	GivenName     string   `json:"given_name"`
	FamilyName    string   `json:"family_name"`
	EMailAddress  string   `json:"email,omitempty"`
	SSHPublicKeys []string `json:"ssh_public_keys,omitempty"`
	//PasswordHash must be in the format generated by crypt(3).
	PasswordHash string               `json:"password"`
	POSIX        *UserPosixAttributes `json:"posix,omitempty"`
}

// UserPosixAttributes appears in type User.
type UserPosixAttributes struct {
	UID           PosixID `json:"uid"`
	GID           PosixID `json:"gid"`
	HomeDirectory string  `json:"home"`
	LoginShell    string  `json:"shell"` //optional
	GECOS         string  `json:"gecos"` //optional
}

// Cloned returns a deep copy of this user.
func (u User) Cloned() User {
	if u.POSIX != nil {
		val := *u.POSIX
		u.POSIX = &val
	}
	if u.SSHPublicKeys != nil {
		u.SSHPublicKeys = append([]string(nil), u.SSHPublicKeys...)
	}
	return u
}

// FullName returns the user's full name.
func (u User) FullName() string {
	return u.GivenName + " " + u.FamilyName //TODO: allow flipped order (family name first)
}

// RenderToLDAP produces the LDAPObject representing this group.
func (u User) RenderToLDAP(suffix string, allGroups map[string]Group) LDAPObject {
	var memberOfGroupDNames []string
	for _, group := range allGroups {
		if group.ContainsUser(u) {
			dn := fmt.Sprintf("cn=%s,ou=groups,%s", group.Name, suffix)
			memberOfGroupDNames = append(memberOfGroupDNames, dn)
		}
	}

	obj := LDAPObject{
		DN: fmt.Sprintf("uid=%s,ou=users,%s", u.LoginName, suffix),
		Attributes: map[string][]string{
			"uid":          {u.LoginName},
			"cn":           {u.FullName()},
			"sn":           {u.FamilyName},
			"givenName":    {u.GivenName},
			"userPassword": {u.PasswordHash},
			"isMemberOf":   memberOfGroupDNames,
			"objectClass":  {"portunusPerson", "inetOrgPerson", "organizationalPerson", "person", "top"},
		},
	}

	if u.EMailAddress != "" {
		obj.Attributes["mail"] = []string{u.EMailAddress}
	}
	if len(u.SSHPublicKeys) > 0 {
		obj.Attributes["sshPublicKey"] = u.SSHPublicKeys
	}

	if u.POSIX != nil {
		obj.Attributes["uidNumber"] = []string{u.POSIX.UID.String()}
		obj.Attributes["gidNumber"] = []string{u.POSIX.GID.String()}
		obj.Attributes["homeDirectory"] = []string{u.POSIX.HomeDirectory}
		if u.POSIX.LoginShell != "" {
			obj.Attributes["loginShell"] = []string{u.POSIX.LoginShell}
		}
		if u.POSIX.GECOS == "" {
			obj.Attributes["gecos"] = []string{u.FullName()}
		} else {
			obj.Attributes["gecos"] = []string{u.POSIX.GECOS}
		}
		obj.Attributes["objectClass"] = append(obj.Attributes["objectClass"], "posixAccount")
	}

	return obj
}

////////////////////////////////////////////////////////////////////////////////

// UserWithPerms is a User that carries its computed set of permissions.
type UserWithPerms struct {
	User
	Perms            Permissions
	GroupMemberships []Group
}
