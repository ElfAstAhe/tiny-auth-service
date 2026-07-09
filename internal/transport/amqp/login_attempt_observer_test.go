package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/Azure/go-amqp"
	libamqp "github.com/ElfAstAhe/go-service-template/pkg/transport/amqp"
	"github.com/ElfAstAhe/go-service-template/pkg/transport/amqp/mocks"
	"github.com/ElfAstAhe/tiny-auth-service/internal/facade/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/goleak"
)

// Гарантируем, что наши тесты обсервера не плодят подвисшие горутины
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

// 1. Тест успешного прохождения события (Happy Path) с глубокой валидацией полей DTO
func TestLoginAttemptObserver_OnNotify_Success(t *testing.T) {
	// Создаем expecter-мок интерфейса ClientSender с помощью mockery
	mockClient := mocks.NewMockClientSingleSender[*amqp.SendOptions](t)

	observerName := "test-login-observer"
	observer := NewLoginAttemptObserver(observerName, mockClient)

	fixedTime := time.Now()
	// Подготавливаем тестовые данные (Ваша полная структура DTO)
	testDTO := &dto.LoginAttemptEventDTO{
		NodeName:  "auth-pod-xyz",
		EventDate: fixedTime,
		RequestID: "req-111-222",
		TraceID:   "trace-777-888",
		IP:        "10.0.0.15",
		Username:  "dev_user",
		Success:   true,
		Error:     "",
	}

	// ИСПОЛЬЗУЕМ СТРОГИЙ СИНТАКСИС TYPE-SAFE EXPECTER (.EXPECT())
	// Проверяем, что обсервер корректно перевел DTO в JSON и сохранил ВСЕ поля контракта обмена
	mockClient.EXPECT().
		Publish(
			mock.Anything,
			mock.MatchedBy(func(msg *libamqp.Message) bool {
				var parsed dto.LoginAttemptEventDTO
				err := json.Unmarshal(msg.Payload, &parsed)

				// Глубокая сверка полей JSON-нагрузки
				return err == nil &&
					parsed.NodeName == "auth-pod-xyz" &&
					parsed.EventDate.Equal(fixedTime) &&
					parsed.RequestID == "req-111-222" &&
					parsed.TraceID == "trace-777-888" &&
					parsed.IP == "10.0.0.15" &&
					parsed.Username == "dev_user" &&
					parsed.Success == true &&
					parsed.Error == ""
			}),
			mock.Anything,
		).
		Return(nil).
		Once()

	// Запуск тестируемого метода OnNotify
	err := observer.OnNotify(context.Background(), testDTO)

	// Проверки (Assertions)
	assert.NoError(t, err)
}

// 2. Тест обработки ошибки сетевого клиента (Publish Failure)
func TestLoginAttemptObserver_OnNotify_PublishError(t *testing.T) {
	mockClient := mocks.NewMockClientSingleSender[*amqp.SendOptions](t)
	mockClient.On("GetTargetName").Return("test-target::test-queue")
	observer := NewLoginAttemptObserver("test-login-observer", mockClient)

	testDTO := &dto.LoginAttemptEventDTO{
		Username: "unstable_user",
		Success:  false,
		Error:    "invalid password",
	}

	// Имитируем сетевой облом на уровне ClientSender
	publishErr := errors.New("amqp connection closed unexpectedly by remote broker")

	mockClient.EXPECT().
		Publish(mock.Anything, mock.Anything, mock.Anything).
		Return(publishErr).
		Once()

	err := observer.OnNotify(context.Background(), testDTO)

	// Проверяем, что ошибка обернута в ваш кастомный errs.NewCommonError
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "test-login-observer observer failed to publish")
}

// 3. Тест защиты от nil-указателя на входе (Nil Data Defense)
func TestLoginAttemptObserver_OnNotify_NilData(t *testing.T) {
	mockClient := mocks.NewMockClientSingleSender[*amqp.SendOptions](t)
	observer := NewLoginAttemptObserver("test-login-observer", mockClient)

	// Передаем nil вместо DTO
	err := observer.OnNotify(context.Background(), nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "test-login-observer observer got nil event data")
}
