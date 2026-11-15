package trace

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/telemetry"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/usecase"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// RoleAdminGetByNameInteractor implements the usecase.RoleAdminGetByNameUseCase interface,
// acting as a non-invasive distributed tracing decorator powered by OpenTelemetry API.
// It intercepts role query paths using string-based name lookups to manage tracing spans and capture metadata attributes.
type RoleAdminGetByNameInteractor struct {
	*telemetry.BaseTelemetry                                   // Embedded framework-level telemetry engine core orchestrator
	next                     usecase.RoleAdminGetByNameUseCase // The encapsulated downstream active core business usecase logic
	spanName                 string                            // Pre-formatted OpenTelemetry specification span name
}

// Compile-time interface compliance verification
var _ usecase.RoleAdminGetByNameUseCase = (*RoleAdminGetByNameInteractor)(nil)

// NewRoleAdminGetByNameUseCase acts as a factory constructor mounting the non-invasive tracing telemetry layer for name-based role lookups.
// NOTICE: The constructor name contains a minor casing typo ("NewROle..." with an uppercase 'O' instead of "NewRole...").
func NewRoleAdminGetByNameUseCase(ucName string, next usecase.RoleAdminGetByNameUseCase) *RoleAdminGetByNameInteractor {
	return &RoleAdminGetByNameInteractor{
		next:          next,
		spanName:      fmt.Sprintf("%s.Get", ucName),
		BaseTelemetry: telemetry.NewBaseTelemetry(ucName),
	}
}

// Get executes a traced wrapper sequence over the underlying core business scenario "Get" retrieval procedure using role names.
func (gnt *RoleAdminGetByNameInteractor) Get(ctx context.Context, name string) (*domain.Role, error) {
	// Spawns a dedicated child execution tracker span injected into the propagation context
	ctx, span := gnt.StartSpan(ctx, gnt.spanName)
	defer span.End()

	// Capture high-cardinality lookup parameters into tracking attributes safely (fully aligned with the param. prefix)
	span.SetAttributes(attribute.String("param.name", name))

	res, err := gnt.next.Get(ctx, name)
	if err != nil {
		// Log explicit infrastructure degradation events directly into the telemetry collector record
		span.AddEvent("Get_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
