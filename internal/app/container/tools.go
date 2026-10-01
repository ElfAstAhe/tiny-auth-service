package container

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	// InstanceHashCipher defines the lookup token key targeted for secure irreversible hashing operations.
	InstanceHashCipher string = "sha256-cipher"
	// InstanceDataCipher structures the system key handle mapping symmetric block data cell encryption utilities.
	InstanceDataCipher string = "aes-gcm-cipher"
	// InstanceDataCipherHelper tracks the registration identifier for common framework data encryption wrappers.
	InstanceDataCipherHelper string = "data-cipher-helper"
	// InstanceKeysHelper sets the lookup tag managing automated asymmetric public-private key blocks generation.
	InstanceKeysHelper string = "rsa-2048-keys-helper"
	// InstanceJWTHelper orchestrates framework validation credentials serializing target cryptographic JSON web tokens.
	InstanceJWTHelper string = "jwt-helper"
	// InstanceJWTHTTPHelper maps unique configuration keys driving transport-level HTTP token extraction pipelines.
	InstanceJWTHTTPHelper string = "jwt-http-helper"
	// InstanceJWTGRPCHelper maps unique configuration keys driving transport-level gRPC metadata interceptor routines.
	InstanceJWTGRPCHelper string = "jwt-grpc-helper"
	// InstanceAuthHelper sets global registration parameters handling identity claims context verification.
	InstanceAuthHelper string = "auth-helper"
)

// ToolsContainer structures a lazy-loaded lifecycle dependency injection container managing cryptographic and cryptographic utility helpers.
type ToolsContainer struct {
	*container.BaseLazyContainer // Generic framework-level baseline container orchestration handle
}

// Compile-time interface compliance verifications
var _ container.Container = (*ToolsContainer)(nil)
var _ container.LazyContainer = (*ToolsContainer)(nil)

// NewToolsContainer acts as a factory constructor deploying structural configuration, logger, and orchestrator boundaries.
func NewToolsContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *ToolsContainer {
	return &ToolsContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(ToolsContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

// Init triggers synchronous registration sequences linking required lazy programmatic utility callbacks inside the dependency graph.
func (tc *ToolsContainer) Init(ctx context.Context) error {
	err := errors.Join(
		tc.RegisterProvider(InstanceHashCipher, tc.providerHashCipher),
		tc.RegisterProvider(InstanceDataCipher, tc.providerDataCipher),
		tc.RegisterProvider(InstanceDataCipherHelper, tc.providerDataCipherHelper),
		tc.RegisterProvider(InstanceKeysHelper, tc.providerKeysHelper),
		tc.RegisterProvider(InstanceJWTHelper, tc.providerJWTHelper),
		tc.RegisterProvider(InstanceJWTHTTPHelper, tc.providerJWTHTTPHelper),
		tc.RegisterProvider(InstanceJWTGRPCHelper, tc.providerJWTGRPCHelper),
		tc.RegisterProvider(InstanceAuthHelper, tc.providerAuthHelper),
	)
	if err != nil {
		return errs.NewContainerError(tc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
