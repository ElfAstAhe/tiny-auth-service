package rest

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	libconfig "github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libhttp "github.com/ElfAstAhe/go-service-template/pkg/transport/http"
	libmware "github.com/ElfAstAhe/go-service-template/pkg/transport/http/middleware"
	_ "github.com/ElfAstAhe/tiny-auth-service/docs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	appmware "github.com/ElfAstAhe/tiny-auth-service/internal/transport/rest/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hellofresh/health-go/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/riandyrn/otelchi"
	swagh "github.com/swaggo/http-swagger"
)

// AppChiRouter orchestrates the HTTP delivery layer, encapsulating the chi.Mux multiplexer matrix.
// It wires cross-cutting middleware pipelines, diagnostics endpoints, and handles structural domain route matching.
type AppChiRouter struct {
	router          *chi.Mux
	log             logger.Logger
	config          *config.Config
	health          *health.Health
	healthz         libhttp.HealthzFunc
	readyz          libhttp.ReadyzFunc
	authFacade      facade.AuthFacade
	userFacade      facade.UserFacade
	userAdminFacade facade.UserAdminFacade
	roleAdminFacade facade.RoleAdminFacade
}

// Compile-time interface compliance verification
var _ libhttp.Router = (*AppChiRouter)(nil)

// NewAppRouter instantiates a new AppChiRouter, processing incoming functional Options parameters and executing strict lifecycle evaluations.
func NewAppRouter(opts ...Option) (*AppChiRouter, error) {
	options := &AppRouterOptions{}

	for _, opt := range opts {
		opt(options)
	}

	err := options.Validate()
	if err != nil {
		return nil, err
	}

	return newAppChiRouter(
		options.Conf,
		options.Logger,
		options.AuthHelper,
		options.Health,
		options.Healthz,
		options.Readyz,
		options.AuthFacade,
		options.UserFacade,
		options.UserAdminFacade,
		options.RoleAdminFacade), nil
}

// newAppChiRouter executes the logical internal assembly sequence allocating routers, profiling hooks, metrics handles, and structural routing groups.
func newAppChiRouter(
	config *config.Config,
	logger logger.Logger,
	authHelper auth.Helper,
	health *health.Health,
	healthz libhttp.HealthzFunc,
	readyz libhttp.ReadyzFunc,
	authFacade facade.AuthFacade,
	userFacade facade.UserFacade,
	userAdminFacade facade.UserAdminFacade,
	roleAdminFacade facade.RoleAdminFacade,
) *AppChiRouter {
	res := &AppChiRouter{
		router:          chi.NewRouter(),
		log:             logger.GetLogger("app-chi-router"),
		config:          config,
		health:          health,
		healthz:         healthz,
		readyz:          readyz,
		authFacade:      authFacade,
		userFacade:      userFacade,
		userAdminFacade: userAdminFacade,
		roleAdminFacade: roleAdminFacade,
	}

	// setup middleware
	res.setupMiddleware(authHelper, logger)

	// mount debug
	res.router.Mount("/debug", middleware.Profiler())
	// mount swagger
	res.router.Mount("/swagger/", swagh.WrapHandler)
	// mount status
	res.router.Mount("/status", res.health.Handler())
	// mount metrics
	res.router.Mount("/metrics", promhttp.Handler())

	// setup routes
	res.setupRoutes()

	return res
}

// GetRouter yields the underlying compiled chi.Mux handler instance serving the web network connection loops.
func (cr *AppChiRouter) GetRouter() http.Handler {
	return cr.router
}

