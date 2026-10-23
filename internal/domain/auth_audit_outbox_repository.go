package domain

import (
	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	libextdom "github.com/ElfAstAhe/go-service-template/pkg/ext/domain"
)

// AuthAuditOutboxRepository определяет специализированный интерфейс доменного уровня
// для управления персистентным состоянием записей транзакционного журнала Outbox.
type AuthAuditOutboxRepository interface {
	libdom.CRUDRepository[*AuthAuditOutbox, string]
	libextdom.TailRepository[string]
	libextdom.OutboxRepository[*AuthAuditOutbox, string]
}
