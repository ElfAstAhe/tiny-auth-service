package app

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Close - метод освобождения ресурсов приложения
//
//	if err := app.Close(); err != nil {
//		log.Errorf("app close error [%v]", err)
//
//		panic(errs.NewAppCommonError("app close failed", err))
//	}
func (app *App) Close() {
	log := app.logger.GetLogger("App.Close")

	log.Info("stop services")
	app.stopStartup()

	log.Info("close db connection")
	if app.db != nil {
		if err := app.db.Close(); err != nil {
			log.Errorf("failed to close db [%v]", err)
		}
	}

	log.Info("close telemetry service")
	if app.telemetryShutdown != nil {
		log.Info("shutting down telemetry batcher...")
		// Используем свежий контекст, так как app.ctx может быть уже отменен
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.telemetryShutdown(ctx); err != nil {
			log.Errorf("telemetry shutdown error: [%v]", err)
		} else {
			log.Info("telemetry flushed and closed")
		}
	}
}

func (app *App) stopStartup() {
	log := app.logger.GetLogger("App.stopStartup")
	var stopWG sync.WaitGroup

	if app.httpServer != nil {
		stopWG.Add(1)
		go func() { // stop HTTP
			defer stopWG.Done()

			ctxTimed, cancelTimed := context.WithTimeout(context.Background(), app.config.HTTP.ShutdownTimeout)
			defer cancelTimed()

			// stop http
			log.Info("shutdown http server...")
			if err := app.httpServer.Shutdown(ctxTimed); err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					log.Error("http server shutdown timed out (force close)")
				} else {
					log.Warnf("shutdown http server with error [%v]", err)
				}
			}
			log.Info("shutdown http server complete")
		}()
	}

	if app.grpcServer != nil {
		stopWG.Add(1)
		go func() { // stop gRPC
			defer stopWG.Done()

			log.Info("shutdown gRPC server...")
			defer log.Info("done shutdown gRPC server")

			doneChan := make(chan struct{})
			go func() {
				app.grpcServer.GracefulStop()
				close(doneChan)
			}()
			select {
			case <-doneChan:
				log.Info("shutdown gRPC server complete")
			case <-time.After(app.config.GRPC.ShutdownTimeout):
				log.Error("gRPC graceful shutdown timed out: forcing stop")
				app.grpcServer.Stop()
			}
		}()
	}

	if app.authAuditClient != nil {
		stopWG.Add(1)
		go func() {
			defer stopWG.Done()

			log.Info("stoping auth audit client...")
			defer log.Info("done stop auth audit client")

			if err := app.authAuditClient.Stop(15 * time.Second); err != nil {
				log.Errorf("failed to stop auth audit client [%v]", err)
			}
		}()
	}

	if app.dataAuditClient != nil {
		stopWG.Add(1)
		go func() {
			defer stopWG.Done()
			log.Info("stoping data audit client...")
			defer log.Info("done stop data audit client")

			if err := app.dataAuditClient.Stop(15 * time.Second); err != nil {
				log.Errorf("failed to stop data audit client [%v]", err)
			}
		}()
	}

	// Ожидаем завершения остановки всех серверов и сервисов
	stopWG.Wait()
}
