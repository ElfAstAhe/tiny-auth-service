package rest

import (
	"net/http"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport"
	libmware "github.com/ElfAstAhe/go-service-template/pkg/transport/middleware"
	_ "github.com/ElfAstAhe/tiny-auth-service/docs"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	trmware "github.com/ElfAstAhe/tiny-auth-service/internal/transport/rest/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hellofresh/health-go/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/riandyrn/otelchi"
	swagh "github.com/swaggo/http-swagger"
)

type AppChiRouter struct {
	router          *chi.Mux
	log             logger.Logger
	config          *config.Config
	health          *health.Health
	healthz         transport.HealthzFunc
	readyz          transport.ReadyzFunc
	authFacade      facade.AuthFacade
	userFacade      facade.UserFacade
	userAdminFacade facade.UserAdminFacade
	roleAdminFacade facade.RoleAdminFacade
}

var _ transport.HTTPRouter = (*AppChiRouter)(nil)

func NewAppChiRouter(
	config *config.Config,
	logger logger.Logger,
	authHelper auth.Helper,
	health *health.Health,
	healthz transport.HealthzFunc,
	readyz transport.ReadyzFunc,
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

func (cr *AppChiRouter) GetRouter() http.Handler {
	return cr.router
}

func (cr *AppChiRouter) setupMiddleware(
	authHelper auth.Helper,
	logger logger.Logger,
) {
	// tracing
	cr.router.Use(otelchi.Middleware(cr.config.Telemetry.ServiceName, otelchi.WithChiRoutes(cr.router)))
	// metrics
	cr.router.Use(libmware.HTTPMetricsMiddleware)
	// requestID
	cr.router.Use(middleware.RequestID)
	// realIP
	cr.router.Use(middleware.RealIP)
	// recoverer
	cr.router.Use(middleware.Recoverer)
	// timeout
	cr.router.Use(middleware.Timeout(cr.config.HTTP.ReadTimeout))
	// compress (add any content-types)
	cr.router.Use(libmware.NewHTTPCompress(logger,
		"application/json", "plain/text",
	).Handle)
	// decompress
	cr.router.Use(libmware.NewHTTPDecompress(int64(cr.config.HTTP.MaxRequestBodySize), logger).Handle)
	// jwt auth extractor - extract user info from token
	cr.router.Use(trmware.NewAuthExtractor(
		transport.NewHTTPPathMatchers([]*transport.HTTPPathMatcher{
			transport.NewHTTPPathMatcher(http.MethodGet, "/metrics", "^/metrics.*$"),
			transport.NewHTTPPathMatcher(http.MethodGet, "/swagger", "^/swagger.*$"),
			transport.NewHTTPPathMatcher(http.MethodGet, "/status", "^/status.*$"),
			transport.NewHTTPPathMatcher(http.MethodGet, "/healthz", "^/healthz.*$"),
			transport.NewHTTPPathMatcher(http.MethodGet, "/readyz", "^/readyz.*$"),
			transport.NewHTTPPathMatcher(http.MethodGet, "/debug", "^/debug.*$"),
			transport.NewHTTPPathMatcher(http.MethodGet, "/config", "^/config.*$"),
			transport.NewHTTPPathMatcher(http.MethodPost, "/api/v1/auth", "/api/v1/auth"),
			transport.NewHTTPPathMatcher(http.MethodPost, "/api/v1/auth/simple", "/api/v1/auth/simple"),
			transport.NewHTTPPathMatcher(http.MethodPost, "/api/v1/users/register", "/api/v1/users/register"),
		}),
		authHelper,
		logger,
	).Handle)
	// income/outcome logger
	cr.router.Use(libmware.NewHTTPRequestLogger(logger).Handle)
}

func (cr *AppChiRouter) setupRoutes() {
	// health check
	cr.router.Get("/healthz", cr.getHealthz)
	// readiness check
	cr.router.Get("/readyz", cr.getReadyz)
	// config (debug)
	if cr.config.App.Env != config.AppEnvProduction {
		cr.router.Get("/config", cr.getConfig)
	}

	// api
	cr.router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// /auth
			r.Post("/auth", cr.postAPIV1Auth)
			// /auth/simple
			if cr.config.App.Env != config.AppEnvProduction {
				r.Post("/auth/simple", cr.postAPIV1AuthSimple)
			}
			// users sub-router
			r.Route("/users", func(r chi.Router) {
				r.Get("/profile", cr.getAPIV1UserProfile)
				if cr.config.App.Env != config.AppEnvProduction {
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
