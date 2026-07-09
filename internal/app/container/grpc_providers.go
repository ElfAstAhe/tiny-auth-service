package container

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	libconfig "github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	libgrpc "github.com/ElfAstAhe/go-service-template/pkg/transport/grpc"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/grpc/interceptors"
	auditlibgrpcintercept "github.com/ElfAstAhe/tiny-audit-service/pkg/transport/interceptor"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade"
	domgrpc "github.com/ElfAstAhe/tiny-auth-service/internal/transport/grpc"
	"github.com/ElfAstAhe/tiny-auth-service/internal/transport/grpc/interceptor"
	pb "github.com/ElfAstAhe/tiny-auth-service/pkg/api/grpc/tiny-auth-service/v1"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func (gc *GRPCContainer) serviceRegister(server *grpc.Server) error {
	authServiceInst, err := container.GetInstance[*domgrpc.AuthGRPCService](InstanceAuthGRPCService)
	if err != nil {
		return errs.NewContainerError(gc.GetName(), "service register: retrieve instance failed", err)
	}
	userServiceInst, err := container.GetInstance[*domgrpc.UserGRPCService](InstanceUserGRPCService)
	if err != nil {
		return errs.NewContainerError(gc.GetName(), "service register: retrieve instance failed", err)
	}
	userAdminServiceInst, err := container.GetInstance[*domgrpc.UserAdminGRPCService](InstanceUserAdminGRPCService)
	if err != nil {
		return errs.NewContainerError(gc.GetName(), "service register: retrieve instance failed", err)
	}
	roleAdminServiceInst, err := container.GetInstance[*domgrpc.RoleAdminGRPCService](InstanceRoleAdminGRPCService)
	if err != nil {
		return errs.NewContainerError(gc.GetName(), "service register: retrieve instance failed", err)
	}

	pb.RegisterAuthServiceServer(server, authServiceInst)
	pb.RegisterUserServiceServer(server, userServiceInst)
	pb.RegisterAdminUsersServiceServer(server, userAdminServiceInst)
	pb.RegisterAdminRolesServiceServer(server, roleAdminServiceInst)

	return nil
}

