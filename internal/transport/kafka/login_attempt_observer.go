package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
	libamqpkafka "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp/kafka"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type LoginAttemptObserver struct {
	name           string
	sender         libamqp.Sender[any]
	senderKindConf string
}

var _ pubsub.Observer[*dto.LoginAttemptEventDTO] = (*LoginAttemptObserver)(nil)

func NewLoginAttemptObserver(
	name string,
	sender libamqp.Sender[any],
	senderKind string,
) *LoginAttemptObserver {
	return &LoginAttemptObserver{
		name:           name,
		sender:         sender,
		senderKindConf: senderKind,
	}
}

func (lak *LoginAttemptObserver) GetName() string {
	return lak.name
}

func (lak *LoginAttemptObserver) OnNotify(ctx context.Context, data *dto.LoginAttemptEventDTO) error {
	if utils.IsNil(data) {
		return errs.NewCommonError(fmt.Sprintf("%s observer got nil event data", lak.GetName()), nil)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return errs.NewCommonError("json encode failed", err)
	}

	msg := &libamqpkafka.Message{
		TargetName: lak.sender.GetTargetName(),
		Payload:    payload,
		Props:      make(map[string]any),
	}
	msg.Props["content-type"] = "application/json"
	msg.Props["kafka_message_key"] = data.Username

	if err = lak.sender.Publish(ctx, msg, nil); err != nil {
		return errs.NewCommonError(fmt.Sprintf("%s observer failed to publish to target %s", lak.GetName(), lak.sender.GetTargetName()), err)
	}

	return nil
}
