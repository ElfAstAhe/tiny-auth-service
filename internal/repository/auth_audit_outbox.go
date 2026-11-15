package repository

import (
	"context"
	"strings"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/db"
	libpg "github.com/ElfAstAhe/go-service-template/pkg/db/postgres"
	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	libextdom "github.com/ElfAstAhe/go-service-template/pkg/ext/domain"
	librepo "github.com/ElfAstAhe/go-service-template/pkg/repository/v2"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

const (
	statementAuthAuditOutboxGetTail       string = "statement-auth-audit-outbox-get-tail"
	statementAuthAuditOutboxListPending   string = "statement-auth-audit-outbox-list-pending"
	statementAuthAuditOutboxListProcessed string = "statement-auth-audit-outbox-list-processed"
	statementAuthAuditOutboxChangeStatus  string = "statement-auth-audit-outbox-change-status"
)

type AuthAuditOutbox struct {
	*librepo.CRUD[*domain.AuthAuditOutbox, string]
}

var _ libdom.CRUDRepository[*domain.AuthAuditOutbox, string] = (*AuthAuditOutbox)(nil)
var _ domain.AuthAuditOutboxRepository = (*AuthAuditOutbox)(nil)
var _ libextdom.TailRepository[string] = (*AuthAuditOutbox)(nil)
var _ libextdom.OutboxRepository[*domain.AuthAuditOutbox, string] = (*AuthAuditOutbox)(nil)

func NewAuthAuditOutbox(
	driverName string,
	executor db.Executor,
	errDecipher db.ErrorDecipher,
) (*AuthAuditOutbox, error) {
	// instance
	res := &AuthAuditOutbox{}
	// embed
	embed, err := librepo.NewCRUD[*domain.AuthAuditOutbox, string](
		librepo.WithExecutor[*domain.AuthAuditOutbox, string](executor),
		librepo.WithErrorDecipher[*domain.AuthAuditOutbox, string](errDecipher),
		librepo.WithHooks[*domain.AuthAuditOutbox, string](
			librepo.WithHookNewEntityFactory[*domain.AuthAuditOutbox, string](domain.NewEmptyAuthAuditOutbox),
			librepo.WithHookCreator[*domain.AuthAuditOutbox, string](res.creator),
			librepo.WithHookChanger[*domain.AuthAuditOutbox, string](res.changer),
		),
		librepo.WithEntityInfo[*domain.AuthAuditOutbox, string]("auth_audit", "AuthAudit"),
		librepo.WithDriverName[*domain.AuthAuditOutbox, string](driverName),
		librepo.WithStatementsBuilder[*domain.AuthAuditOutbox, string](libpg.DriverName, res.pgStatementsBuilder),
		librepo.WithStatementsBuilder[*domain.AuthAuditOutbox, string]("pgx", res.pgStatementsBuilder),
		librepo.WithStatementsBuilder[*domain.AuthAuditOutbox, string]("postgres", res.pgStatementsBuilder),
	)
	if err != nil {
		return nil, errs.NewDalError("NewAuthAuditOutbox", "create embed CRUD failed", err)
	}
	// setup
	res.CRUD = embed

	return res, nil
}

// GetTail возвращает список идентификаторов старых записей аудита, созданных до указанной даты.
func (aao *AuthAuditOutbox) GetTail(ctx context.Context, tailDate time.Time) ([]string, error) {
	// fail fast
	if tailDate.IsZero() {
		return []string{}, errs.NewInvalidArgumentError("TailDate", "tail date is required")
	}
	// act
	res, err := aao.GetHelper().List(ctx, statementAuthAuditOutboxGetTail, tailDate)
	if err != nil {
		return []string{}, errs.NewDalError("AuthAuditOutbox.GetTail", "List failed", err)
	}

	return libdom.MapEntitiesToIDs(res), nil
}

// CutTail выполняет физическое удаление записи аудита из "хвоста" таблицы по её идентификатору.
func (aao *AuthAuditOutbox) CutTail(ctx context.Context, id string) error {
	// fail fast
	if strings.TrimSpace(id) == "" {
		return errs.NewInvalidArgumentError("ID", "id can not be empty")
	}
	err := aao.Delete(ctx, id)
	if err != nil {
		return errs.NewDalError("AuthAuditOutbox.CutTail", "Delete failed", err)
	}

	return nil
}

// Save выполняет сохранение (создание) новой записи outbox в базе данных.
// Реализует интерфейс extdom.OutboxRepository.
func (aao *AuthAuditOutbox) Save(ctx context.Context, entity *domain.AuthAuditOutbox) error {
	_, err := aao.Create(ctx, entity)

	return err
}

// ChangeStatus обновляет статус конкретной записи outbox по её идентификатору.
// Реализует интерфейс extdom.OutboxRepository.
func (aao *AuthAuditOutbox) ChangeStatus(ctx context.Context, id string, status string) error {
	if err := aao.GetHelper().ExecNoCheck(ctx, statementAuthAuditOutboxChangeStatus, id, status); err != nil {
		return errs.NewDalError("AuthAuditOutbox.ChangeStatus", "change status failed", err)
	}

	return nil
}

// ListPending извлекает список необработанных записей outbox с ограничением по количеству.
// Реализует интерфейс extdom.OutboxRepository.
func (aao *AuthAuditOutbox) ListPending(ctx context.Context, limit int) ([]*domain.AuthAuditOutbox, error) {
	return aao.GetHelper().List(ctx, statementAuthAuditOutboxListPending, []string{libextdom.OutboxStatusPending}, limit)
}

// ListProcessed извлекает список успешно обработанных записей, обновленных после указанного времени.
// Реализует интерфейс extdom.OutboxRepository.
func (aao *AuthAuditOutbox) ListProcessed(ctx context.Context, lastUpdatedAt time.Time, limit int) ([]*domain.AuthAuditOutbox, error) {
	return aao.GetHelper().List(ctx, statementAuthAuditOutboxListProcessed, []string{libextdom.OutboxStatusFailed, libextdom.OutboxStatusSuccess}, lastUpdatedAt, limit)
}

func (aao *AuthAuditOutbox) creator(ctx context.Context, querier db.Querier, entity *domain.AuthAuditOutbox, params ...any) (librepo.Scannable, error) {
	query, err := aao.GetHelper().GetStatements().GetSQL(librepo.StatementCreate)
	if err != nil {
		return nil, errs.NewDalError("AuthAuditOutbox.creator", "create statement not found", err)
	}

	return querier.QueryRowContext(ctx, query,
		entity.ID,
		entity.Username,
		entity.ContentType,
		entity.Payload,
		entity.Status,
		entity.CreatedAt,
	), nil
}

func (aao *AuthAuditOutbox) changer(ctx context.Context, querier db.Querier, entity *domain.AuthAuditOutbox, params ...any) (librepo.Scannable, error) {
	query, err := aao.GetHelper().GetStatements().GetSQL(librepo.StatementChange)
	if err != nil {
		return nil, errs.NewDalError("AuthAuditOutbox.changer", "change statement not found", err)
	}

	return querier.QueryRowContext(ctx, query,
		entity.ID,
		entity.Username,
		entity.ContentType,
		entity.Payload,
		entity.Status,
		entity.UpdatedAt,
	), nil
}

func (aao *AuthAuditOutbox) mapDefault(ctx context.Context, scannable librepo.Scannable, params ...any) (*domain.AuthAuditOutbox, error) {
	res := aao.GetHelper().GetHooks().NewEntityFactory()
	err := scannable.Scan(
		&res.ID,
		&res.Username,
		&res.ContentType,
		&res.Payload,
		&res.Status,
		&res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return nil, errs.NewDalError("AuthAuditOutbox.mapDefault", "default mapper failed", err)
	}

	return res, nil
}

func (aao *AuthAuditOutbox) mapTail(ctx context.Context, scannable librepo.Scannable, params ...any) (*domain.AuthAuditOutbox, error) {
	res := aao.GetHelper().GetHooks().NewEntityFactory()
	err := scannable.Scan(
		&res.ID,
	)
	if err != nil {
		return nil, errs.NewDalError("AuthAuditOutbox.mapTail", "tail mapper failed", err)
	}

	return res, nil
}