//goland:noinspection DuplicatedCode
func (gc *GRPCContainer) serverProvider(conf *libconfig.GRPCConfig) (*grpc.Server, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}
	authHelperInst, err := container.GetInstance[auth.Helper](InstanceAuthHelper)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}

	// Настраиваем KeepAlive на основе твоего GRPCConfig
	kasp := keepalive.ServerParameters{
		MaxConnectionIdle:     conf.MaxConnIdle,
		MaxConnectionAge:      conf.MaxConnAge,
		MaxConnectionAgeGrace: conf.MaxConnAgeGrace,
		Time:                  conf.KeepAliveTime,
		Timeout:               conf.KeepAliveTimeout,
	}
	// Метрики
	srvMetrics := grpcprom.NewServerMetrics(
		grpcprom.WithServerHandlingTimeHistogram(
			grpcprom.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
		// Add tenant_name as a context label. This server option is necessary
		// to initialize the metrics with the labels that will be provided
		// dynamically from the context. This should be used in tandem with
		// WithLabelsFromContext in the interceptor options.
		grpcprom.WithContextLabels("tenant_name"),
	)
	if err := prometheus.Register(srvMetrics); err != nil {
		return nil, errs.NewCommonError("failed to register grpc metrics", err)
	}
	exemplarFromContext := func(ctx context.Context) prometheus.Labels {
		if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
			return prometheus.Labels{"traceID": span.TraceID().String()}
		}

		return nil
	}
	// Extract the tenant name value from gRPC metadata
	// and use it as a label on our metrics.
	labelsFromContext := func(ctx context.Context) prometheus.Labels {
		labels := prometheus.Labels{}

		md := metadata.ExtractIncoming(ctx)
		tenantName := md.Get("tenant-name")
		if tenantName == "" {
			tenantName = "unknown"
		}
		labels["tenant_name"] = tenantName

		return labels
	}
	// Setup metric for panic recoveries.
	panicsTotal := promauto.NewCounter(prometheus.CounterOpts{
		Name: "grpc_req_panics_recovered_total",
		Help: "Total number of gRPC requests recovered from internal panic.",
	})
	grpcPanicRecoveryHandler := func(p any) (err error) {
		panicsTotal.Inc()
		//		rpcLogger.Error("recovered from panic", "panic", p, "stack", debug.Stack())
		return status.Errorf(codes.Internal, "%s", p)
	}

	authExtractor := interceptor.NewAuthExtractor(authHelperInst, logInst)
	// Собираем опции сервера
	opts := []grpc.ServerOption{
		// keepalive
		grpc.KeepaliveParams(kasp),
		// timeout
		grpc.ConnectionTimeout(conf.Timeout),
		// tracing
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		// metrics
		grpc.ChainUnaryInterceptor(
			srvMetrics.UnaryServerInterceptor(
				grpcprom.WithExemplarFromContext(exemplarFromContext),
				grpcprom.WithLabelsFromContext(labelsFromContext),
			),
			interceptors.RequestIDExtractorUSInterceptor(
				interceptors.MDXRequestID,
				interceptors.MDXCorrelationID,
				interceptors.MDRequestID,
			),
			interceptors.TraceIDExtractorUSInterceptor(
				interceptors.MDXCloudTraceContext,
				interceptors.MDTraceParent,
				interceptors.MDXTraceID,
				interceptors.MDTraceID,
			),
			interceptors.NewDefaultRealIPExtractorUSInterceptor().UnaryServerInterceptor(),
			auditlibgrpcintercept.AuditRequestIDExtractorUnaryServerInterceptor(
				auditlibgrpcintercept.MDXRequestID,
				auditlibgrpcintercept.MDXCorrelationID,
				auditlibgrpcintercept.MDRequestID,
			),
			auditlibgrpcintercept.AuditTraceIDExtractorUnaryServerInterceptor(
				auditlibgrpcintercept.MDXCloudTraceContext,
				auditlibgrpcintercept.MDTraceParent,
				auditlibgrpcintercept.MDXTraceID,
				auditlibgrpcintercept.MDTraceID,
			),
			authExtractor.UnaryServerInterceptor,
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			srvMetrics.StreamServerInterceptor(
				grpcprom.WithExemplarFromContext(exemplarFromContext),
				grpcprom.WithLabelsFromContext(labelsFromContext),
			),
			interceptors.RequestIDExtractorSSInterceptor(
				interceptors.MDXRequestID,
				interceptors.MDXCorrelationID,
				interceptors.MDRequestID,
			),
			interceptors.TraceIDExtractorSSInterceptor(
				interceptors.MDXCloudTraceContext,
				interceptors.MDTraceParent,
				interceptors.MDXTraceID,
				interceptors.MDTraceID,
			),
			auditlibgrpcintercept.AuditRequestIDExtractorStreamServerInterceptor(
				auditlibgrpcintercept.MDXRequestID,
				auditlibgrpcintercept.MDXCorrelationID,
				auditlibgrpcintercept.MDRequestID,
			),
			auditlibgrpcintercept.AuditTraceIDExtractorStreamServerInterceptor(
				auditlibgrpcintercept.MDXCloudTraceContext,
				auditlibgrpcintercept.MDTraceParent,
				auditlibgrpcintercept.MDXTraceID,
				auditlibgrpcintercept.MDTraceID,
			),
			authExtractor.StreamServerInterceptor,
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
		),
	}

	srv := grpc.NewServer(opts...)

	if confInst.App.Env != libconfig.AppEnvProduction {
		reflection.Register(srv)
	}

	// Инициализация метрик с нулевыми рядами
	srvMetrics.InitializeMetrics(srv)

	return srv, nil
}

func (gc *GRPCContainer) providerAuthService() (any, error) {
	authFacadeInst, err := container.GetInstance[facade.AuthFacade](InstanceAuthFacade)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}

	return domgrpc.NewAuthGRPCService(authFacadeInst), nil
}

func (gc *GRPCContainer) providerUserService() (any, error) {
	userFacadeInst, err := container.GetInstance[facade.UserFacade](InstanceUserFacade)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}

	return domgrpc.NewUserGRPCService(userFacadeInst), nil
}

func (gc *GRPCContainer) providerUserAdminService() (any, error) {
	userAdminFacadeInst, err := container.GetInstance[facade.UserAdminFacade](InstanceUserAdminFacade)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}

	return domgrpc.NewUserAdminGRPCService(userAdminFacadeInst), nil
}

func (gc *GRPCContainer) providerRoleAdminService() (any, error) {
	roleAdminFacadeInst, err := container.GetInstance[facade.RoleAdminFacade](InstanceRoleAdminFacade)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}

	return domgrpc.NewRoleAdminGRPCService(roleAdminFacadeInst), nil
}

func (gc *GRPCContainer) providerGRPCRunner() (any, error) {
	confInst, err := container.GetInstance[*config.Config](InstanceConfig)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}
	logInst, err := container.GetInstance[logger.Logger](InstanceLogger)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), "provider: retrieve instance failed", err)
	}

	runner, err := libgrpc.NewRunner(
		libgrpc.WithName("main-grpc-server"),
		libgrpc.WithConfig(confInst.GRPC),
		libgrpc.WithLogger("grpc_server", logInst),
		libgrpc.WithServiceRegister(gc.serviceRegister),
		libgrpc.WithServerProvider(gc.serverProvider),
	)
	if err != nil {
		return nil, errs.NewContainerError(gc.GetName(), fmt.Sprintf("provider: create %s failed", InstanceGRPCRunner), err)
	}

	return runner, nil
}
