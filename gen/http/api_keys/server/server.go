// Code generated by goa v3.20.0, DO NOT EDIT.
//
// api_keys HTTP server
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"net/http"
	"regexp"

	apikeys "github.com/juicycleff/frank/gen/api_keys"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the api_keys service endpoint HTTP handlers.
type Server struct {
	Mounts   []*MountPoint
	List     http.Handler
	Create   http.Handler
	Get      http.Handler
	Update   http.Handler
	Delete   http.Handler
	Validate http.Handler
	CORS     http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the api_keys service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *apikeys.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"List", "GET", "/v1/api-keys"},
			{"Create", "POST", "/v1/api-keys"},
			{"Get", "GET", "/v1/api-keys/{id}"},
			{"Update", "PUT", "/v1/api-keys/{id}"},
			{"Delete", "DELETE", "/v1/api-keys/{id}"},
			{"Validate", "GET", "/v1/api-keys/validate"},
			{"CORS", "OPTIONS", "/v1/api-keys"},
			{"CORS", "OPTIONS", "/v1/api-keys/{id}"},
			{"CORS", "OPTIONS", "/v1/api-keys/validate"},
		},
		List:     NewListHandler(e.List, mux, decoder, encoder, errhandler, formatter),
		Create:   NewCreateHandler(e.Create, mux, decoder, encoder, errhandler, formatter),
		Get:      NewGetHandler(e.Get, mux, decoder, encoder, errhandler, formatter),
		Update:   NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		Delete:   NewDeleteHandler(e.Delete, mux, decoder, encoder, errhandler, formatter),
		Validate: NewValidateHandler(e.Validate, mux, decoder, encoder, errhandler, formatter),
		CORS:     NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "api_keys" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.List = m(s.List)
	s.Create = m(s.Create)
	s.Get = m(s.Get)
	s.Update = m(s.Update)
	s.Delete = m(s.Delete)
	s.Validate = m(s.Validate)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return apikeys.MethodNames[:] }

// Mount configures the mux to serve the api_keys endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListHandler(mux, h.List)
	MountCreateHandler(mux, h.Create)
	MountGetHandler(mux, h.Get)
	MountUpdateHandler(mux, h.Update)
	MountDeleteHandler(mux, h.Delete)
	MountValidateHandler(mux, h.Validate)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the api_keys endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListHandler configures the mux to serve the "api_keys" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAPIKeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/api-keys", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "api_keys" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListRequest(mux, decoder)
		encodeResponse = EncodeListResponse(encoder)
		encodeError    = EncodeListError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api_keys")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateHandler configures the mux to serve the "api_keys" service
// "create" endpoint.
func MountCreateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAPIKeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/api-keys", f)
}

// NewCreateHandler creates a HTTP handler which loads the HTTP request and
// calls the "api_keys" service "create" endpoint.
func NewCreateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateRequest(mux, decoder)
		encodeResponse = EncodeCreateResponse(encoder)
		encodeError    = EncodeCreateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api_keys")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetHandler configures the mux to serve the "api_keys" service "get"
// endpoint.
func MountGetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAPIKeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/api-keys/{id}", f)
}

// NewGetHandler creates a HTTP handler which loads the HTTP request and calls
// the "api_keys" service "get" endpoint.
func NewGetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetRequest(mux, decoder)
		encodeResponse = EncodeGetResponse(encoder)
		encodeError    = EncodeGetError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api_keys")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateHandler configures the mux to serve the "api_keys" service
// "update" endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAPIKeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/api-keys/{id}", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "api_keys" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = EncodeUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api_keys")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteHandler configures the mux to serve the "api_keys" service
// "delete" endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAPIKeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/api-keys/{id}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "api_keys" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, decoder)
		encodeResponse = EncodeDeleteResponse(encoder)
		encodeError    = EncodeDeleteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api_keys")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountValidateHandler configures the mux to serve the "api_keys" service
// "validate" endpoint.
func MountValidateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAPIKeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/api-keys/validate", f)
}

// NewValidateHandler creates a HTTP handler which loads the HTTP request and
// calls the "api_keys" service "validate" endpoint.
func NewValidateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeValidateRequest(mux, decoder)
		encodeResponse = EncodeValidateResponse(encoder)
		encodeError    = EncodeValidateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "validate")
		ctx = context.WithValue(ctx, goa.ServiceKey, "api_keys")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service api_keys.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleAPIKeysOrigin(h)
	mux.Handle("OPTIONS", "/v1/api-keys", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/api-keys/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/api-keys/validate", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleAPIKeysOrigin applies the CORS response headers corresponding to the
// origin for the service api_keys.
func HandleAPIKeysOrigin(h http.Handler) http.Handler {
	spec1 := regexp.MustCompile(".*localhost.*")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*.frank.com") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "100")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Headers", "X-Shared-Secret, X-Api-Version")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec1) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Request-Id")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "localhost") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
