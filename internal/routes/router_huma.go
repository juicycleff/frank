package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/juicycleff/frank/config"
	"github.com/juicycleff/frank/internal/handlers"
	middleware3 "github.com/juicycleff/frank/internal/middleware"
	"github.com/juicycleff/frank/internal/services"
	"github.com/juicycleff/frank/pkg/data"
	"github.com/juicycleff/frank/pkg/logging"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/swag"
)

// HumaRouter manages HTTP routing
type HumaRouter struct {
	svcs    *services.Services
	router  chi.Router
	config  *config.Config
	logger  logging.Logger
	faktory *Factory
	huma    huma.API
}

// NewHumaRouter creates a new router
func NewHumaRouter(clients *data.Clients, svcs *services.Services, cfg *config.Config, logger logging.Logger) *HumaRouter {
	// Create Chi r
	r := chi.NewRouter()

	// Add built-in Chi middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// Add custom middleware
	r.Use(middleware3.Logging(logger))
	r.Use(logging.Middleware)
	r.Use(middleware3.RequestLogging)

	// Add rate limiter if enabled
	if cfg.Security.RateLimitEnabled {
		r.Use(middleware3.RateLimiter(cfg.Security.RateLimitPerSecond, cfg.Security.RateLimitBurst))
	}

	// Add recovery middleware
	r.Use(middleware3.Recovery(logger))
	r.Use(middleware3.ErrorHandler(logger))

	// CORS middleware
	r.Use(middleware3.CORS(cfg))

	// Rate limiting middleware if enabled
	if cfg.Security.RateLimitEnabled {
		r.Use(middleware3.RateLimiter(cfg.Security.RateLimitPerSecond, cfg.Security.RateLimitBurst))
	}

	// Security headers middleware
	if cfg.Security.SecHeadersEnabled {
		// r.Use(customMiddleware.SecurityHeaders(cfg.Security))
	}

	// Wrap the router with Huma to create an API instance.
	api := humachi.New(r, huma.DefaultConfig(
		cfg.Server.Name,
		cfg.Version,
	))

	faktory := NewFactory(svcs, clients, cfg, logger)

	return &HumaRouter{
		router:  r,
		config:  cfg,
		logger:  logger,
		svcs:    svcs,
		faktory: faktory,
		huma:    api,
	}
}

