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
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// Group represents a single group of users. Membership in a group implicitly
// grants its Permissions to all users in that group.
type Group struct {
	Name             string           `json:"name"`
	LongName         string           `json:"long_name"`
	MemberLoginNames GroupMemberNames `json:"members"`
	Permissions      Permissions      `json:"permissions"`
	PosixGID         *PosixID         `json:"posix_gid,omitempty"`
}

// Cloned returns a deep copy of this user.
func (g Group) Cloned() Group {
	logins := g.MemberLoginNames
	g.MemberLoginNames = make(GroupMemberNames)
	for name, isMember := range logins {
		if isMember {
			g.MemberLoginNames[name] = true
		}
	}
	if g.PosixGID != nil {
		val := *g.PosixGID
		g.PosixGID = &val
	}
	return g
}

// ContainsUser checks whether this group contains the given user.
func (g Group) ContainsUser(u User) bool {
	return g.MemberLoginNames[u.LoginName]
}

// IsEqualTo is a type-safe wrapper around reflect.DeepEqual().
func (g Group) IsEqualTo(other Group) bool {
	return reflect.DeepEqual(g, other)
}

// RenderToLDAP produces the LDAPObject representing this group.
func (g Group) RenderToLDAP(suffix string) []LDAPObject {
	memberDNames := make([]string, 0, len(g.MemberLoginNames))
	memberLoginNames := make([]string, 0, len(g.MemberLoginNames))
	for name, isMember := range g.MemberLoginNames {
		if isMember {
			memberDNames = append(memberDNames, fmt.Sprintf("uid=%s,ou=users,%s", name, suffix))
			memberLoginNames = append(memberLoginNames, name)
		}
	}
	if len(memberDNames) == 0 {
		//The OpenLDAP core.schema requires that `groupOfNames` contain at least
		//one `member` attribute. If the group does not have any proper members,
		//add the dummy user account "nobody" to it.
		memberDNames = append(memberDNames, "cn=nobody,"+suffix)
	}

	objs := []LDAPObject{{
		DN: fmt.Sprintf("cn=%s,ou=groups,%s", g.Name, suffix),
		Attributes: map[string][]string{
			"cn":          {g.Name},
			"member":      memberDNames,
			"objectClass": {"groupOfNames", "top"},
		},
	}}
	if g.PosixGID != nil {
		objs = append(objs, LDAPObject{
			DN: fmt.Sprintf("cn=%s,ou=posix-groups,%s", g.Name, suffix),
			Attributes: map[string][]string{
				"cn":          {g.Name},
				"gidNumber":   {g.PosixGID.String()},
				"memberUid":   memberLoginNames,
				"objectClass": {"posixGroup", "top"},
			},
		})
	}
	return objs
}

// GroupMemberNames is the type of Group.MemberLoginNames.
type GroupMemberNames map[string]bool

// MarshalJSON implements the json.Marshaler interface.
func (g GroupMemberNames) MarshalJSON() ([]byte, error) {
	names := make([]string, 0, len(g))
	for name, isMember := range g {
		if isMember {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return json.Marshal(names)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (g *GroupMemberNames) UnmarshalJSON(data []byte) error {
	var names []string
	err := json.Unmarshal(data, &names)
	if err != nil {
		return err
	}
	*g = make(map[string]bool)
	for _, name := range names {
		(*g)[name] = true
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// PosixID represents a POSIX user or group ID.
type PosixID uint16

func (id PosixID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
