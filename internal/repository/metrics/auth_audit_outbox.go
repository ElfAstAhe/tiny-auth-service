package metrics

import (
	"context"
	"time"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	libextdom "github.com/ElfAstAhe/go-service-template/pkg/ext/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository/metrics"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type AuthAuditOutbox struct {
	*metrics.CRUD[*domain.AuthAuditOutbox, string]
	next domain.AuthAuditOutboxRepository
}

var _ libdom.CRUDRepository[*domain.AuthAuditOutbox, string] = (*AuthAuditOutbox)(nil)
var _ domain.AuthAuditOutboxRepository = (*AuthAuditOutbox)(nil)
var _ libextdom.TailRepository[string] = (*AuthAuditOutbox)(nil)
var _ libextdom.OutboxRepository[*domain.AuthAuditOutbox, string] = (*AuthAuditOutbox)(nil)

func NewAuthAuditOutboxRepository(next domain.AuthAuditOutboxRepository) *AuthAuditOutbox {
	return &AuthAuditOutbox{
		next: next,
		CRUD: metrics.NewCRUD[*domain.AuthAuditOutbox, string]("AuthAuditOutboxRepository", next),
	}
}

// GetTail получение списка на удаление, реализация интерфейса domain.TailRepository
func (aao *AuthAuditOutbox) GetTail(ctx context.Context, tailDate time.Time) (res []string, err error) {
	defer func(startTime time.Time) {
		metrics.ObserveRepositoryOp(aao.GetRepositoryName(), "GetTail", err, startTime)
	}(time.Now())

	return aao.next.GetTail(ctx, tailDate)
}

// CutTail удаление одного элемента, реализация domain.TailRepository
func (aao *AuthAuditOutbox) CutTail(ctx context.Context, id string) (err error) {
	defer func(startTime time.Time) {
		metrics.ObserveRepositoryOp(aao.GetRepositoryName(), "CutTail", err, startTime)
	}(time.Now())

	return aao.next.CutTail(ctx, id)
}

// Save сохранение элемента domain.DataAuditOutbox, реализация domain.OutboxRepository
func (aao *AuthAuditOutbox) Save(ctx context.Context, entity *domain.AuthAuditOutbox) (err error) {
	defer func(startTime time.Time) {
		metrics.ObserveRepositoryOp(aao.GetRepositoryName(), "Save", err, startTime)
	}(time.Now())

	return aao.next.Save(ctx, entity)
}

// ChangeStatus меняет статус элемента domain.DataAuditOutbox, реализация domain.OutboxRepository
func (aao *AuthAuditOutbox) ChangeStatus(ctx context.Context, id string, status string) (err error) {
	defer func(startTime time.Time) {
		metrics.ObserveRepositoryOp(aao.GetRepositoryName(), "ChangeStatus", err, startTime)
	}(time.Now())

	return aao.next.ChangeStatus(ctx, id, status)
}

// ListPending возвращает список необработанных элементов domain.DataAuditOutbox, реализация domain.OutboxRepository
func (aao *AuthAuditOutbox) ListPending(ctx context.Context, limit int) (res []*domain.AuthAuditOutbox, err error) {
	defer func(startTime time.Time) {
		metrics.ObserveRepositoryOp(aao.GetRepositoryName(), "ListPending", err, startTime)
	}(time.Now())

	return aao.next.ListPending(ctx, limit)
}

// ListProcessed возвращает список обработанных элементов domain.DataAuditOutbox, реализация domain.OutboxRepository
func (aao *AuthAuditOutbox) ListProcessed(ctx context.Context, lastUpdatedAt time.Time, limit int) (res []*domain.AuthAuditOutbox, err error) {
	defer func(startTime time.Time) {
		metrics.ObserveRepositoryOp(aao.GetRepositoryName(), "ListProcessed", err, startTime)
	}(time.Now())

	return aao.next.ListProcessed(ctx, lastUpdatedAt, limit)
}
