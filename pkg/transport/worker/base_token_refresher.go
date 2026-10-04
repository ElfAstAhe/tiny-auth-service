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

const baseTokenRefresherNameTemplate = "token-refresher-%s"

// BaseTokenRefresher implements auth.TokenProvider and worker.Scheduler interfaces,
// orchestrating atomic background thread-safe rotations of expired authentication identities.
type BaseTokenRefresher struct {
	*worker.BaseScheduler
	name               string
	mutex              sync.RWMutex
	token              *utils.AtomicString
	opts               *BaseTokenRefresherOptions
	tokenRefreshAction TokenRefreshAction
	logger             logger.Logger
}

// Compile-time interface compliance verifications
var _ auth.TokenProvider = (*BaseTokenRefresher)(nil)
var _ worker.Scheduler = (*BaseTokenRefresher)(nil)

// NewBaseTokenRefresher constructs an isolated, standalone refresher unit and mounts internal abstract time.Timer event listeners.
func NewBaseTokenRefresher(options ...BaseTokenRefresherOption) (*BaseTokenRefresher, error) {
	opts := NewBaseTokenRefresherOptions()
	for _, option := range options {
		option(opts)
	}
	if err := opts.Validate(); err != nil {
		return nil, errs.NewTlCommonError("NewBaseTokenRefresher", "base token refresher options validation failed", err)
	}

	// instance
	res := &BaseTokenRefresher{
		name:               fmt.Sprintf(baseTokenRefresherNameTemplate, opts.Name),
		opts:               opts,
		token:              utils.NewAtomicString(""),
		tokenRefreshAction: opts.TokenRefreshAction,
		logger:             opts.Logger.GetLogger(fmt.Sprintf(baseTokenRefresherNameTemplate, opts.Name)),
	}
	// scheduler
	scheduler, err := worker.NewBaseScheduler(
		worker.WithSchedulerName(opts.Name),
		worker.WithSchedulerStartInterval(opts.StartInterval),
		worker.WithSchedulerScheduleInterval(opts.ScheduleInterval),
		worker.WithSchedulerStopTimeout(opts.StopTimeout),
		worker.WithSchedulerLogger(opts.Logger),
		worker.WithSchedulerTimerDispatcher(res.timerDispatcher),
	)
	if err != nil {
		return nil, errs.NewTlCommonError("NewBaseTokenRefresher", "base token refresher scheduler create failed", err)
	}
	// setup
	res.BaseScheduler = scheduler

	return res, nil
}

// timerDispatcher bridges the base underlying time tick fired from kernel tickers to the actual token renewal procedure.
func (btr *BaseTokenRefresher) timerDispatcher(ctx context.Context, eventTime time.Time) error {
	btr.GetLogger().Debugf("token refresher %s timer event %s start", btr.GetName(), eventTime.Format(time.DateTime))
	defer btr.GetLogger().Debugf("token refresher %s timer event %s finish", btr.GetName(), eventTime.Format(time.DateTime))

	if utils.IsNil(btr.tokenRefreshAction) {
		return errs.NewCommonError(fmt.Sprintf("token refresher %s timer event %s refresh action not applied", btr.GetName(), eventTime.Format(time.DateTime)), nil)
	}

	btr.token.Store("")

	token, err := btr.tokenRefreshAction(ctx, eventTime)
	if err != nil {
		// Dynamically downgrade scheduler pace to prevent spamming upstream OIDC providers under networking outage
		btr.BaseScheduler.GetOpts().ScheduleInterval = btr.GetOpts().ErrorScheduleInterval

		return errs.NewCommonError(fmt.Sprintf("token refresher %s timer event %s token refresh action failed", btr.GetName(), eventTime.Format(time.DateTime)), err)
	}

	btr.mutex.Lock()
	defer btr.mutex.Unlock()

	// Persist the newly acquired signed payload safely
	btr.token.Store(token)

	// Since btr.GetOpts() references the exact same embedded pointer, the original schedule duration is permanently lost after the first error.
	btr.BaseScheduler.GetOpts().ScheduleInterval = btr.GetOpts().ScheduleInterval

	return nil
}

// GetAccessToken yields the currently cached, unexpired valid authorization token sequence.
// Leverages efficient sync.RWMutex shared read paths to guarantee zero contention under extreme high-throughput traffic.
func (btr *BaseTokenRefresher) GetAccessToken() (string, error) {
	btr.mutex.RLock()
	defer btr.mutex.RUnlock()

	if btr.token.Load() == "" {
		return "", errs.NewCommonError("no actual access token", nil)
	}

	return btr.token.Load(), nil
}

// GetName extracts component name
func (btr *BaseTokenRefresher) GetName() string {
	return btr.name
}

// GetOpts extracts a direct reference to the complete configured BaseTokenRefresherOptions passport parameters blueprint.
func (btr *BaseTokenRefresher) GetOpts() *BaseTokenRefresherOptions {
	return btr.opts
}

// GetLogger extracts the isolated, granular structural reporting handles mapped directly onto this module instance boundary.
func (btr *BaseTokenRefresher) GetLogger() logger.Logger {
	return btr.logger
}
