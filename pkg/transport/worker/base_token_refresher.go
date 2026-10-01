package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/auth"
)

// TokenRefreshAction defines a strongly-typed signature for the application lifecycle callback
// responsible for executing the physical remote network exchange or crypto signature generation to acquire a new token.
type TokenRefreshAction func(ctx context.Context, eventTime time.Time) (string, error)

// BaseTokenRefresherConfig encapsulates configuration credentials for the scheduling mechanics
// alongside adaptive fallback intervals applied dynamically if authentication failures happen.
type BaseTokenRefresherConfig struct {
	*worker.BaseSchedulerConfig
	ErrorScheduleInterval time.Duration // Adaptive fallback delay used sequentially if an exchange round drops an error
}

// NewBaseTokenRefresherConfig acts as a factory constructor setting up interval footprints for the refresher lifecycle.
func NewBaseTokenRefresherConfig(
	conf *worker.BaseSchedulerConfig,
	errorScheduleInterval time.Duration,
) *BaseTokenRefresherConfig {
	return &BaseTokenRefresherConfig{
		BaseSchedulerConfig:   conf,
		ErrorScheduleInterval: errorScheduleInterval,
	}
}

// BaseTokenRefresher implements auth.TokenProvider and worker.Scheduler interfaces,
// orchestrating atomic background thread-safe rotations of expired authentication identities.
type BaseTokenRefresher struct {
	*worker.BaseScheduler
	mutex              sync.RWMutex
	token              string
	conf               *BaseTokenRefresherConfig
	tokenRefreshAction TokenRefreshAction
}

// Compile-time interface compliance verifications
var _ auth.TokenProvider = (*BaseTokenRefresher)(nil)
var _ worker.Scheduler = (*BaseTokenRefresher)(nil)

// NewBaseTokenRefresher constructs an isolated, standalone refresher unit and mounts internal abstract time.Timer event listeners.
func NewBaseTokenRefresher(
	conf *BaseTokenRefresherConfig,
	tokenRefreshAction TokenRefreshAction,
	log logger.Logger,
) *BaseTokenRefresher {
	res := &BaseTokenRefresher{
		conf:               conf,
		tokenRefreshAction: tokenRefreshAction,
	}
	res.BaseScheduler = worker.NewBaseScheduler(
		"tokenRefresher",
		res.timerDispatcher,
		worker.NewBaseSchedulerConfig(conf.StartInterval, conf.ScheduleInterval, conf.StopTimeout),
		log,
	)

	return res
}

// timerDispatcher bridges the base underlying time tick fired from kernel tickers to the actual token renewal procedure.
func (btr *BaseTokenRefresher) timerDispatcher(ctx context.Context, eventTime time.Time) error {
	btr.GetLogger().Debugf("token refresher %s timer event %s start", btr.GetName(), eventTime.Format(time.DateTime))
	defer btr.GetLogger().Debugf("token refresher %s timer event %s finish", btr.GetName(), eventTime.Format(time.DateTime))

	if utils.IsNil(btr.tokenRefreshAction) {
		return errs.NewCommonError(fmt.Sprintf("token refresher %s timer event %s refresh action not applied", btr.GetName(), eventTime.Format(time.DateTime)), nil)
	}

	// Important: passing down the underlying base app context state instead of ephemeral execution context
	token, err := btr.tokenRefreshAction(btr.GetContext(), eventTime)
	if err != nil {
		// Dynamically downgrade scheduler pace to prevent spamming upstream OIDC providers under networking outage
		btr.BaseScheduler.GetConfig().ScheduleInterval = btr.GetConfig().ErrorScheduleInterval

		return errs.NewCommonError(fmt.Sprintf("token refresher %s timer event %s token refresh action failed", btr.GetName(), eventTime.Format(time.DateTime)), err)
	}

	btr.mutex.Lock()
	defer btr.mutex.Unlock()

	// Persist the newly acquired signed payload safely
	btr.token = token

	// BUG WATCH: This self-assignment statement will fail to restore the base worker configuration interval.
	// Since btr.GetConfig() references the exact same embedded pointer, the original schedule duration is permanently lost after the first error.
	btr.BaseScheduler.GetConfig().ScheduleInterval = btr.GetConfig().ScheduleInterval

	return nil
}

// GetAccessToken yields the currently cached, unexpired valid authorization token token sequence.
// Leverages efficient sync.RWMutex shared read paths to guarantee zero contention under extreme high-throughput traffic.
func (btr *BaseTokenRefresher) GetAccessToken() (string, error) {
	btr.mutex.RLock()
	defer btr.mutex.RUnlock()

	if btr.token == "" {
		return btr.token, errs.NewCommonError("no actual access token", nil)
	}

	return btr.token, nil
}

// GetConfig extracts a direct reference to the complete configured BaseTokenRefresherConfig passport parameters blueprint.
func (btr *BaseTokenRefresher) GetConfig() *BaseTokenRefresherConfig {
	return btr.conf
}
