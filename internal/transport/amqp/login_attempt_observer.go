package amqp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
	libamqpazure "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp/azure"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type LoginAttemptObserver struct {
	name           string
	sender         libamqp.Sender[*amqp.SendOptions]
	sendOpts       *amqp.SendOptions
	senderKindConf string
}

var _ pubsub.Observer[*dto.LoginAttemptEventDTO] = (*LoginAttemptObserver)(nil)

func NewLoginAttemptObserver(
	name string,
	sender libamqp.Sender[*amqp.SendOptions],
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

func (laa *LoginAttemptObserver) GetName() string {
	return laa.name
}

func (laa *LoginAttemptObserver) OnNotify(ctx context.Context, data *dto.LoginAttemptEventDTO) error {
	if utils.IsNil(data) {
		return errs.NewCommonError(fmt.Sprintf("%s observer got nil event data", laa.GetName()), nil)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return errs.NewCommonError("json encode failed", err)
	}

	msg := &libamqpazure.Message{
		Header: &amqp.MessageHeader{
			Durable: true,
		},
		Payload: payload,
		Props:   make(map[string]any),
	}

	if err = laa.sender.Publish(ctx, msg, laa.sendOpts); err != nil {
		return errs.NewCommonError(fmt.Sprintf("%s observer failed to publish to target %s", laa.GetName(), laa.sender.GetTargetName()), err)
	}

	return nil
}
