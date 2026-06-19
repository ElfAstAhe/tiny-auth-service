package amqp

import (
	"github.com/Azure/go-amqp"
)

type LoginAttemptObserver struct {
	client *amqp.Sender
}
