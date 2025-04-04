// Code generated by goa v3.20.0, DO NOT EDIT.
//
// User types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package designtypes

// MFA verification request
type MFAVerifyRequest struct {
	// MFA method to verify
	Method string
	// Verification code
	Code string
	// Phone number for SMS verification
	PhoneNumber *string
}
