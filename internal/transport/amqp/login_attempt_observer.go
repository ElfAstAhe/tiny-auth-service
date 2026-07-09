package amqp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/go-amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/infra/pubsub"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
)

type LoginAttemptObserver struct {
	name     string
	client   libamqp.ClientSingleSender[*amqp.SendOptions]
	sendOpts *amqp.SendOptions
}

var _ pubsub.Observer[*dto.LoginAttemptEventDTO] = (*LoginAttemptObserver)(nil)

func NewLoginAttemptObserver(
	name string,
	client libamqp.ClientSingleSender[*amqp.SendOptions],
) *LoginAttemptObserver {
	return &LoginAttemptObserver{
		name:   name,
		client: client,
		sendOpts: &amqp.SendOptions{
			Settled: true,
		},
	}
}

func (lao *LoginAttemptObserver) GetName() string {
	return lao.name
}

func (lao *LoginAttemptObserver) OnNotify(ctx context.Context, data *dto.LoginAttemptEventDTO) error {
	if utils.IsNil(data) {
		return errs.NewCommonError(fmt.Sprintf("%s observer got nil event data", lao.GetName()), nil)
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return errs.NewCommonError("json encode failed", err)
	}

	msg := &libamqp.Message{
		Payload:    payload,
		Properties: make(map[string]any),
	}

	if err = lao.client.Publish(ctx, msg, lao.sendOpts); err != nil {
		return errs.NewCommonError(fmt.Sprintf("%s observer failed to publish to target name %s", lao.GetName(), lao.client.GetTargetName()), err)
	}

	return nil
}
