// Code generated by goa v3.20.0, DO NOT EDIT.
//
// oauth_client HTTP client types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	oauthclient "github.com/juicycleff/frank/gen/oauth_client"
	goa "goa.design/goa/v3/pkg"
)

// ListProvidersResponseBody is the type of the "oauth_client" service
// "list_providers" endpoint HTTP response body.
type ListProvidersResponseBody struct {
	Providers []*SSOProviderResponseBody `form:"providers,omitempty" json:"providers,omitempty" xml:"providers,omitempty"`
}

// ProviderCallbackResponseBody is the type of the "oauth_client" service
// "provider_callback" endpoint HTTP response body.
type ProviderCallbackResponseBody struct {
	// Whether authentication was successful
	Authenticated *bool `form:"authenticated,omitempty" json:"authenticated,omitempty" xml:"authenticated,omitempty"`
	// User data if authentication successful
	User *UserResponseBody `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
	// Success or error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListProvidersBadRequestResponseBody is the type of the "oauth_client"
// service "list_providers" endpoint HTTP response body for the "bad_request"
// error.
type ListProvidersBadRequestResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ListProvidersForbiddenResponseBody is the type of the "oauth_client" service
// "list_providers" endpoint HTTP response body for the "forbidden" error.
type ListProvidersForbiddenResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ListProvidersInternalErrorResponseBody is the type of the "oauth_client"
// service "list_providers" endpoint HTTP response body for the
// "internal_error" error.
type ListProvidersInternalErrorResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ListProvidersNotFoundResponseBody is the type of the "oauth_client" service
// "list_providers" endpoint HTTP response body for the "not_found" error.
type ListProvidersNotFoundResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ListProvidersUnauthorizedResponseBody is the type of the "oauth_client"
// service "list_providers" endpoint HTTP response body for the "unauthorized"
// error.
type ListProvidersUnauthorizedResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderAuthBadRequestResponseBody is the type of the "oauth_client" service
// "provider_auth" endpoint HTTP response body for the "bad_request" error.
type ProviderAuthBadRequestResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderAuthForbiddenResponseBody is the type of the "oauth_client" service
// "provider_auth" endpoint HTTP response body for the "forbidden" error.
type ProviderAuthForbiddenResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderAuthInternalErrorResponseBody is the type of the "oauth_client"
// service "provider_auth" endpoint HTTP response body for the "internal_error"
// error.
type ProviderAuthInternalErrorResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderAuthNotFoundResponseBody is the type of the "oauth_client" service
// "provider_auth" endpoint HTTP response body for the "not_found" error.
type ProviderAuthNotFoundResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderAuthUnauthorizedResponseBody is the type of the "oauth_client"
// service "provider_auth" endpoint HTTP response body for the "unauthorized"
// error.
type ProviderAuthUnauthorizedResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderCallbackBadRequestResponseBody is the type of the "oauth_client"
// service "provider_callback" endpoint HTTP response body for the
// "bad_request" error.
type ProviderCallbackBadRequestResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderCallbackForbiddenResponseBody is the type of the "oauth_client"
// service "provider_callback" endpoint HTTP response body for the "forbidden"
// error.
type ProviderCallbackForbiddenResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderCallbackInternalErrorResponseBody is the type of the "oauth_client"
// service "provider_callback" endpoint HTTP response body for the
// "internal_error" error.
type ProviderCallbackInternalErrorResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderCallbackNotFoundResponseBody is the type of the "oauth_client"
// service "provider_callback" endpoint HTTP response body for the "not_found"
// error.
type ProviderCallbackNotFoundResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ProviderCallbackUnauthorizedResponseBody is the type of the "oauth_client"
// service "provider_callback" endpoint HTTP response body for the
// "unauthorized" error.
type ProviderCallbackUnauthorizedResponseBody struct {
	// Error code
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// Error message
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Additional error details
	Details any `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
	// Unique error ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// SSOProviderResponseBody is used to define fields on response body types.
type SSOProviderResponseBody struct {
	// Provider ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Provider name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Provider type (oauth2, oidc, saml)
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Provider icon URL
	IconURL *string `form:"icon_url,omitempty" json:"icon_url,omitempty" xml:"icon_url,omitempty"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	// Whether account is active
	Active *bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// Whether email is verified
	EmailVerified *bool `json:"email_verified,emailVerified"`
	// Whether phone is verified
	PhoneVerified *bool `json:"phone_verified,phoneVerified"`
	// URL to user's profile image
	ProfileImageURL *string `json:"profile_image_url,profileImageUrl"`
	// User first name
	FirstName *string `json:"first_name,firstName"`
	// User last name
	LastName *string `json:"last_name,lastName"`
	// ID of the entity
	ID *string `json:"id"`
	// User phone number
	PhoneNumber *string `json:"phone_number,phoneNumber"`
	// User metadata
	Metadata map[string]any `json:"metadata"`
	// User locale
	Locale *string `json:"locale"`
	// Email address
	Email *string `json:"email"`
}

// NewListProvidersResultOK builds a "oauth_client" service "list_providers"
// endpoint result from a HTTP "OK" response.
func NewListProvidersResultOK(body *ListProvidersResponseBody) *oauthclient.ListProvidersResult {
	v := &oauthclient.ListProvidersResult{}
	v.Providers = make([]*oauthclient.SSOProvider, len(body.Providers))
	for i, val := range body.Providers {
		v.Providers[i] = unmarshalSSOProviderResponseBodyToOauthclientSSOProvider(val)
	}

	return v
}

// NewListProvidersBadRequest builds a oauth_client service list_providers
// endpoint bad_request error.
func NewListProvidersBadRequest(body *ListProvidersBadRequestResponseBody) *oauthclient.BadRequestError {
	v := &oauthclient.BadRequestError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewListProvidersForbidden builds a oauth_client service list_providers
// endpoint forbidden error.
func NewListProvidersForbidden(body *ListProvidersForbiddenResponseBody) *oauthclient.ForbiddenError {
	v := &oauthclient.ForbiddenError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewListProvidersInternalError builds a oauth_client service list_providers
// endpoint internal_error error.
func NewListProvidersInternalError(body *ListProvidersInternalErrorResponseBody) *oauthclient.InternalServerError {
	v := &oauthclient.InternalServerError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewListProvidersNotFound builds a oauth_client service list_providers
// endpoint not_found error.
func NewListProvidersNotFound(body *ListProvidersNotFoundResponseBody) *oauthclient.NotFoundError {
	v := &oauthclient.NotFoundError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewListProvidersUnauthorized builds a oauth_client service list_providers
// endpoint unauthorized error.
func NewListProvidersUnauthorized(body *ListProvidersUnauthorizedResponseBody) *oauthclient.UnauthorizedError {
	v := &oauthclient.UnauthorizedError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderAuthBadRequest builds a oauth_client service provider_auth
// endpoint bad_request error.
func NewProviderAuthBadRequest(body *ProviderAuthBadRequestResponseBody) *oauthclient.BadRequestError {
	v := &oauthclient.BadRequestError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderAuthForbidden builds a oauth_client service provider_auth
// endpoint forbidden error.
func NewProviderAuthForbidden(body *ProviderAuthForbiddenResponseBody) *oauthclient.ForbiddenError {
	v := &oauthclient.ForbiddenError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderAuthInternalError builds a oauth_client service provider_auth
// endpoint internal_error error.
func NewProviderAuthInternalError(body *ProviderAuthInternalErrorResponseBody) *oauthclient.InternalServerError {
	v := &oauthclient.InternalServerError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderAuthNotFound builds a oauth_client service provider_auth endpoint
// not_found error.
func NewProviderAuthNotFound(body *ProviderAuthNotFoundResponseBody) *oauthclient.NotFoundError {
	v := &oauthclient.NotFoundError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderAuthUnauthorized builds a oauth_client service provider_auth
// endpoint unauthorized error.
func NewProviderAuthUnauthorized(body *ProviderAuthUnauthorizedResponseBody) *oauthclient.UnauthorizedError {
	v := &oauthclient.UnauthorizedError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderCallbackResultOK builds a "oauth_client" service
// "provider_callback" endpoint result from a HTTP "OK" response.
func NewProviderCallbackResultOK(body *ProviderCallbackResponseBody) *oauthclient.ProviderCallbackResult {
	v := &oauthclient.ProviderCallbackResult{
		Authenticated: *body.Authenticated,
		Message:       *body.Message,
	}
	if body.User != nil {
		v.User = unmarshalUserResponseBodyToDesigntypesUser(body.User)
	}

	return v
}

// NewProviderCallbackBadRequest builds a oauth_client service
// provider_callback endpoint bad_request error.
func NewProviderCallbackBadRequest(body *ProviderCallbackBadRequestResponseBody) *oauthclient.BadRequestError {
	v := &oauthclient.BadRequestError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderCallbackForbidden builds a oauth_client service provider_callback
// endpoint forbidden error.
func NewProviderCallbackForbidden(body *ProviderCallbackForbiddenResponseBody) *oauthclient.ForbiddenError {
	v := &oauthclient.ForbiddenError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderCallbackInternalError builds a oauth_client service
// provider_callback endpoint internal_error error.
func NewProviderCallbackInternalError(body *ProviderCallbackInternalErrorResponseBody) *oauthclient.InternalServerError {
	v := &oauthclient.InternalServerError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderCallbackNotFound builds a oauth_client service provider_callback
// endpoint not_found error.
func NewProviderCallbackNotFound(body *ProviderCallbackNotFoundResponseBody) *oauthclient.NotFoundError {
	v := &oauthclient.NotFoundError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// NewProviderCallbackUnauthorized builds a oauth_client service
// provider_callback endpoint unauthorized error.
func NewProviderCallbackUnauthorized(body *ProviderCallbackUnauthorizedResponseBody) *oauthclient.UnauthorizedError {
	v := &oauthclient.UnauthorizedError{
		Code:    *body.Code,
		Message: *body.Message,
		Details: body.Details,
		ID:      body.ID,
	}

	return v
}

// ValidateListProvidersResponseBody runs the validations defined on
// list_providers_response_body
func ValidateListProvidersResponseBody(body *ListProvidersResponseBody) (err error) {
	if body.Providers == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("providers", "body"))
	}
	for _, e := range body.Providers {
		if e != nil {
			if err2 := ValidateSSOProviderResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateProviderCallbackResponseBody runs the validations defined on
// provider_callback_response_body
func ValidateProviderCallbackResponseBody(body *ProviderCallbackResponseBody) (err error) {
	if body.Authenticated == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("authenticated", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateListProvidersBadRequestResponseBody runs the validations defined on
// list_providers_bad_request_response_body
func ValidateListProvidersBadRequestResponseBody(body *ListProvidersBadRequestResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProvidersForbiddenResponseBody runs the validations defined on
// list_providers_forbidden_response_body
func ValidateListProvidersForbiddenResponseBody(body *ListProvidersForbiddenResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProvidersInternalErrorResponseBody runs the validations defined
// on list_providers_internal_error_response_body
func ValidateListProvidersInternalErrorResponseBody(body *ListProvidersInternalErrorResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProvidersNotFoundResponseBody runs the validations defined on
// list_providers_not_found_response_body
func ValidateListProvidersNotFoundResponseBody(body *ListProvidersNotFoundResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateListProvidersUnauthorizedResponseBody runs the validations defined
// on list_providers_unauthorized_response_body
func ValidateListProvidersUnauthorizedResponseBody(body *ListProvidersUnauthorizedResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderAuthBadRequestResponseBody runs the validations defined on
// provider_auth_bad_request_response_body
func ValidateProviderAuthBadRequestResponseBody(body *ProviderAuthBadRequestResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderAuthForbiddenResponseBody runs the validations defined on
// provider_auth_forbidden_response_body
func ValidateProviderAuthForbiddenResponseBody(body *ProviderAuthForbiddenResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderAuthInternalErrorResponseBody runs the validations defined
// on provider_auth_internal_error_response_body
func ValidateProviderAuthInternalErrorResponseBody(body *ProviderAuthInternalErrorResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderAuthNotFoundResponseBody runs the validations defined on
// provider_auth_not_found_response_body
func ValidateProviderAuthNotFoundResponseBody(body *ProviderAuthNotFoundResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderAuthUnauthorizedResponseBody runs the validations defined on
// provider_auth_unauthorized_response_body
func ValidateProviderAuthUnauthorizedResponseBody(body *ProviderAuthUnauthorizedResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderCallbackBadRequestResponseBody runs the validations defined
// on provider_callback_bad_request_response_body
func ValidateProviderCallbackBadRequestResponseBody(body *ProviderCallbackBadRequestResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderCallbackForbiddenResponseBody runs the validations defined
// on provider_callback_forbidden_response_body
func ValidateProviderCallbackForbiddenResponseBody(body *ProviderCallbackForbiddenResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderCallbackInternalErrorResponseBody runs the validations
// defined on provider_callback_internal_error_response_body
func ValidateProviderCallbackInternalErrorResponseBody(body *ProviderCallbackInternalErrorResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderCallbackNotFoundResponseBody runs the validations defined on
// provider_callback_not_found_response_body
func ValidateProviderCallbackNotFoundResponseBody(body *ProviderCallbackNotFoundResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateProviderCallbackUnauthorizedResponseBody runs the validations
// defined on provider_callback_unauthorized_response_body
func ValidateProviderCallbackUnauthorizedResponseBody(body *ProviderCallbackUnauthorizedResponseBody) (err error) {
	if body.Code == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("code", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateSSOProviderResponseBody runs the validations defined on
// SSOProviderResponseBody
func ValidateSSOProviderResponseBody(body *SSOProviderResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}

// ValidateUserResponseBody runs the validations defined on UserResponseBody
func ValidateUserResponseBody(body *UserResponseBody) (err error) {
	if body.Active == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("active", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.EmailVerified == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email_verified", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}
