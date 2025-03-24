// Code generated by goa v3.20.0, DO NOT EDIT.
//
// oauth_client HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	designtypes "github.com/juicycleff/frank/gen/designtypes"
	oauthclient "github.com/juicycleff/frank/gen/oauth_client"
	goahttp "goa.design/goa/v3/http"
)

// BuildListProvidersRequest instantiates a HTTP request object with method and
// path set to call the "oauth_client" service "list_providers" endpoint
func (c *Client) BuildListProvidersRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListProvidersOauthClientPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("oauth_client", "list_providers", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListProvidersRequest returns an encoder for requests sent to the
// oauth_client list_providers server.
func EncodeListProvidersRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*oauthclient.ListProvidersPayload)
		if !ok {
			return goahttp.ErrInvalidType("oauth_client", "list_providers", "*oauthclient.ListProvidersPayload", v)
		}
		if p.Oauth2 != nil {
			head := *p.Oauth2
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.XAPIKey != nil {
			head := *p.XAPIKey
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeListProvidersResponse returns a decoder for responses returned by the
// oauth_client list_providers endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeListProvidersResponse may return the following errors:
//   - "bad_request" (type *oauthclient.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *oauthclient.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *oauthclient.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *oauthclient.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *oauthclient.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeListProvidersResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListProvidersResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "list_providers", err)
			}
			err = ValidateListProvidersResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "list_providers", err)
			}
			res := NewListProvidersResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListProvidersBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "list_providers", err)
			}
			err = ValidateListProvidersBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "list_providers", err)
			}
			return nil, NewListProvidersBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ListProvidersForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "list_providers", err)
			}
			err = ValidateListProvidersForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "list_providers", err)
			}
			return nil, NewListProvidersForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body ListProvidersInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "list_providers", err)
			}
			err = ValidateListProvidersInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "list_providers", err)
			}
			return nil, NewListProvidersInternalError(&body)
		case http.StatusNotFound:
			var (
				body ListProvidersNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "list_providers", err)
			}
			err = ValidateListProvidersNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "list_providers", err)
			}
			return nil, NewListProvidersNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ListProvidersUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "list_providers", err)
			}
			err = ValidateListProvidersUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "list_providers", err)
			}
			return nil, NewListProvidersUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("oauth_client", "list_providers", resp.StatusCode, string(body))
		}
	}
}

// BuildProviderAuthRequest instantiates a HTTP request object with method and
// path set to call the "oauth_client" service "provider_auth" endpoint
func (c *Client) BuildProviderAuthRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		provider string
	)
	{
		p, ok := v.(*oauthclient.ProviderAuthPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("oauth_client", "provider_auth", "*oauthclient.ProviderAuthPayload", v)
		}
		provider = p.Provider
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ProviderAuthOauthClientPath(provider)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("oauth_client", "provider_auth", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeProviderAuthRequest returns an encoder for requests sent to the
// oauth_client provider_auth server.
func EncodeProviderAuthRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*oauthclient.ProviderAuthPayload)
		if !ok {
			return goahttp.ErrInvalidType("oauth_client", "provider_auth", "*oauthclient.ProviderAuthPayload", v)
		}
		if p.Oauth2 != nil {
			head := *p.Oauth2
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.XAPIKey != nil {
			head := *p.XAPIKey
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		if p.RedirectURI != nil {
			values.Add("redirect_uri", *p.RedirectURI)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeProviderAuthResponse returns a decoder for responses returned by the
// oauth_client provider_auth endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeProviderAuthResponse may return the following errors:
//   - "bad_request" (type *oauthclient.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *oauthclient.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *oauthclient.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *oauthclient.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *oauthclient.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeProviderAuthResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusFound:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body ProviderAuthBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_auth", err)
			}
			err = ValidateProviderAuthBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_auth", err)
			}
			return nil, NewProviderAuthBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ProviderAuthForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_auth", err)
			}
			err = ValidateProviderAuthForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_auth", err)
			}
			return nil, NewProviderAuthForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body ProviderAuthInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_auth", err)
			}
			err = ValidateProviderAuthInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_auth", err)
			}
			return nil, NewProviderAuthInternalError(&body)
		case http.StatusNotFound:
			var (
				body ProviderAuthNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_auth", err)
			}
			err = ValidateProviderAuthNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_auth", err)
			}
			return nil, NewProviderAuthNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ProviderAuthUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_auth", err)
			}
			err = ValidateProviderAuthUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_auth", err)
			}
			return nil, NewProviderAuthUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("oauth_client", "provider_auth", resp.StatusCode, string(body))
		}
	}
}

