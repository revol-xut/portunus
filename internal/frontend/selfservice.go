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

package frontend

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/majewsky/portunus/internal/core"
	h "github.com/majewsky/portunus/internal/html"
)

// TODO: allow flipped order (family name first)
var userFullNameSnippet = h.NewSnippet(`
	<span class="given-name">{{.GivenName}}</span> <span class="family-name">{{.FamilyName}}</span>
`)
var userEMailAddressSnippet = h.NewSnippet(`
	{{if .EMailAddress}}{{.EMailAddress}}{{else}}<em>Not specified</em>{{end}}
`)

func useSelfServiceForm(e core.Engine) HandlerStep {
	return func(i *Interaction) {
		user := i.CurrentUser

		isAdmin := user.Perms.Portunus.IsAdmin
		visibleGroups := user.GroupMemberships
		if isAdmin {
			visibleGroups = e.ListGroups()
		}
		sort.Slice(visibleGroups, func(i, j int) bool {
			return visibleGroups[i].LongName < visibleGroups[j].LongName
		})

		var memberships []h.SelectOptionSpec
		isSelected := make(map[string]bool)
		for _, group := range visibleGroups {
			membership := h.SelectOptionSpec{
				Value: group.Name,
				Label: group.LongName,
			}
			memberships = append(memberships, membership)
			isSelected[group.Name] = group.ContainsUser(user.User)
		}

		i.FormState = &h.FormState{
			Fields: map[string]*h.FieldState{
				"memberships": {
					Selected: isSelected,
				},
				"ssh_public_keys": {
					Value: strings.Join(user.SSHPublicKeys, "\r\n"),
				},
			},
		}

		i.FormSpec = &h.FormSpec{
			PostTarget:  "/self",
			SubmitLabel: "Update profile",
			Fields: []h.FormField{
				h.StaticField{
					Label: "Login name",
					Value: codeTagSnippet.Render(user.LoginName),
				},
				h.StaticField{
					Label: "Full name",
					Value: userFullNameSnippet.Render(user),
				},
				h.StaticField{
					Label: "Email address",
					Value: userEMailAddressSnippet.Render(user),
				},
				h.SelectFieldSpec{
					Name:     "memberships",
					Label:    "Group memberships",
					Options:  memberships,
					ReadOnly: true,
				},
				h.MultilineInputFieldSpec{
					Name:  "ssh_public_keys",
					Label: "SSH public key(s)",
					Rules: []h.ValidationRule{
						core.MustNotHaveSurroundingSpaces,
						core.MustBeSSHPublicKeys,
					},
				},
				h.FieldSet{
					Name:       "change_password",
					Label:      "Change password",
					IsFoldable: true,
					Fields: []h.FormField{
						h.InputFieldSpec{
							InputType: "password",
							Name:      "old_password",
							Label:     "Old password",
							Rules: []h.ValidationRule{
								core.MustNotBeEmpty,
							},
						},
						h.InputFieldSpec{
							InputType: "password",
							Name:      "new_password",
							Label:     "New password",
							Rules: []h.ValidationRule{
								core.MustNotBeEmpty,
							},
						},
						h.InputFieldSpec{
							InputType: "password",
							Name:      "repeat_password",
							Label:     "Repeat password",
							Rules: []h.ValidationRule{
								core.MustNotBeEmpty,
							},
						},
					},
				},
			},
		}
	}
}

func getSelfHandler(e core.Engine) http.Handler {
	return Do(
		LoadSession,
		VerifyLogin(e),
		useSelfServiceForm(e),
		ShowForm("My profile"),
	)
}

func postSelfHandler(e core.Engine) http.Handler {
	return Do(
		LoadSession,
		VerifyLogin(e),
		useSelfServiceForm(e),
		ReadFormStateFromRequest,
		validateSelfServiceForm,
		ShowFormIfErrors("My profile"),
		executeSelfServiceForm(e),
		ShowForm("My profile"),
	)
}

func validateSelfServiceForm(i *Interaction) {
	fs := i.FormState

	if fs.Fields["change_password"].IsUnfolded {
		if fs.IsValid() {
			newPassword1 := fs.Fields["new_password"].Value
			newPassword2 := fs.Fields["repeat_password"].Value
			if newPassword1 != newPassword2 {
				fs.Fields["repeat_password"].ErrorMessage = "did not match"
			}
		}

		if fs.IsValid() {
			oldPassword := fs.Fields["old_password"].Value
			if !core.CheckPasswordHash(oldPassword, i.CurrentUser.PasswordHash) {
				fs.Fields["old_password"].ErrorMessage = "is not correct"
			}
		}
	}
}

func executeSelfServiceForm(e core.Engine) HandlerStep {
	return func(i *Interaction) {
		fs := i.FormState

		err := e.ChangeUser(i.CurrentUser.LoginName, func(u core.User) (*core.User, error) {
			if u.LoginName == "" {
				return nil, fmt.Errorf("no such user")
			}
			if fs.Fields["change_password"].IsUnfolded {
				u.PasswordHash = core.HashPasswordForLDAP(i.FormState.Fields["new_password"].Value)
			}
			u.SSHPublicKeys = core.SplitSSHPublicKeys(i.FormState.Fields["ssh_public_keys"].Value)
			return &u, nil
		})

		if err == nil {
			i.Session.AddFlash(Flash{"success", "Profile updated."})
		} else {
			i.Session.AddFlash(Flash{"danger", err.Error()})
		}
	}
}
