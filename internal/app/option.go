package app

import (
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-auth-service/internal/config"
)

// Option defines a functional configuration closure pattern designed to lazily populate Application bootstrap fields.
type Option func(*Application)

// WithConfig returns an Option configuring the application layout Config credentials block.
func WithConfig(conf *config.Config) Option {
	return func(app *Application) {
		app.conf = conf
	}
}

// WithLogger returns an Option mapping the structured logging subsystem handle into the application bootstrap graph.
func WithLogger(log logger.Logger) Option {
	return func(app *Application) {
		app.log = log
	}
}
