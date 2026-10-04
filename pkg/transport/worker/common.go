package worker

import (
	"context"
	"time"
)

// TokenRefreshAction defines a strongly-typed signature for the application lifecycle callback
// responsible for executing the physical remote network exchange or crypto signature generation to acquire a new token.
type TokenRefreshAction func(ctx context.Context, eventTime time.Time) (string, error)
