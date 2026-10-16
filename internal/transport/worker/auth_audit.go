package worker

import (
	"github.com/ElfAstAhe/go-service-template/pkg/transport/worker"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client/dto"
)

type AuthAudit struct {
	*worker.BaseSchedulerDispatcher[*dto.AuthAuditDTO]
}
