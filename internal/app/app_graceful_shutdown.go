package app

import (
	"os"
	"os/signal"
	"syscall"
)

// gracefulShutdown - внутренний метод "агрессивного" закрытия приложения (ctrl+c) + остальные сигналы OS на закрытие
func (app *App) gracefulShutdown() {
	defer app.wg.Done()

	//log := app.logger.GetLogger("App.gracefulShutdown")
	// channel
	sig := make(chan os.Signal, 1)
	// register channel signals
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// awaiting signal
	select {
	case <-sig:
		app.cancel()
	case <-app.ctx.Done():
		signal.Stop(sig)
	}
}
