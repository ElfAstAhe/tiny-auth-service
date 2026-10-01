package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/broker"
	libkafka "github.com/ElfAstAhe/go-service-template/pkg/transport/broker/kafka"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

// LoginAttemptObserver implements the pubsub.Observer interface, specialized in intercepting internal
// login event streams and asynchronously translating them into persistent Kafka broker message wire payloads.
type LoginAttemptObserver struct {
	name           string        // Unique identifier naming the current observer component handle
	sender         broker.Sender // Framework abstraction responsible for routing outbound messages to active brokers
	senderKindConf string        // Configuration metric defining the specific sender protocol implementation variant
}

// Compile-time interface compliance verification
var _ pubsub.Observer[*dto.LoginAttemptEventDTO] = (*LoginAttemptObserver)(nil)

// NewLoginAttemptObserver acts as a factory constructor mounting pubsub broadcast listener endpoints backed by broker senders.
func NewLoginAttemptObserver(
	name string,
	sender broker.Sender,
	senderKind string,
) *LoginAttemptObserver {
	return &LoginAttemptObserver{
		name:           name,
		sender:         sender,
		senderKindConf: senderKind,
	}
}

// GetName extracts the registration identifier string present in the observer setup configuration.
func (lak *LoginAttemptObserver) GetName() string {
	return lak.name
}

// OnNotify triggers when a user login event materializes, marshaling structured DTO data into JSON and streaming it via Kafka routing protocols.
func (lak *LoginAttemptObserver) OnNotify(ctx context.Context, data *dto.LoginAttemptEventDTO) error {
	if utils.IsNil(data) {
		return errs.NewCommonError(fmt.Sprintf("%s observer got nil event data", lak.GetName()), nil)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return errs.NewCommonError("json encode failed", err)
	}

	msg := &libkafka.Message{
		TargetName: lak.sender.GetTargetName(),
		Payload:    payload,
		Props:      make(map[string]any),
	}
	msg.Props["content-type"] = "application/json"
	msg.Props["kafka_message_key"] = data.Username

	if err = lak.sender.Publish(ctx, msg); err != nil {
		return errs.NewCommonError(fmt.Sprintf("%s observer failed to publish to target %s", lak.GetName(), lak.sender.GetTargetName()), err)
	}

	return nil
}
