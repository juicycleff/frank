// Code generated by goa v3.20.0, DO NOT EDIT.
//
// User types
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package designtypes

// User is the result type of the auth service me method.
type User struct {
	// Whether account is active
	Active bool
	// Whether email is verified
	EmailVerified bool `json:"email_verified,emailVerified"`
	// Whether phone is verified
	PhoneVerified *bool `json:"phone_verified,phoneVerified"`
	// URL to user's profile image
	ProfileImageURL *string `json:"profile_image_url,profileImageUrl"`
	// User first name
	FirstName *string `json:"first_name,firstName"`
	// User last name
	LastName *string `json:"last_name,lastName"`
	// User phone number
	PhoneNumber *string `json:"phone_number,phoneNumber"`
	// User metadata
	Metadata map[string]any `json:"metadata"`
	// User locale
	Locale string `json:"locale"`
	// Email address
	Email string `json:"email"`
}
