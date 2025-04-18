// Code generated by goa v3.20.0, DO NOT EDIT.
//
// oauth_client HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	oauthclient "github.com/juicycleff/frank/gen/oauth_client"
)

// BuildListProvidersPayload builds the payload for the oauth_client
// list_providers endpoint from CLI flags.
func BuildListProvidersPayload(oauthClientListProvidersOauth2 string, oauthClientListProvidersXAPIKey string, oauthClientListProvidersJWT string) (*oauthclient.ListProvidersPayload, error) {
	var oauth2 *string
	{
		if oauthClientListProvidersOauth2 != "" {
			oauth2 = &oauthClientListProvidersOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthClientListProvidersXAPIKey != "" {
			xAPIKey = &oauthClientListProvidersXAPIKey
		}
	}
	var jwt *string
	{
		if oauthClientListProvidersJWT != "" {
			jwt = &oauthClientListProvidersJWT
		}
	}
	v := &oauthclient.ListProvidersPayload{}
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildProviderAuthPayload builds the payload for the oauth_client
// provider_auth endpoint from CLI flags.
func BuildProviderAuthPayload(oauthClientProviderAuthProvider string, oauthClientProviderAuthRedirectURI string, oauthClientProviderAuthOauth2 string, oauthClientProviderAuthXAPIKey string, oauthClientProviderAuthJWT string) (*oauthclient.ProviderAuthPayload, error) {
	var provider string
	{
		provider = oauthClientProviderAuthProvider
	}
	var redirectURI *string
	{
		if oauthClientProviderAuthRedirectURI != "" {
			redirectURI = &oauthClientProviderAuthRedirectURI
		}
	}
	var oauth2 *string
	{
		if oauthClientProviderAuthOauth2 != "" {
			oauth2 = &oauthClientProviderAuthOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthClientProviderAuthXAPIKey != "" {
			xAPIKey = &oauthClientProviderAuthXAPIKey
		}
	}
	var jwt *string
	{
		if oauthClientProviderAuthJWT != "" {
			jwt = &oauthClientProviderAuthJWT
		}
	}
	v := &oauthclient.ProviderAuthPayload{}
	v.Provider = provider
	v.RedirectURI = redirectURI
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildProviderCallbackPayload builds the payload for the oauth_client
// provider_callback endpoint from CLI flags.
func BuildProviderCallbackPayload(oauthClientProviderCallbackProvider string, oauthClientProviderCallbackCode string, oauthClientProviderCallbackState string, oauthClientProviderCallbackOauth2 string, oauthClientProviderCallbackXAPIKey string, oauthClientProviderCallbackJWT string) (*oauthclient.ProviderCallbackPayload, error) {
	var provider string
	{
		provider = oauthClientProviderCallbackProvider
	}
	var code *string
	{
		if oauthClientProviderCallbackCode != "" {
			code = &oauthClientProviderCallbackCode
		}
	}
	var state *string
	{
		if oauthClientProviderCallbackState != "" {
			state = &oauthClientProviderCallbackState
		}
	}
	var oauth2 *string
	{
		if oauthClientProviderCallbackOauth2 != "" {
			oauth2 = &oauthClientProviderCallbackOauth2
		}
	}
	var xAPIKey *string
	{
		if oauthClientProviderCallbackXAPIKey != "" {
			xAPIKey = &oauthClientProviderCallbackXAPIKey
		}
	}
	var jwt *string
	{
		if oauthClientProviderCallbackJWT != "" {
			jwt = &oauthClientProviderCallbackJWT
		}
	}
	v := &oauthclient.ProviderCallbackPayload{}
	v.Provider = provider
	v.Code = code
	v.State = state
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}
