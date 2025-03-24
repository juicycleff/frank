// Code generated by goa v3.20.0, DO NOT EDIT.
//
// api_keys service
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package apikeys

import (
	"context"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// API key management service
type Service interface {
	// List API keys
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Create a new API key
	Create(context.Context, *CreatePayload) (res *APIKeyWithSecretResponse, err error)
	// Get API key by ID
	Get(context.Context, *GetPayload) (res *APIKeyResponse, err error)
	// Update API key
	Update(context.Context, *UpdatePayload) (res *APIKeyResponse, err error)
	// Delete API key
	Delete(context.Context, *DeletePayload) (err error)
	// Validate API key
	Validate(context.Context, *ValidatePayload) (res *ValidateResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
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
const ServiceName = "api_keys"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [6]string{"list", "create", "get", "update", "delete", "validate"}

// APIKeyResponse is the result type of the api_keys service get method.
type APIKeyResponse struct {
	// API key ID
	ID string
	// API key name
	Name string
	// User ID who owns the key
	UserID *string
	// Organization ID
	OrganizationID *string
	// API key type (client/server)
	Type string
	// Whether API key is active
	Active bool
	// Key permissions
	Permissions []string
	// Key scopes
	Scopes []string
	// Key metadata
	Metadata map[string]any
	// Last used timestamp
	LastUsed *string
	// Expiry timestamp
	ExpiresAt *string
	// Creation timestamp
	CreatedAt string
	// Last update timestamp
	UpdatedAt *string
}

// APIKeyWithSecretResponse is the result type of the api_keys service create
// method.
type APIKeyWithSecretResponse struct {
	// API key secret
	Key string
	// API key ID
	ID string
	// API key name
	Name string
	// User ID who owns the key
	UserID *string
	// Organization ID
	OrganizationID *string
	// API key type (client/server)
	Type string
	// Whether API key is active
	Active bool
	// Key permissions
	Permissions []string
	// Key scopes
	Scopes []string
	// Key metadata
	Metadata map[string]any
	// Last used timestamp
	LastUsed *string
	// Expiry timestamp
	ExpiresAt *string
	// Creation timestamp
	CreatedAt string
	// Last update timestamp
	UpdatedAt *string
}

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

// Create API key request
type CreateAPIKeyRequest struct {
	// API key name
	Name string
	// API key type
	Type string
	// Key permissions
	Permissions []string
	// Key scopes
	Scopes []string
	// Key metadata
	Metadata map[string]any
	// Expiry in seconds
	ExpiresIn *int
}

// CreatePayload is the payload type of the api_keys service create method.
type CreatePayload struct {
	JWT *string
	Key *CreateAPIKeyRequest
}

// DeletePayload is the payload type of the api_keys service delete method.
type DeletePayload struct {
	JWT *string
	// API key ID
	ID string
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

// GetPayload is the payload type of the api_keys service get method.
type GetPayload struct {
	JWT *string
	// API key ID
	ID string
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

// ListPayload is the payload type of the api_keys service list method.
type ListPayload struct {
	JWT *string
	// Pagination offset
	Offset int
	// Number of items to return
	Limit int
	// Filter by key type
	Type *string
	// Filter by organization ID
	OrganizationID *string
}

// ListResult is the result type of the api_keys service list method.
type ListResult struct {
	Data []*APIKeyResponse
	// Total number of keys
	Total      int
	Pagination *designtypes.Pagination
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

// Update API key request
type UpdateAPIKeyRequest struct {
	// API key name
	Name *string
	// Whether API key is active
	Active *bool
	// Key permissions
	Permissions []string
	// Key scopes
	Scopes []string
	// Key metadata
	Metadata map[string]any
	// Expiry timestamp
	ExpiresAt *string
}

// UpdatePayload is the payload type of the api_keys service update method.
type UpdatePayload struct {
	JWT *string
	// API key ID
	ID  string
	Key *UpdateAPIKeyRequest
}

// ValidatePayload is the payload type of the api_keys service validate method.
type ValidatePayload struct {
	// API key to validate
	APIKey string
}

// ValidateResult is the result type of the api_keys service validate method.
type ValidateResult struct {
	// Whether key is valid
	Valid bool
	// Key details if valid
	Key *APIKeyResponse
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

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "unauthorized", false, false, false)
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "bad_request", false, false, false)
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "not_found", false, false, false)
}
