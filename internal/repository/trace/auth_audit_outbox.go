package trace

import (
	"context"
	"fmt"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	libextdom "github.com/ElfAstAhe/go-service-template/pkg/ext/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/trace"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type AuthAuditOutbox struct {
	*trace.CRUD[*domain.AuthAuditOutbox, string]
	next domain.AuthAuditOutboxRepository
}

var _ libdom.CRUDRepository[*domain.AuthAuditOutbox, string] = (*AuthAuditOutbox)(nil)
var _ domain.AuthAuditOutboxRepository = (*AuthAuditOutbox)(nil)
var _ libextdom.TailRepository[string] = (*AuthAuditOutbox)(nil)
var _ libextdom.OutboxRepository[*domain.AuthAuditOutbox, string] = (*AuthAuditOutbox)(nil)

func NewAuthAuditOutboxRepository(next domain.AuthAuditOutboxRepository) *AuthAuditOutbox {
	return &AuthAuditOutbox{
		next: next,
		CRUD: trace.NewCRUD("AuthAuditOutboxRepository", next),
	}
}

// GetTail возвращает срез идентификаторов старых ("хвостовых") записей outbox, созданных до указанной даты.
func (aao *AuthAuditOutbox) GetTail(ctx context.Context, tailDate time.Time, limit int) ([]string, error) {
	repoCtx, span := aao.StartSpan(ctx, fmt.Sprintf("%s.GetTail", aao.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.Int64("param.tailDate", tailDate.Unix()),
	)

	res, err := aao.next.GetTail(repoCtx, tailDate, limit)
	if err != nil {
		span.AddEvent("GetTail_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}

// CutTail физически удаляет запись из "хвоста" таблицы outbox по её уникальному идентификатору.
func (aao *AuthAuditOutbox) CutTail(ctx context.Context, id string) error {
	repoCtx, span := aao.StartSpan(ctx, fmt.Sprintf("%s.CutTail", aao.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.String("params.id", id),
	)

	err := aao.next.CutTail(repoCtx, id)
	if err != nil {
		span.AddEvent("CutTail_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}

// Save сохраняет или обновляет сущность DataAuditOutbox в хранилище в рамках распределенной трассировки.
func (aao *AuthAuditOutbox) Save(ctx context.Context, entity *domain.AuthAuditOutbox) error {
	repoCtx, span := aao.StartSpan(ctx, fmt.Sprintf("%s.Save", aao.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.String("params.entity_id", fmt.Sprintf("%v", entity.GetID())),
	)

	err := aao.next.Save(repoCtx, entity)
	if err != nil {
		span.AddEvent("Save_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}

// ChangeStatus атомарно изменяет текущий статус обработки записи outbox (например, с pending на processed).
func (aao *AuthAuditOutbox) ChangeStatus(ctx context.Context, id string, status string) error {
	repoCtx, span := aao.StartSpan(ctx, fmt.Sprintf("%s.ChangeStatus", aao.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.String("params.id", id),
		attribute.String("params.status", status),
	)

	err := aao.next.ChangeStatus(repoCtx, id, status)
	if err != nil {
		span.AddEvent("ChangeStatus_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}

// ListPending возвращает список записей outbox, ожидающих отправки, с ограничением по количеству.
func (aao *AuthAuditOutbox) ListPending(ctx context.Context, limit int) ([]*domain.AuthAuditOutbox, error) {
	repoCtx, span := aao.StartSpan(ctx, fmt.Sprintf("%s.ListPending", aao.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.Int("param.limit", limit),
	)

	res, err := aao.next.ListPending(repoCtx, limit)
	if err != nil {
		span.AddEvent("ListPending_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}

// ListProcessed возвращает список успешно обработанных записей outbox, измененных после указанного времени.
func (aao *AuthAuditOutbox) ListProcessed(ctx context.Context, lastUpdatedAt time.Time, limit int) ([]*domain.AuthAuditOutbox, error) {
	repoCtx, span := aao.StartSpan(ctx, fmt.Sprintf("%s.ListProcessed", aao.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(
		attribute.Int64("params.last_updated_at", lastUpdatedAt.Unix()),
		attribute.Int("params.limit", limit),
	)

	res, err := aao.next.ListProcessed(repoCtx, lastUpdatedAt, limit)
	if err != nil {
		span.AddEvent("ListProcessed_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
