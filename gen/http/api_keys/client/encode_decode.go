// Code generated by goa v3.20.0, DO NOT EDIT.
//
// api_keys HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	apikeys "github.com/juicycleff/frank/gen/api_keys"
	designtypes "github.com/juicycleff/frank/gen/designtypes"
	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "api_keys" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListAPIKeysPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_keys", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the api_keys list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*apikeys.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("api_keys", "list", "*apikeys.ListPayload", v)
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
		values.Add("offset", fmt.Sprintf("%v", p.Offset))
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		if p.Type != nil {
			values.Add("type", *p.Type)
		}
		if p.OrganizationID != nil {
			values.Add("organization_id", *p.OrganizationID)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the api_keys
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListResponse may return the following errors:
//   - "bad_request" (type *apikeys.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *apikeys.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *apikeys.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *apikeys.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *apikeys.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "list", err)
			}
			res := NewListResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ListBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "list", err)
			}
			err = ValidateListBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "list", err)
			}
			return nil, NewListBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ListForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "list", err)
			}
			err = ValidateListForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "list", err)
			}
			return nil, NewListForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body ListInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "list", err)
			}
			err = ValidateListInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "list", err)
			}
			return nil, NewListInternalError(&body)
		case http.StatusNotFound:
			var (
				body ListNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "list", err)
			}
			err = ValidateListNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "list", err)
			}
			return nil, NewListNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ListUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "list", err)
			}
			err = ValidateListUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "list", err)
			}
			return nil, NewListUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_keys", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateRequest instantiates a HTTP request object with method and path
// set to call the "api_keys" service "create" endpoint
func (c *Client) BuildCreateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAPIKeysPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_keys", "create", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateRequest returns an encoder for requests sent to the api_keys
// create server.
func EncodeCreateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*apikeys.CreatePayload)
		if !ok {
			return goahttp.ErrInvalidType("api_keys", "create", "*apikeys.CreatePayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("api_keys", "create", err)
		}
		return nil
	}
}

// DecodeCreateResponse returns a decoder for responses returned by the
// api_keys create endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateResponse may return the following errors:
//   - "bad_request" (type *apikeys.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *apikeys.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *apikeys.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *apikeys.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *apikeys.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeCreateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusCreated:
			var (
				body CreateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "create", err)
			}
			err = ValidateCreateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "create", err)
			}
			res := NewCreateAPIKeyWithSecretResponseCreated(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "create", err)
			}
			err = ValidateCreateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "create", err)
			}
			return nil, NewCreateBadRequest(&body)
		case http.StatusForbidden:
			var (
				body CreateForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "create", err)
			}
			err = ValidateCreateForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "create", err)
			}
			return nil, NewCreateForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body CreateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "create", err)
			}
			err = ValidateCreateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "create", err)
			}
			return nil, NewCreateInternalError(&body)
		case http.StatusNotFound:
			var (
				body CreateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "create", err)
			}
			err = ValidateCreateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "create", err)
			}
			return nil, NewCreateNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body CreateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "create", err)
			}
			err = ValidateCreateUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "create", err)
			}
			return nil, NewCreateUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_keys", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "api_keys" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*apikeys.GetPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("api_keys", "get", "*apikeys.GetPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAPIKeysPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_keys", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the api_keys get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*apikeys.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("api_keys", "get", "*apikeys.GetPayload", v)
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

// DecodeGetResponse returns a decoder for responses returned by the api_keys
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetResponse may return the following errors:
//   - "bad_request" (type *apikeys.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *apikeys.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *apikeys.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *apikeys.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *apikeys.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "get", err)
			}
			err = ValidateGetResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "get", err)
			}
			res := NewGetAPIKeyResponseOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body GetBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "get", err)
			}
			err = ValidateGetBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "get", err)
			}
			return nil, NewGetBadRequest(&body)
		case http.StatusForbidden:
			var (
				body GetForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "get", err)
			}
			err = ValidateGetForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "get", err)
			}
			return nil, NewGetForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body GetInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "get", err)
			}
			err = ValidateGetInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "get", err)
			}
			return nil, NewGetInternalError(&body)
		case http.StatusNotFound:
			var (
				body GetNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "get", err)
			}
			err = ValidateGetNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "get", err)
			}
			return nil, NewGetNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body GetUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "get", err)
			}
			err = ValidateGetUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "get", err)
			}
			return nil, NewGetUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_keys", "get", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateRequest instantiates a HTTP request object with method and path