// setupMiddleware builds the global cascaded interceptor pipeline, applying telemetry, compression limits, and security boundaries.
func (cr *AppChiRouter) setupMiddleware(
	authHelper auth.Helper,
	logger logger.Logger,
) {
	// 1. CRITICAL: Recoverer must be placed at the absolute top of the stack to intercept panics from all downstream handlers.
	cr.router.Use(middleware.Recoverer)

	// 2. LOGGING: HTTPRequestLogger records every single inbound request, including those that panic or fail auth validation.
	cr.router.Use(libmware.NewHTTPRequestLogger(logger).Handle)

	// 3. TELEMETRY & OBSERVABILITY: OpenTelemetry and Prometheus catch the final HTTP execution states and write precise metrics.
	cr.router.Use(otelchi.Middleware(cr.config.Telemetry.ServiceName, otelchi.WithChiRoutes(cr.router)))
	cr.router.Use(libmware.MetricsMiddleware)

	// 4. REQUEST IDENTIFICATION: Correlation tokens propagation for end-to-end distributed tracking.
	cr.router.Use(middleware.RequestID)
	cr.router.Use(libmware.NewDefaultRequestIDExtractor().Handler)
	cr.router.Use(libmware.NewDefaultTraceIDExtractor().Handler)

	// 5. NETWORKING CORE: Real IP extraction boundaries.
	cr.router.Use(libmware.NewRealIPExtractor().Handler)
	// realIP
	//cr.router.Use(middleware.RealIP)

	// 6. RESOURCE BUDGETING & PROTECTION: Network stream limits and defensive runtime constraints.
	cr.router.Use(middleware.Timeout(cr.config.HTTP.ReadTimeout))
	cr.router.Use(libmware.NewCompress(logger,
		libhttp.MediaTypeApplicationJSON,
		libhttp.MediaTypeTextPlain,
	).Handle)
	cr.router.Use(libmware.NewDecompress(int64(cr.config.HTTP.MaxRequestBodySize), logger).Handle)

	// 7. SECURITY BOUNDARY: Authentication extraction executes right before routing to core application aggregates.
	// ROUTING OPTIMIZATION NOTICE: PathMatchers match regex rules against raw endpoint patterns securely.
	cr.router.Use(appmware.NewAuthExtractor(
		libhttp.NewHTTPPathMatchers([]*libhttp.PathMatcher{
			libhttp.NewPathMatcher(http.MethodGet, "/metrics", "^/metrics.*$"),
			libhttp.NewPathMatcher(http.MethodGet, "/swagger", "^/swagger.*$"),
			libhttp.NewPathMatcher(http.MethodGet, "/status", "^/status.*$"),
			libhttp.NewPathMatcher(http.MethodGet, "/healthz", "^/healthz.*$"),
			libhttp.NewPathMatcher(http.MethodGet, "/readyz", "^/readyz.*$"),
			libhttp.NewPathMatcher(http.MethodGet, "/debug", "^/debug.*$"),
			libhttp.NewPathMatcher(http.MethodGet, "/config", "^/config.*$"),
			libhttp.NewPathMatcher(http.MethodPost, "/api/v1/auth", "/api/v1/auth"),
			libhttp.NewPathMatcher(http.MethodPost, "/api/v1/auth/simple", "/api/v1/auth/simple"),
			libhttp.NewPathMatcher(http.MethodPost, "/api/v1/users/register", "/api/v1/users/register"),
		}),
		authHelper,
		logger,
	).Handle)
}

// setupRoutes registers operational routing endpoints, guarding administrative sub-routers and environments conditional blocks.
func (cr *AppChiRouter) setupRoutes() {
	// health check
	cr.router.Get("/healthz", cr.getHealthz)
	// readiness check
	cr.router.Get("/readyz", cr.getReadyz)
	// config (debug)
	if cr.config.App.Env != libconfig.AppEnvProduction {
		cr.router.Get("/config", cr.getConfig)
	}

	// api
	cr.router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// /auth
			r.Post("/auth", cr.postAPIV1Auth)
			// /auth/simple
			if cr.config.App.Env != libconfig.AppEnvProduction {
				r.Post("/auth/simple", cr.postAPIV1AuthSimple)
			}
			// users sub-router
			r.Route("/users", func(r chi.Router) {
				r.Get("/profile", cr.getAPIV1UserProfile)
				if cr.config.App.Env != libconfig.AppEnvProduction {
					r.Post("/register", cr.postAPIV1UserRegister)
				}
				r.Put("/password", cr.putAPIV1UserChangePassword)
				r.Put("/keys", cr.putAPIV1UserChangeKeys)
			})
			// admin sub-router
			r.Route("/admin", func(r chi.Router) {
				// /users sub-router
				r.Route("/users", func(r chi.Router) {
					r.Get("/{id}", cr.getAPIV1AdminUser)
					r.Get("/search", cr.getAPIV1AdminUserSearch)
					r.Get("/", cr.getAPIV1AdminUsers)
					r.Post("/", cr.postAPIV1AdminUser)
					r.Put("/{id}", cr.putAPIV1AdminUser)
					r.Delete("/{id}", cr.deleteAPIV1AdminUser)
				})
				// /roles sub-route
				r.Route("/roles", func(r chi.Router) {
					r.Get("/{id}", cr.getAPIV1AdminRole)
					r.Get("/search", cr.getAPIV1AdminRoleSearch)
					r.Get("/", cr.getAPIV1AdminRoles)
					r.Post("/", cr.postAPIV1AdminRole)
					r.Put("/{id}", cr.putAPIV1AdminRole)
					r.Delete("/{id}", cr.deleteAPIV1AdminRole)
				})
			})
		})
	})
}
