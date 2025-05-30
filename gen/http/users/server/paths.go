// Code generated by goa v3.20.0, DO NOT EDIT.
//
// HTTP request path constructors for the users service.
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"fmt"
)

// ListUsersPath returns the URL path to the users service list HTTP endpoint.
func ListUsersPath() string {
	return "/v1/users"
}

// CreateUsersPath returns the URL path to the users service create HTTP endpoint.
func CreateUsersPath() string {
	return "/v1/users"
}

// GetUsersPath returns the URL path to the users service get HTTP endpoint.
func GetUsersPath(id string) string {
	return fmt.Sprintf("/v1/users/%v", id)
}

// UpdateUsersPath returns the URL path to the users service update HTTP endpoint.
func UpdateUsersPath(id string) string {
	return fmt.Sprintf("/v1/users/%v", id)
}

// DeleteUsersPath returns the URL path to the users service delete HTTP endpoint.
func DeleteUsersPath(id string) string {
	return fmt.Sprintf("/v1/users/%v", id)
}

// UpdateMeUsersPath returns the URL path to the users service update_me HTTP endpoint.
func UpdateMeUsersPath() string {
	return "/v1/users/me"
}

// UpdatePasswordUsersPath returns the URL path to the users service update_password HTTP endpoint.
func UpdatePasswordUsersPath() string {
	return "/v1/users/me/password"
}

// GetSessionsUsersPath returns the URL path to the users service get_sessions HTTP endpoint.
func GetSessionsUsersPath() string {
	return "/v1/users/me/sessions"
}

// DeleteSessionUsersPath returns the URL path to the users service delete_session HTTP endpoint.
func DeleteSessionUsersPath(sessionID string) string {
	return fmt.Sprintf("/v1/users/me/sessions/%v", sessionID)
}

// GetOrganizationsUsersPath returns the URL path to the users service get_organizations HTTP endpoint.
func GetOrganizationsUsersPath(id string) string {
	return fmt.Sprintf("/v1/users/%v/organizations", id)
}