// set to call the "api_keys" service "update" endpoint
func (c *Client) BuildUpdateRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*apikeys.UpdatePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("api_keys", "update", "*apikeys.UpdatePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateAPIKeysPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_keys", "update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateRequest returns an encoder for requests sent to the api_keys
// update server.
func EncodeUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*apikeys.UpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("api_keys", "update", "*apikeys.UpdatePayload", v)
		}
		if p.JWT != nil {
			head := *p.JWT
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("api_keys", "update", err)
		}
		return nil
	}
}

// DecodeUpdateResponse returns a decoder for responses returned by the
// api_keys update endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateResponse may return the following errors:
//   - "bad_request" (type *apikeys.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *apikeys.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *apikeys.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *apikeys.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *apikeys.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body UpdateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "update", err)
			}
			err = ValidateUpdateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "update", err)
			}
			res := NewUpdateAPIKeyResponseOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body UpdateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "update", err)
			}
			err = ValidateUpdateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "update", err)
			}
			return nil, NewUpdateBadRequest(&body)
		case http.StatusForbidden:
			var (
				body UpdateForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "update", err)
			}
			err = ValidateUpdateForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "update", err)
			}
			return nil, NewUpdateForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body UpdateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "update", err)
			}
			err = ValidateUpdateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "update", err)
			}
			return nil, NewUpdateInternalError(&body)
		case http.StatusNotFound:
			var (
				body UpdateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "update", err)
			}
			err = ValidateUpdateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "update", err)
			}
			return nil, NewUpdateNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body UpdateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "update", err)
			}
			err = ValidateUpdateUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "update", err)
			}
			return nil, NewUpdateUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_keys", "update", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteRequest instantiates a HTTP request object with method and path
// set to call the "api_keys" service "delete" endpoint
func (c *Client) BuildDeleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*apikeys.DeletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("api_keys", "delete", "*apikeys.DeletePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteAPIKeysPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_keys", "delete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteRequest returns an encoder for requests sent to the api_keys
// delete server.
func EncodeDeleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*apikeys.DeletePayload)
		if !ok {
			return goahttp.ErrInvalidType("api_keys", "delete", "*apikeys.DeletePayload", v)
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

// DecodeDeleteResponse returns a decoder for responses returned by the
// api_keys delete endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteResponse may return the following errors:
//   - "bad_request" (type *apikeys.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *apikeys.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *apikeys.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *apikeys.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *apikeys.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeDeleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body DeleteBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "delete", err)
			}
			err = ValidateDeleteBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "delete", err)
			}
			return nil, NewDeleteBadRequest(&body)
		case http.StatusForbidden:
			var (
				body DeleteForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "delete", err)
			}
			err = ValidateDeleteForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "delete", err)
			}
			return nil, NewDeleteForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body DeleteInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "delete", err)
			}
			err = ValidateDeleteInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "delete", err)
			}
			return nil, NewDeleteInternalError(&body)
		case http.StatusNotFound:
			var (
				body DeleteNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "delete", err)
			}
			err = ValidateDeleteNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "delete", err)
			}
			return nil, NewDeleteNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body DeleteUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "delete", err)
			}
			err = ValidateDeleteUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "delete", err)
			}
			return nil, NewDeleteUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_keys", "delete", resp.StatusCode, string(body))
		}
	}
}