// RegisterRoutes registers all API routes
func (rt *HumaRouter) RegisterRoutes() {
	faktory := rt.faktory

	// Health check routes
	faktory.Health.RegisterPublicRoutes(rt.router)

	// Apply Auth middleware - you might need to modify your Auth middleware to work with Chi
	// Since Chi uses middleware differently, the Auth middleware should be adapted
	authMw := middleware3.AuthHuma(rt.config, rt.logger, rt.svcs.Session, rt.svcs.APIKey)

	// Organization middleware (for routes that need organization context)
	// orgMiddleware := customMiddleware.NewOrganizationMiddleware(rt.config, rt.svcs.Organization, rt.logger)

	// API routes
	apiGroup := huma.NewGroup(rt.huma, "/api")
	v1Group := huma.NewGroup(apiGroup, "/v1")
	protectedGroup := huma.NewGroup(v1Group)

	// Apply auth middleware
	protectedGroup.UseMiddleware(authMw)
	protectedGroup.UseMiddleware(middleware3.RequireAuthenticationHuma)

	faktory.Auth.RegisterRoutesHuma(protectedGroup)

	rt.router.Route("/api", func(r chi.Router) {
		cfg := huma.DefaultConfig(rt.config.Server.Name, rt.config.Version)
		cfg.Servers = []*huma.Server{
			{URL: fmt.Sprintf("%s/%s", rt.config.GetServerAddress(), "/api"), Description: "Server"},
		}

		// Version-specific routes
		r.Route("/v1", func(r chi.Router) {
			// Public routes (no auth required)
			r.Group(func(r chi.Router) {
				// Auth routes (login, register, etc.)
				faktory.Auth.RegisterPublicRoutes(r)

				// OAuth2 provider endpoints (authorize, token, etc.)
				faktory.OAuth.RegisterProviderRoutes(r)

				// Passwordless auth endpoints
				faktory.Passwordless.RegisterPublicRoutes(r)

				// Passkey auth endpoints
				faktory.Passkey.RegisterPublicRoutes(r)

				// SSO endpoints
				faktory.SSO.RegisterPublicRoutes(r)
			})

			// // Protected routes (auth required)
			// r.Group(func(r chi.Router) {
			// 	// Apply auth middleware
			// 	r.Use(authMw)
			// 	r.Use(customMiddleware.RequireAuthentication)
			//
			// 	// Auth routes
			// 	faktory.Auth.RegisterRoutes(r)
			//
			// 	// User routes
			// 	faktory.User.RegisterRoutes(r)
			//
			// 	// MFA routes
			// 	faktory.MFA.RegisterRoutes(r)
			//
			// 	// API Key routes
			// 	faktory.APIKey.RegisterRoutes(r)
			//
			// 	// Organization routes (basic routes)
			// 	faktory.Organization.RegisterRoutes(r)
			//
			// 	// Webhooks routes
			// 	faktory.Webhook.RegisterRoutes(r)
			//
			// 	// OAuth client routes
			// 	faktory.OAuth.RegisterClientRoutes(r)
			//
			// 	// Passkey routes
			// 	faktory.Passkey.RegisterRoutes(r)
			//
			// 	// Passwordless routes
			// 	faktory.Passwordless.RegisterRoutes(r)
			//
			// 	// SSO routes
			// 	faktory.SSO.RegisterRoutes(r)
			//
			// 	// Roles routes
			// 	faktory.RBAC.RegisterRolesRoutes(r)
			//
			// 	// Permissions routes
			// 	faktory.RBAC.RegisterPermissionsRoutes(r)
			//
			// 	// Email template routes
			// 	faktory.Email.RegisterRoutes(r)
			// })
			//
			// // Organization-specific routes (require org context)
			// r.Group(func(r chi.Router) {
			// 	// Apply auth middleware plus organization middleware
			// 	r.Use(authMw)
			// 	r.Use(customMiddleware.RequireAuthentication)
			// 	r.Use(orgMiddleware.RequireOrganization)
			// 	r.Use(orgMiddleware.RequireOrganizationMember)
			//
			// 	// Organization management routes
			// 	faktory.Organization.RegisterOrganizationRoutes(r)
			//
			// 	// Feature-specific routes that require organization context
			// 	r.Group(func(r chi.Router) {
			// 		r.Use(orgMiddleware.RequireFeatureEnabled("webhooks"))
			// 		faktory.Webhook.RegisterOrganizationRoutes(r)
			// 	})
			//
			// 	r.Group(func(r chi.Router) {
			// 		r.Use(orgMiddleware.RequireFeatureEnabled("sso"))
			// 		faktory.SSO.RegisterOrganizationRoutes(r)
			// 	})
			// })

			// OAuth provider endpoints
			rt.router.Route("/oauth", func(r chi.Router) {
				r.Get("/authorize", handlers.OAuthAuthorize)
				r.Post("/token", handlers.OAuthToken)
				r.Get("/userinfo", handlers.OAuthUserInfo)
				r.Get("/.well-known/openid-configuration", handlers.OAuthConfiguration)
				r.Get("/.well-known/jwks.json", handlers.OAuthJWKS)
			})

			// SAML endpoints
			rt.router.Route("/saml", func(r chi.Router) {
				r.Post("/acs", handlers.SAMLAssertionConsumerService)
				r.Get("/metadata", handlers.SAMLMetadata)
			})

			// OIDC discovery endpoints
			rt.router.Get("/.well-known/openid-configuration", faktory.OAuth.OAuthConfiguration)
			rt.router.Get("/.well-known/jwks.json", faktory.OAuth.OAuthJWKS)

			// Webhook endpoints (for receiving notifications)
			rt.router.Post("/webhooks/{id}", handlers.ReceiveWebhook)
		})
	})

	// Swagger endpoint
	rt.router.Get("/docs/*", httpSwagger.WrapHandler)

	// Static assets (if needed)
	fileServer := http.FileServer(http.Dir("./web/static"))
	rt.router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	rt.router.Handle("/*", handlers.FileServer("./web/apps/client/dist", rt.router))
}

// Handler returns the HTTP handler
func (rt *HumaRouter) Handler() http.Handler {
	return rt.router
}

// Group adds a new route group
func (rt *HumaRouter) Group(fn func(r chi.Router)) {
	rt.router.Group(fn)
}

// Route adds a new route group with a pattern
func (rt *HumaRouter) Route(pattern string, fn func(r chi.Router)) {
	rt.router.Route(pattern, fn)
}

// Use appends a middleware to the chain
func (rt *HumaRouter) Use(middleware ...func(http.Handler) http.Handler) {
	rt.router.Use(middleware...)
}

// Method adds a method-specific route
func (rt *HumaRouter) Method(method, pattern string, handler http.HandlerFunc) {
	rt.router.Method(method, pattern, handler)
}

// NotFound sets the not found handler
func (rt *HumaRouter) NotFound(handler http.HandlerFunc) {
	rt.router.NotFound(handler)
}

// MethodNotAllowed sets the method not allowed handler
func (rt *HumaRouter) MethodNotAllowed(handler http.HandlerFunc) {
	rt.router.MethodNotAllowed(handler)
}
