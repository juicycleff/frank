// Code generated by goa v3.20.0, DO NOT EDIT.
//
// passwordless HTTP client encoders and decoders
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
	passwordless "github.com/juicycleff/frank/gen/passwordless"
	goahttp "goa.design/goa/v3/http"
)

// BuildEmailRequest instantiates a HTTP request object with method and path
// set to call the "passwordless" service "email" endpoint
func (c *Client) BuildEmailRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: EmailPasswordlessPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("passwordless", "email", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeEmailRequest returns an encoder for requests sent to the passwordless
// email server.
func EncodeEmailRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*passwordless.EmailPayload)
		if !ok {
			return goahttp.ErrInvalidType("passwordless", "email", "*passwordless.EmailPayload", v)
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
		body := NewEmailRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("passwordless", "email", err)
		}
		return nil
	}
}

// DecodeEmailResponse returns a decoder for responses returned by the
// passwordless email endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeEmailResponse may return the following errors:
//   - "bad_request" (type *passwordless.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *passwordless.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *passwordless.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *passwordless.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *passwordless.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeEmailResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body EmailResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "email", err)
			}
			err = ValidateEmailResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "email", err)
			}
			res := NewEmailResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body EmailBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "email", err)
			}
			err = ValidateEmailBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "email", err)
			}
			return nil, NewEmailBadRequest(&body)
		case http.StatusForbidden:
			var (
				body EmailForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "email", err)
			}
			err = ValidateEmailForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "email", err)
			}
			return nil, NewEmailForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body EmailInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "email", err)
			}
			err = ValidateEmailInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "email", err)
			}
			return nil, NewEmailInternalError(&body)
		case http.StatusNotFound:
			var (
				body EmailNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "email", err)
			}
			err = ValidateEmailNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "email", err)
			}
			return nil, NewEmailNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body EmailUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "email", err)
			}
			err = ValidateEmailUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "email", err)
			}
			return nil, NewEmailUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("passwordless", "email", resp.StatusCode, string(body))
		}
	}
}

// BuildSmsRequest instantiates a HTTP request object with method and path set
// to call the "passwordless" service "sms" endpoint
func (c *Client) BuildSmsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SmsPasswordlessPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("passwordless", "sms", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSmsRequest returns an encoder for requests sent to the passwordless
// sms server.
func EncodeSmsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*passwordless.SmsPayload)
		if !ok {
			return goahttp.ErrInvalidType("passwordless", "sms", "*passwordless.SmsPayload", v)
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
		body := NewSmsRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("passwordless", "sms", err)
		}
		return nil
	}
}

// DecodeSmsResponse returns a decoder for responses returned by the
// passwordless sms endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSmsResponse may return the following errors:
//   - "bad_request" (type *passwordless.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *passwordless.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *passwordless.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *passwordless.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *passwordless.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeSmsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body SmsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "sms", err)
			}
			err = ValidateSmsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "sms", err)
			}
			res := NewSmsResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body SmsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "sms", err)
			}
			err = ValidateSmsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "sms", err)
			}
			return nil, NewSmsBadRequest(&body)
		case http.StatusForbidden:
			var (
				body SmsForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "sms", err)
			}
			err = ValidateSmsForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "sms", err)
			}
			return nil, NewSmsForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body SmsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "sms", err)
			}
			err = ValidateSmsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "sms", err)
			}
			return nil, NewSmsInternalError(&body)
		case http.StatusNotFound:
			var (
				body SmsNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "sms", err)
			}
			err = ValidateSmsNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "sms", err)
			}
			return nil, NewSmsNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body SmsUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "sms", err)
			}
			err = ValidateSmsUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "sms", err)
			}
			return nil, NewSmsUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("passwordless", "sms", resp.StatusCode, string(body))
		}
	}
}

// BuildVerifyRequest instantiates a HTTP request object with method and path
// set to call the "passwordless" service "verify" endpoint
func (c *Client) BuildVerifyRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: VerifyPasswordlessPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("passwordless", "verify", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeVerifyRequest returns an encoder for requests sent to the passwordless
// verify server.
func EncodeVerifyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*passwordless.VerifyPayload)
		if !ok {
			return goahttp.ErrInvalidType("passwordless", "verify", "*passwordless.VerifyPayload", v)
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
		body := NewVerifyRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("passwordless", "verify", err)
		}
		return nil
	}
}