// BuildValidateRequest instantiates a HTTP request object with method and path
// set to call the "api_keys" service "validate" endpoint
func (c *Client) BuildValidateRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ValidateAPIKeysPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_keys", "validate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeValidateRequest returns an encoder for requests sent to the api_keys
// validate server.
func EncodeValidateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*apikeys.ValidatePayload)
		if !ok {
			return goahttp.ErrInvalidType("api_keys", "validate", "*apikeys.ValidatePayload", v)
		}
		values := req.URL.Query()
		values.Add("api_key", p.APIKey)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeValidateResponse returns a decoder for responses returned by the
// api_keys validate endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeValidateResponse may return the following errors:
//   - "bad_request" (type *apikeys.BadRequestError): http.StatusBadRequest
//   - "forbidden" (type *apikeys.ForbiddenError): http.StatusForbidden
//   - "internal_error" (type *apikeys.InternalServerError): http.StatusInternalServerError
//   - "not_found" (type *apikeys.NotFoundError): http.StatusNotFound
//   - "unauthorized" (type *apikeys.UnauthorizedError): http.StatusUnauthorized
//   - error: internal error
func DecodeValidateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
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
				body ValidateResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "validate", err)
			}
			err = ValidateValidateResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "validate", err)
			}
			res := NewValidateResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body ValidateBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "validate", err)
			}
			err = ValidateValidateBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "validate", err)
			}
			return nil, NewValidateBadRequest(&body)
		case http.StatusForbidden:
			var (
				body ValidateForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "validate", err)
			}
			err = ValidateValidateForbiddenResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "validate", err)
			}
			return nil, NewValidateForbidden(&body)
		case http.StatusInternalServerError:
			var (
				body ValidateInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "validate", err)
			}
			err = ValidateValidateInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "validate", err)
			}
			return nil, NewValidateInternalError(&body)
		case http.StatusNotFound:
			var (
				body ValidateNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "validate", err)
			}
			err = ValidateValidateNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "validate", err)
			}
			return nil, NewValidateNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ValidateUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("api_keys", "validate", err)
			}
			err = ValidateValidateUnauthorizedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("api_keys", "validate", err)
			}
			return nil, NewValidateUnauthorized(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_keys", "validate", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAPIKeyResponseResponseBodyToApikeysAPIKeyResponse builds a value of
// type *apikeys.APIKeyResponse from a value of type
// *APIKeyResponseResponseBody.
func unmarshalAPIKeyResponseResponseBodyToApikeysAPIKeyResponse(v *APIKeyResponseResponseBody) *apikeys.APIKeyResponse {
	res := &apikeys.APIKeyResponse{
		ID:             *v.ID,
		Name:           *v.Name,
		UserID:         v.UserID,
		OrganizationID: v.OrganizationID,
		Type:           *v.Type,
		Active:         *v.Active,
		LastUsed:       v.LastUsed,
		ExpiresAt:      v.ExpiresAt,
		CreatedAt:      *v.CreatedAt,
		UpdatedAt:      v.UpdatedAt,
	}
	if v.Permissions != nil {
		res.Permissions = make([]string, len(v.Permissions))
		for i, val := range v.Permissions {
			res.Permissions[i] = val
		}
	}
	if v.Scopes != nil {
		res.Scopes = make([]string, len(v.Scopes))
		for i, val := range v.Scopes {
			res.Scopes[i] = val
		}
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}

// unmarshalPaginationResponseBodyToDesigntypesPagination builds a value of
// type *designtypes.Pagination from a value of type *PaginationResponseBody.
func unmarshalPaginationResponseBodyToDesigntypesPagination(v *PaginationResponseBody) *designtypes.Pagination {
	res := &designtypes.Pagination{
		Offset:      *v.Offset,
		Limit:       *v.Limit,
		Total:       *v.Total,
		TotalPages:  *v.TotalPages,
		CurrentPage: *v.CurrentPage,
		HasNext:     *v.HasNext,
		HasPrevious: *v.HasPrevious,
	}

	return res
}

// marshalApikeysCreateAPIKeyRequestToCreateAPIKeyRequestRequestBody builds a
// value of type *CreateAPIKeyRequestRequestBody from a value of type
// *apikeys.CreateAPIKeyRequest.
func marshalApikeysCreateAPIKeyRequestToCreateAPIKeyRequestRequestBody(v *apikeys.CreateAPIKeyRequest) *CreateAPIKeyRequestRequestBody {
	res := &CreateAPIKeyRequestRequestBody{
		Name:      v.Name,
		Type:      v.Type,
		ExpiresIn: v.ExpiresIn,
	}
	{
		var zero string
		if res.Type == zero {
			res.Type = "client"
		}
	}
	if v.Permissions != nil {
		res.Permissions = make([]string, len(v.Permissions))
		for i, val := range v.Permissions {
			res.Permissions[i] = val
		}
	}
	if v.Scopes != nil {
		res.Scopes = make([]string, len(v.Scopes))
		for i, val := range v.Scopes {
			res.Scopes[i] = val
		}
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}

// marshalCreateAPIKeyRequestRequestBodyToApikeysCreateAPIKeyRequest builds a
// value of type *apikeys.CreateAPIKeyRequest from a value of type
// *CreateAPIKeyRequestRequestBody.
func marshalCreateAPIKeyRequestRequestBodyToApikeysCreateAPIKeyRequest(v *CreateAPIKeyRequestRequestBody) *apikeys.CreateAPIKeyRequest {
	res := &apikeys.CreateAPIKeyRequest{
		Name:      v.Name,
		Type:      v.Type,
		ExpiresIn: v.ExpiresIn,
	}
	{
		var zero string
		if res.Type == zero {
			res.Type = "client"
		}
	}
	if v.Permissions != nil {
		res.Permissions = make([]string, len(v.Permissions))
		for i, val := range v.Permissions {
			res.Permissions[i] = val
		}
	}
	if v.Scopes != nil {
		res.Scopes = make([]string, len(v.Scopes))
		for i, val := range v.Scopes {
			res.Scopes[i] = val
		}
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}

// marshalApikeysUpdateAPIKeyRequestToUpdateAPIKeyRequestRequestBody builds a
// value of type *UpdateAPIKeyRequestRequestBody from a value of type
// *apikeys.UpdateAPIKeyRequest.
func marshalApikeysUpdateAPIKeyRequestToUpdateAPIKeyRequestRequestBody(v *apikeys.UpdateAPIKeyRequest) *UpdateAPIKeyRequestRequestBody {
	res := &UpdateAPIKeyRequestRequestBody{
		Name:      v.Name,
		Active:    v.Active,
		ExpiresAt: v.ExpiresAt,
	}
	if v.Permissions != nil {
		res.Permissions = make([]string, len(v.Permissions))
		for i, val := range v.Permissions {
			res.Permissions[i] = val
		}
	}
	if v.Scopes != nil {
		res.Scopes = make([]string, len(v.Scopes))
		for i, val := range v.Scopes {
			res.Scopes[i] = val
		}
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}

// marshalUpdateAPIKeyRequestRequestBodyToApikeysUpdateAPIKeyRequest builds a
// value of type *apikeys.UpdateAPIKeyRequest from a value of type
// *UpdateAPIKeyRequestRequestBody.
func marshalUpdateAPIKeyRequestRequestBodyToApikeysUpdateAPIKeyRequest(v *UpdateAPIKeyRequestRequestBody) *apikeys.UpdateAPIKeyRequest {
	res := &apikeys.UpdateAPIKeyRequest{
		Name:      v.Name,
		Active:    v.Active,
		ExpiresAt: v.ExpiresAt,
	}
	if v.Permissions != nil {
		res.Permissions = make([]string, len(v.Permissions))
		for i, val := range v.Permissions {
			res.Permissions[i] = val
		}
	}
	if v.Scopes != nil {
		res.Scopes = make([]string, len(v.Scopes))
		for i, val := range v.Scopes {
			res.Scopes[i] = val
		}
	}
	if v.Metadata != nil {
		res.Metadata = make(map[string]any, len(v.Metadata))
		for key, val := range v.Metadata {
			tk := key
			tv := val
			res.Metadata[tk] = tv
		}
	}

	return res
}
