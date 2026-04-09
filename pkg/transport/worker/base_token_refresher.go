package worker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-auth-service/pkg/transport/auth"
)

type TokenRefreshAction func(ctx context.Context, eventTime time.Time) (string, error)

type BaseTokenRefresherConfig struct {
	*worker.BaseSchedulerConfig
	ErrorScheduleInterval time.Duration
}

func NewBaseTokenRefresherConfig(
	conf *worker.BaseSchedulerConfig,
	errorScheduleInterval time.Duration,
) *BaseTokenRefresherConfig {
	return &BaseTokenRefresherConfig{
		BaseSchedulerConfig:   conf,
		ErrorScheduleInterval: errorScheduleInterval,
	}
}

type BaseTokenRefresher struct {
	*worker.BaseScheduler
	mutex              sync.RWMutex
	token              string
	conf               *BaseTokenRefresherConfig
	tokenRefreshAction TokenRefreshAction
}

var _ auth.TokenProvider = (*BaseTokenRefresher)(nil)
var _ worker.Scheduler = (*BaseTokenRefresher)(nil)

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
		worker.NewBaseSchedulerConfig(conf.StartInterval, conf.ScheduleInterval),
		log,
	)

	return res
}

func (btr *BaseTokenRefresher) timerDispatcher(eventTime time.Time) error {
	btr.GetLogger().Debugf("token refresher %s timer event %s start", btr.GetName(), eventTime.Format(time.DateTime))
	defer btr.GetLogger().Debugf("token refresher %s timer event %s finish", btr.GetName(), eventTime.Format(time.DateTime))

	if btr.tokenRefreshAction == nil {
		return errs.NewCommonError(fmt.Sprintf("token refresher %s timer event %s refresh acton not applied", btr.GetName(), eventTime.Format(time.DateTime)), nil)
	}

	token, err := btr.tokenRefreshAction(btr.GetContext(), eventTime)
	if err != nil {
		// интервал таймера следующей итерации по ошибке
		btr.BaseScheduler.GetConfig().ScheduleInterval = btr.GetConfig().ErrorScheduleInterval

		return errs.NewCommonError(fmt.Sprintf("token refresher %s timer event %s token refresh action failed", btr.GetName(), eventTime.Format(time.DateTime)), err)
	}

	btr.mutex.Lock()
	defer btr.mutex.Unlock()

	// new token
	btr.token = token
	// restore timer interval
	btr.BaseScheduler.GetConfig().ScheduleInterval = btr.GetConfig().ScheduleInterval

	return nil
}

func (btr *BaseTokenRefresher) GetAccessToken() (string, error) {
	btr.mutex.RLock()
	defer btr.mutex.RUnlock()

	if btr.token == "" {
		return btr.token, errs.NewCommonError("no actual access token", nil)
	}

	return btr.token, nil
}

func (btr *BaseTokenRefresher) GetConfig() *BaseTokenRefresherConfig {
	return btr.conf
}
