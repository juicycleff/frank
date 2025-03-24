// Code generated by goa v3.20.0, DO NOT EDIT.
//
// oauth_client service
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package oauthclient

import (
	"context"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	"goa.design/goa/v3/security"
)

// OAuth2 client service for authenticating with external providers
type Service interface {
	// List available OAuth providers
	ListProviders(context.Context, *ListProvidersPayload) (res *ListProvidersResult, err error)
	// Initiate authentication with an OAuth provider
	ProviderAuth(context.Context, *ProviderAuthPayload) (err error)
	// Handle OAuth provider callback
	ProviderCallback(context.Context, *ProviderCallbackPayload) (res *ProviderCallbackResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
	// APIKeyAuth implements the authorization logic for the APIKey security scheme.
	APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error)
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// APIName is the name of the API as defined in the design.
const APIName = "frank"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "1.0.0"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "oauth_client"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"list_providers", "provider_auth", "provider_callback"}

// Bad request response
type BadRequestError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// Forbidden response
type ForbiddenError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// Internal server error response
type InternalServerError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// ListProvidersPayload is the payload type of the oauth_client service
// list_providers method.
type ListProvidersPayload struct {
	// OAuth2 access token
	Oauth2 *string
	// API key
	XAPIKey *string
	// JWT token
	JWT *string
}

// ListProvidersResult is the result type of the oauth_client service
// list_providers method.
type ListProvidersResult struct {
	Providers []*SSOProvider
}

// Not found response
type NotFoundError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// ProviderAuthPayload is the payload type of the oauth_client service
// provider_auth method.
type ProviderAuthPayload struct {
	// OAuth2 access token
	Oauth2 *string
	// API key
	XAPIKey *string
	// JWT token
	JWT *string
	// Provider ID
	Provider string
	// Redirect URI after authentication
	RedirectURI *string
}

// ProviderCallbackPayload is the payload type of the oauth_client service
// provider_callback method.
type ProviderCallbackPayload struct {
	// OAuth2 access token
	Oauth2 *string
	// API key
	XAPIKey *string
	// JWT token
	JWT *string
	// Provider ID
	Provider string
	// Authorization code
	Code *string
	// State parameter
	State *string
}

// ProviderCallbackResult is the result type of the oauth_client service
// provider_callback method.
type ProviderCallbackResult struct {
	// Whether authentication was successful
	Authenticated bool
	// User data if authentication successful
	User *designtypes.User
	// Success or error message
	Message string
}

// SSO Provider information
type SSOProvider struct {
	// Provider ID
	ID string
	// Provider name
	Name string
	// Provider type (oauth2, oidc, saml)
	Type string
	// Provider icon URL
	IconURL *string
}

// Unauthorized response
type UnauthorizedError struct {
	// Error code
	Code string
	// Error message
	Message string
	// Additional error details
	Details any
	// Unique error ID
	ID *string
}

// Error returns an error description.
func (e *BadRequestError) Error() string {
	return "Bad request response"
}

// ErrorName returns "BadRequestError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *BadRequestError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "BadRequestError".
func (e *BadRequestError) GoaErrorName() string {
	return "bad_request"
}

// Error returns an error description.
func (e *ForbiddenError) Error() string {
	return "Forbidden response"
}

// ErrorName returns "ForbiddenError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *ForbiddenError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "ForbiddenError".
func (e *ForbiddenError) GoaErrorName() string {
	return "forbidden"
}

// Error returns an error description.
func (e *InternalServerError) Error() string {
	return "Internal server error response"
}

// ErrorName returns "InternalServerError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *InternalServerError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "InternalServerError".
func (e *InternalServerError) GoaErrorName() string {
	return "internal_error"
}

// Error returns an error description.
func (e *NotFoundError) Error() string {
	return "Not found response"
}

// ErrorName returns "NotFoundError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *NotFoundError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "NotFoundError".
func (e *NotFoundError) GoaErrorName() string {
	return "not_found"
}

// Error returns an error description.
func (e *UnauthorizedError) Error() string {
	return "Unauthorized response"
}

// ErrorName returns "UnauthorizedError".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e *UnauthorizedError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "UnauthorizedError".
func (e *UnauthorizedError) GoaErrorName() string {
	return "unauthorized"
}
