// Code generated by goa v3.20.0, DO NOT EDIT.
//
// User types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package designtypes

// Organization information
type OrganizationSettings struct {
	// Signup fields
	SignupFields []*FormField `json:"signupFields"`
	// Signup fields
	Verification []*OrganizationVerificationConfig `json:"verification"`
}
