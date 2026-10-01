package amqp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/broker/amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/broker/amqp/azure"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// LoginAttemptObserver implements the pubsub.Observer interface, specialized in intercepting internal
// login event streams and asynchronously translating them into persistent AMQP broker message wire payloads.
type LoginAttemptObserver struct {
	name           string            // Unique identifier naming the current observer component handle
	sender         libamqp.Sender    // Framework abstraction responsible for routing outbound messages to AMQP-compatible brokers
	sendOpts       *amqp.SendOptions // Technical parameters defining the explicit network settlement boundaries for delivery
	senderKindConf string            // Configuration metric defining the specific sender protocol implementation variant
}

// Compile-time interface compliance verification
var _ pubsub.Observer[*dto.LoginAttemptEventDTO] = (*LoginAttemptObserver)(nil)

// NewLoginAttemptObserver acts as a factory constructor mounting pubsub broadcast listener endpoints backed by AMQP senders.
func NewLoginAttemptObserver(
	name string,
	sender libamqp.Sender,
	senderKind string,
) *LoginAttemptObserver {
	return &LoginAttemptObserver{
		name:           name,
		sender:         sender,
		senderKindConf: senderKind,
		sendOpts: &amqp.SendOptions{
			Settled: true,
		},
	}
}

// GetName extracts the registration identifier string present in the observer setup configuration.
func (laa *LoginAttemptObserver) GetName() string {
	return laa.name
}

// OnNotify triggers when a user login event materializes, marshaling structured DTO data into JSON and streaming it via AMQP routing protocols.
func (laa *LoginAttemptObserver) OnNotify(ctx context.Context, data *dto.LoginAttemptEventDTO) error {
	if utils.IsNil(data) {
		return errs.NewCommonError(fmt.Sprintf("%s observer got nil event data", laa.GetName()), nil)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return errs.NewCommonError("json encode failed", err)
	}

	msg := &azure.Message{
		Header: &amqp.MessageHeader{
			Durable: true,
		},
		Payload: payload,
		Props:   make(map[string]any),
	}

	if err = laa.sender.PublishWithOpts(ctx, msg, laa.sendOpts); err != nil {
		return errs.NewCommonError(fmt.Sprintf("%s observer failed to publish to target %s", laa.GetName(), laa.sender.GetTargetName()), err)
	}

	return nil
}
