package amqp

import (
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
)

type LoginAttemptObserver struct {
	client *libamqp.ClientSender
}
