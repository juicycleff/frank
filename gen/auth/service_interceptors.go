// Code generated by goa v3.20.0, DO NOT EDIT.
//
// authServer Interceptors
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// ServerInterceptors defines the interface for all server-side interceptors.
// Server interceptors execute after the request is decoded and before the
// payload is sent to the service. The implementation is responsible for calling
// next to complete the request.
type ServerInterceptors interface {
	// CSRF token interceptor
	CSRFToken(ctx context.Context, info *CSRFTokenInfo, next goa.Endpoint) (any, error)
}

// Access interfaces for interceptor payloads and results
type (
	// CSRFTokenInfo provides metadata about the current interception.
	// It includes service name, method name, and access to the endpoint.
	CSRFTokenInfo struct {
		service    string
		method     string
		callType   goa.InterceptorCallType
		rawPayload any
	}

	// CSRFTokenResult provides type-safe access to the method result.
	// It allows reading and writing specific fields of the result as defined
	// in the design.
	CSRFTokenResult interface {
		CsrfToken() string
		SetCsrfToken(string)
	}
)

// Private implementation types
type (
	cSRFTokenLoginPayload struct {
		payload *LoginPayload
	}
	cSRFTokenLoginResult struct {
		result *LoginResponse
	}
)

// WrapLoginEndpoint wraps the login endpoint with the server-side interceptors
// defined in the design.
func WrapLoginEndpoint(endpoint goa.Endpoint, i ServerInterceptors) goa.Endpoint {
	endpoint = wrapLoginCSRFToken(endpoint, i)
	return endpoint
}

// Public accessor methods for Info types

// Service returns the name of the service handling the request.
func (info *CSRFTokenInfo) Service() string {
	return info.service
}

// Method returns the name of the method handling the request.
func (info *CSRFTokenInfo) Method() string {
	return info.method
}

// CallType returns the type of call the interceptor is handling.
func (info *CSRFTokenInfo) CallType() goa.InterceptorCallType {
	return info.callType
}

// RawPayload returns the raw payload of the request.
func (info *CSRFTokenInfo) RawPayload() any {
	return info.rawPayload
}

// Result returns a type-safe accessor for the method result.
func (info *CSRFTokenInfo) Result(res any) CSRFTokenResult {
	return &cSRFTokenLoginResult{result: res.(*LoginResponse)}
}

// Private implementation methods

func (r *cSRFTokenLoginResult) CsrfToken() string {
	return r.result.CsrfToken
}
func (r *cSRFTokenLoginResult) SetCsrfToken(v string) {
	r.result.CsrfToken = v
}
