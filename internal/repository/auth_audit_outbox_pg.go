package repository

import (
	librepo "github.com/ElfAstAhe/go-service-template/pkg/repository/v2"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/repository/postgres"
)

//goland:noinspection DuplicatedCode
func (aao *AuthAuditOutbox) pgStatementsBuilder() *librepo.Statements[*domain.AuthAuditOutbox, string] {
	return librepo.NewStatements[*domain.AuthAuditOutbox, string]().
		WithStatement(librepo.StatementFind, librepo.NewStatement(postgres.SqlAuthAuditOutboxFind, aao.mapDefault)).
		WithStatement(librepo.StatementList, librepo.NewStatement(postgres.SqlAuthAuditOutboxList, aao.mapDefault)).
		WithStatement(librepo.StatementCreate, librepo.NewStatement(postgres.SqlAuthAuditOutboxCreate, aao.mapDefault)).
		WithStatement(librepo.StatementChange, librepo.NewStatement(postgres.SqlAuthAuditOutboxChange, aao.mapDefault)).
		WithStatement(librepo.StatementDelete, librepo.NewStatement[*domain.AuthAuditOutbox, string](postgres.SqlAuthAuditOutboxDelete, nil)).
		WithStatement(statementAuthAuditOutboxGetTail, librepo.NewStatement(postgres.SqlAuthAuditOutboxGetTail, aao.mapTail)).
		WithStatement(statementAuthAuditOutboxListPending, librepo.NewStatement(postgres.SqlAuthAuditOutboxListPending, aao.mapDefault)).
		WithStatement(statementAuthAuditOutboxListProcessed, librepo.NewStatement(postgres.SqlAuthAuditOutboxListProcessed, aao.mapDefault)).
		WithStatement(statementAuthAuditOutboxChangeStatus, librepo.NewStatement(postgres.SqlAuthAuditOutboxChangeStatus, aao.mapDefault))
}