// BuildProviderCallbackRequest instantiates a HTTP request object with method
// and path set to call the "oauth_client" service "provider_callback" endpoint
func (c *Client) BuildProviderCallbackRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		provider string
	)
	{
		p, ok := v.(*oauthclient.ProviderCallbackPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("oauth_client", "provider_callback", "*oauthclient.ProviderCallbackPayload", v)
		}
		provider = p.Provider
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ProviderCallbackOauthClientPath(provider)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("oauth_client", "provider_callback", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeProviderCallbackRequest returns an encoder for requests sent to the
// oauth_client provider_callback server.
func EncodeProviderCallbackRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*oauthclient.ProviderCallbackPayload)
		if !ok {
			return goahttp.ErrInvalidType("oauth_client", "provider_callback", "*oauthclient.ProviderCallbackPayload", v)
		}
		if p.Oauth2 != nil {
			head := *p.Oauth2
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.XAPIKey != nil {
			head := *p.XAPIKey
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		if p.Code != nil {
			values.Add("code", *p.Code)
		}
		if p.State != nil {
			values.Add("state", *p.State)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeProviderCallbackResponse returns a decoder for responses returned by
// the oauth_client provider_callback endpoint. restoreBody controls whether
// the response body should be restored after having been read.
// DecodeProviderCallbackResponse may return the following errors:
//   - "bad_request" (type *oauthclient.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *oauthclient.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *oauthclient.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *oauthclient.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *oauthclient.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeProviderCallbackResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ProviderCallbackResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_callback", err)
			}
			err = ValidateProviderCallbackResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_callback", err)
			}
			res := NewProviderCallbackResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ProviderCallbackBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_callback", err)
			}
			err = ValidateProviderCallbackBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_callback", err)
			}
			return nil, NewProviderCallbackBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ProviderCallbackForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_callback", err)
			}
			err = ValidateProviderCallbackForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_callback", err)
			}
			return nil, NewProviderCallbackForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body ProviderCallbackInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_callback", err)
			}
			err = ValidateProviderCallbackInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_callback", err)
			}
			return nil, NewProviderCallbackInternalError(&body)
		case http.StatusNotFound:
			var (
				body ProviderCallbackNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_callback", err)
			}
			err = ValidateProviderCallbackNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_callback", err)
			}
			return nil, NewProviderCallbackNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ProviderCallbackUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_client", "provider_callback", err)
			}
			err = ValidateProviderCallbackUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("oauth_client", "provider_callback", err)
			}
			return nil, NewProviderCallbackUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("oauth_client", "provider_callback", resp.StatusCode, string(body))
		}
	}
}

// unmarshalSSOProviderResponseBodyToOauthclientSSOProvider builds a value of
// type *oauthclient.SSOProvider from a value of type *SSOProviderResponseBody.
func unmarshalSSOProviderResponseBodyToOauthclientSSOProvider(v *SSOProviderResponseBody) *oauthclient.SSOProvider {
	res := &oauthclient.SSOProvider{
		ID:      *v.ID,
		Name:    *v.Name,
		Type:    *v.Type,
		IconURL: v.IconURL,
	}

	return res
}

// unmarshalUserResponseBodyToDesigntypesUser builds a value of type
// *designtypes.User from a value of type *UserResponseBody.
func unmarshalUserResponseBodyToDesigntypesUser(v *UserResponseBody) *designtypes.User {
	if v == nil {
		return nil
	}
	res := &designtypes.User{
		Active:          *v.Active,
		EmailVerified:   *v.EmailVerified,
		PhoneVerified:   v.PhoneVerified,
		ProfileImageURL: v.ProfileImageURL,
		FirstName:       v.FirstName,
		LastName:        v.LastName,
		PhoneNumber:     v.PhoneNumber,
		Email:           *v.Email,
	}
	if v.Locale != nil {
		res.Locale = *v.Locale
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}
	if v.Locale == nil {
		res.Locale = "en"
	}

	return res
}