// DecodeVerifyResponse returns a decoder for responses returned by the
// passwordless verify endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeVerifyResponse may return the following errors:
//   - "bad_request" (type *passwordless.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *passwordless.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *passwordless.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *passwordless.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *passwordless.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeVerifyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body VerifyResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "verify", err)
			}
			err = ValidateVerifyResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "verify", err)
			}
			res := NewVerifyResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body VerifyBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "verify", err)
			}
			err = ValidateVerifyBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "verify", err)
			}
			return nil, NewVerifyBadRequest(&body)
		case http.StatusForbidden:
			var (
				body VerifyForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "verify", err)
			}
			err = ValidateVerifyForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "verify", err)
			}
			return nil, NewVerifyForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body VerifyInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "verify", err)
			}
			err = ValidateVerifyInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "verify", err)
			}
			return nil, NewVerifyInternalError(&body)
		case http.StatusNotFound:
			var (
				body VerifyNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "verify", err)
			}
			err = ValidateVerifyNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "verify", err)
			}
			return nil, NewVerifyNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body VerifyUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "verify", err)
			}
			err = ValidateVerifyUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "verify", err)
			}
			return nil, NewVerifyUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("passwordless", "verify", resp.StatusCode, string(body))
		}
	}
}

// BuildMethodsRequest instantiates a HTTP request object with method and path
// set to call the "passwordless" service "methods" endpoint
func (c *Client) BuildMethodsRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MethodsPasswordlessPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("passwordless", "methods", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMethodsRequest returns an encoder for requests sent to the
// passwordless methods server.
func EncodeMethodsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*passwordless.MethodsPayload)
		if !ok {
			return goahttp.ErrInvalidType("passwordless", "methods", "*passwordless.MethodsPayload", v)
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

// DecodeMethodsResponse returns a decoder for responses returned by the
// passwordless methods endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeMethodsResponse may return the following errors:
//   - "bad_request" (type *passwordless.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *passwordless.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *passwordless.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *passwordless.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *passwordless.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeMethodsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body MethodsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "methods", err)
			}
			err = ValidateMethodsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "methods", err)
			}
			res := NewMethodsResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body MethodsBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "methods", err)
			}
			err = ValidateMethodsBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "methods", err)
			}
			return nil, NewMethodsBadRequest(&body)
		case http.StatusForbidden:
			var (
				body MethodsForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "methods", err)
			}
			err = ValidateMethodsForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "methods", err)
			}
			return nil, NewMethodsForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body MethodsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "methods", err)
			}
			err = ValidateMethodsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "methods", err)
			}
			return nil, NewMethodsInternalError(&body)
		case http.StatusNotFound:
			var (
				body MethodsNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "methods", err)
			}
			err = ValidateMethodsNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "methods", err)
			}
			return nil, NewMethodsNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body MethodsUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "methods", err)
			}
			err = ValidateMethodsUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "methods", err)
			}
			return nil, NewMethodsUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("passwordless", "methods", resp.StatusCode, string(body))
		}
	}
}

// BuildMagicLinkRequest instantiates a HTTP request object with method and
// path set to call the "passwordless" service "magic_link" endpoint
func (c *Client) BuildMagicLinkRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MagicLinkPasswordlessPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("passwordless", "magic_link", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMagicLinkRequest returns an encoder for requests sent to the
// passwordless magic_link server.
func EncodeMagicLinkRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*passwordless.MagicLinkPayload)
		if !ok {
			return goahttp.ErrInvalidType("passwordless", "magic_link", "*passwordless.MagicLinkPayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewMagicLinkRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("passwordless", "magic_link", err)
		}
		return nil
	}
}

// DecodeMagicLinkResponse returns a decoder for responses returned by the
// passwordless magic_link endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeMagicLinkResponse may return the following errors:
//   - "bad_request" (type *passwordless.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *passwordless.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *passwordless.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *passwordless.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *passwordless.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeMagicLinkResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body MagicLinkResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "magic_link", err)
			}
			err = ValidateMagicLinkResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "magic_link", err)
			}
			res := NewMagicLinkResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body MagicLinkBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "magic_link", err)
			}
			err = ValidateMagicLinkBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "magic_link", err)
			}
			return nil, NewMagicLinkBadRequest(&body)
		case http.StatusForbidden:
			var (
				body MagicLinkForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "magic_link", err)
			}
			err = ValidateMagicLinkForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "magic_link", err)
			}
			return nil, NewMagicLinkForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body MagicLinkInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "magic_link", err)
			}
			err = ValidateMagicLinkInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "magic_link", err)
			}
			return nil, NewMagicLinkInternalError(&body)
		case http.StatusNotFound:
			var (
				body MagicLinkNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "magic_link", err)
			}
			err = ValidateMagicLinkNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "magic_link", err)
			}
			return nil, NewMagicLinkNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body MagicLinkUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("passwordless", "magic_link", err)
			}
			err = ValidateMagicLinkUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("passwordless", "magic_link", err)
			}
			return nil, NewMagicLinkUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("passwordless", "magic_link", resp.StatusCode, string(body))
		}
	}
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
		ID:              *v.ID,
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
