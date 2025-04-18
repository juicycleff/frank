// Code generated by goa v3.20.0, DO NOT EDIT.
//
// sso HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"

	sso "github.com/juicycleff/frank/gen/sso"
	goa "goa.design/goa/v3/pkg"
)

// BuildListProvidersPayload builds the payload for the sso list_providers
// endpoint from CLI flags.
func BuildListProvidersPayload(ssoListProvidersOrganizationID string, ssoListProvidersOauth2 string, ssoListProvidersXAPIKey string, ssoListProvidersJWT string) (*sso.ListProvidersPayload, error) {
	var organizationID *string
	{
		if ssoListProvidersOrganizationID != "" {
			organizationID = &ssoListProvidersOrganizationID
		}
	}
	var oauth2 *string
	{
		if ssoListProvidersOauth2 != "" {
			oauth2 = &ssoListProvidersOauth2
		}
	}
	var xAPIKey *string
	{
		if ssoListProvidersXAPIKey != "" {
			xAPIKey = &ssoListProvidersXAPIKey
		}
	}
	var jwt *string
	{
		if ssoListProvidersJWT != "" {
			jwt = &ssoListProvidersJWT
		}
	}
	v := &sso.ListProvidersPayload{}
	v.OrganizationID = organizationID
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildProviderAuthPayload builds the payload for the sso provider_auth
// endpoint from CLI flags.
func BuildProviderAuthPayload(ssoProviderAuthProvider string, ssoProviderAuthRedirectURI string, ssoProviderAuthOauth2 string, ssoProviderAuthXAPIKey string, ssoProviderAuthJWT string) (*sso.ProviderAuthPayload, error) {
	var provider string
	{
		provider = ssoProviderAuthProvider
	}
	var redirectURI *string
	{
		if ssoProviderAuthRedirectURI != "" {
			redirectURI = &ssoProviderAuthRedirectURI
		}
	}
	var oauth2 *string
	{
		if ssoProviderAuthOauth2 != "" {
			oauth2 = &ssoProviderAuthOauth2
		}
	}
	var xAPIKey *string
	{
		if ssoProviderAuthXAPIKey != "" {
			xAPIKey = &ssoProviderAuthXAPIKey
		}
	}
	var jwt *string
	{
		if ssoProviderAuthJWT != "" {
			jwt = &ssoProviderAuthJWT
		}
	}
	v := &sso.ProviderAuthPayload{}
	v.Provider = provider
	v.RedirectURI = redirectURI
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildProviderCallbackPayload builds the payload for the sso
// provider_callback endpoint from CLI flags.
func BuildProviderCallbackPayload(ssoProviderCallbackProvider string, ssoProviderCallbackCode string, ssoProviderCallbackState string, ssoProviderCallbackSAMLResponse string, ssoProviderCallbackRelayState string, ssoProviderCallbackOauth2 string, ssoProviderCallbackXAPIKey string, ssoProviderCallbackJWT string) (*sso.ProviderCallbackPayload, error) {
	var provider string
	{
		provider = ssoProviderCallbackProvider
	}
	var code *string
	{
		if ssoProviderCallbackCode != "" {
			code = &ssoProviderCallbackCode
		}
	}
	var state *string
	{
		if ssoProviderCallbackState != "" {
			state = &ssoProviderCallbackState
		}
	}
	var sAMLResponse *string
	{
		if ssoProviderCallbackSAMLResponse != "" {
			sAMLResponse = &ssoProviderCallbackSAMLResponse
		}
	}
	var relayState *string
	{
		if ssoProviderCallbackRelayState != "" {
			relayState = &ssoProviderCallbackRelayState
		}
	}
	var oauth2 *string
	{
		if ssoProviderCallbackOauth2 != "" {
			oauth2 = &ssoProviderCallbackOauth2
		}
	}
	var xAPIKey *string
	{
		if ssoProviderCallbackXAPIKey != "" {
			xAPIKey = &ssoProviderCallbackXAPIKey
		}
	}
	var jwt *string
	{
		if ssoProviderCallbackJWT != "" {
			jwt = &ssoProviderCallbackJWT
		}
	}
	v := &sso.ProviderCallbackPayload{}
	v.Provider = provider
	v.Code = code
	v.State = state
	v.SAMLResponse = sAMLResponse
	v.RelayState = relayState
	v.Oauth2 = oauth2
	v.XAPIKey = xAPIKey
	v.JWT = jwt

	return v, nil
}

// BuildListIdentityProvidersPayload builds the payload for the sso
// list_identity_providers endpoint from CLI flags.
func BuildListIdentityProvidersPayload(ssoListIdentityProvidersOrganizationID string, ssoListIdentityProvidersJWT string) (*sso.ListIdentityProvidersPayload, error) {
	var organizationID string
	{
		organizationID = ssoListIdentityProvidersOrganizationID
	}
	var jwt *string
	{
		if ssoListIdentityProvidersJWT != "" {
			jwt = &ssoListIdentityProvidersJWT
		}
	}
	v := &sso.ListIdentityProvidersPayload{}
	v.OrganizationID = organizationID
	v.JWT = jwt

	return v, nil
}

// BuildCreateIdentityProviderPayload builds the payload for the sso
// create_identity_provider endpoint from CLI flags.
func BuildCreateIdentityProviderPayload(ssoCreateIdentityProviderBody string, ssoCreateIdentityProviderJWT string) (*sso.CreateIdentityProviderPayload, error) {
	var err error
	var body CreateIdentityProviderRequestBody
	{
		err = json.Unmarshal([]byte(ssoCreateIdentityProviderBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"organization_id\": \"Nobis aliquam neque perspiciatis dolores totam.\",\n      \"provider\": {\n         \"active\": true,\n         \"attributes_mapping\": {\n            \"email\": \"email\",\n            \"name\": \"name\"\n         },\n         \"authorization_endpoint\": \"https://accounts.google.com/o/oauth2/auth\",\n         \"certificate\": \"Accusamus sint omnis.\",\n         \"client_id\": \"client_id_123\",\n         \"client_secret\": \"client_secret_456\",\n         \"domains\": [\n            \"example.com\"\n         ],\n         \"issuer\": \"https://accounts.google.com\",\n         \"jwks_uri\": \"https://www.googleapis.com/oauth2/v3/certs\",\n         \"metadata_url\": \"Ipsam velit explicabo eius accusamus.\",\n         \"name\": \"Google\",\n         \"primary\": false,\n         \"private_key\": \"Aperiam quisquam magni rerum facere.\",\n         \"provider_type\": \"oidc\",\n         \"redirect_uri\": \"https://auth.example.com/callback\",\n         \"token_endpoint\": \"https://oauth2.googleapis.com/token\",\n         \"userinfo_endpoint\": \"https://openidconnect.googleapis.com/v1/userinfo\"\n      }\n   }'")
		}
		if body.Provider == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("provider", "body"))
		}
		if body.Provider != nil {
			if err2 := ValidateCreateIdentityProviderRequestRequestBody(body.Provider); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if ssoCreateIdentityProviderJWT != "" {
			jwt = &ssoCreateIdentityProviderJWT
		}
	}
	v := &sso.CreateIdentityProviderPayload{
		OrganizationID: body.OrganizationID,
	}
	if body.Provider != nil {
		v.Provider = marshalCreateIdentityProviderRequestRequestBodyToSsoCreateIdentityProviderRequest(body.Provider)
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetIdentityProviderPayload builds the payload for the sso
// get_identity_provider endpoint from CLI flags.
func BuildGetIdentityProviderPayload(ssoGetIdentityProviderID string, ssoGetIdentityProviderJWT string) (*sso.GetIdentityProviderPayload, error) {
	var id string
	{
		id = ssoGetIdentityProviderID
	}
	var jwt *string
	{
		if ssoGetIdentityProviderJWT != "" {
			jwt = &ssoGetIdentityProviderJWT
		}
	}
	v := &sso.GetIdentityProviderPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildUpdateIdentityProviderPayload builds the payload for the sso
// update_identity_provider endpoint from CLI flags.
func BuildUpdateIdentityProviderPayload(ssoUpdateIdentityProviderBody string, ssoUpdateIdentityProviderID string, ssoUpdateIdentityProviderJWT string) (*sso.UpdateIdentityProviderPayload, error) {
	var err error
	var body UpdateIdentityProviderRequestBody
	{
		err = json.Unmarshal([]byte(ssoUpdateIdentityProviderBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"provider\": {\n         \"active\": false,\n         \"attributes_mapping\": {\n            \"Doloribus fuga officia facere consectetur.\": \"Amet veniam illum iure ut corporis.\",\n            \"Est est et illum quidem veniam.\": \"In ut quibusdam labore sapiente ut.\",\n            \"Ipsum nesciunt et non.\": \"Quod quo voluptas odit ex recusandae dolorem.\"\n         },\n         \"authorization_endpoint\": \"Id accusamus maxime voluptatibus.\",\n         \"certificate\": \"Optio aut eos earum ea.\",\n         \"client_id\": \"Voluptas molestiae accusamus voluptas dolores.\",\n         \"client_secret\": \"Odit nesciunt iusto ut.\",\n         \"domains\": [\n            \"Praesentium architecto libero accusamus est voluptatem.\",\n            \"Neque aspernatur.\"\n         ],\n         \"issuer\": \"Ut ratione error quae quaerat hic ipsam.\",\n         \"jwks_uri\": \"Omnis incidunt sequi qui est.\",\n         \"metadata_url\": \"Officiis quisquam at quas facere quis inventore.\",\n         \"name\": \"Architecto in quisquam nihil nisi.\",\n         \"primary\": true,\n         \"private_key\": \"Voluptatem voluptatibus enim rerum.\",\n         \"redirect_uri\": \"Temporibus et autem aut reprehenderit unde.\",\n         \"token_endpoint\": \"Quis illo sit sint.\",\n         \"userinfo_endpoint\": \"Laborum et iusto.\"\n      }\n   }'")
		}
		if body.Provider == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("provider", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = ssoUpdateIdentityProviderID
	}
	var jwt *string
	{
		if ssoUpdateIdentityProviderJWT != "" {
			jwt = &ssoUpdateIdentityProviderJWT
		}
	}
	v := &sso.UpdateIdentityProviderPayload{}
	if body.Provider != nil {
		v.Provider = marshalUpdateIdentityProviderRequestRequestBodyToSsoUpdateIdentityProviderRequest(body.Provider)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeleteIdentityProviderPayload builds the payload for the sso
// delete_identity_provider endpoint from CLI flags.
func BuildDeleteIdentityProviderPayload(ssoDeleteIdentityProviderID string, ssoDeleteIdentityProviderJWT string) (*sso.DeleteIdentityProviderPayload, error) {
	var id string
	{
		id = ssoDeleteIdentityProviderID
	}
	var jwt *string
	{
		if ssoDeleteIdentityProviderJWT != "" {
			jwt = &ssoDeleteIdentityProviderJWT
		}
	}
	v := &sso.DeleteIdentityProviderPayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildSamlMetadataPayload builds the payload for the sso saml_metadata
// endpoint from CLI flags.
func BuildSamlMetadataPayload(ssoSamlMetadataID string) (*sso.SamlMetadataPayload, error) {
	var id string
	{
		id = ssoSamlMetadataID
	}
	v := &sso.SamlMetadataPayload{}
	v.ID = id

	return v, nil
}

// BuildSamlAcsPayload builds the payload for the sso saml_acs endpoint from
// CLI flags.
func BuildSamlAcsPayload(ssoSamlAcsID string) (*sso.SamlAcsPayload, error) {
	var id string
	{
		id = ssoSamlAcsID
	}
	v := &sso.SamlAcsPayload{}
	v.ID = id

	return v, nil
}
