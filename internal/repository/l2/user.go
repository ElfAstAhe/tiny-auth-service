package l2

import (
	"context"
	"time"

	"github.com/ElfAstAhe/go-service-template/pkg/infra/cache"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserL2Repository struct {
	*repository.BaseCRUDL2Repository[*domain.User, string]
	next      domain.UserRepository
	nameCache cache.Cache[string, *domain.User]
}

var _ domain.UserRepository = (*UserL2Repository)(nil)

func NewUserRepository(
	next domain.UserRepository,
	entityInfo *repository.EntityInfo,
	crudCache cache.Cache[string, *domain.User],
	nameCache cache.Cache[string, *domain.User],
	log logger.Logger,
) *UserL2Repository {
	res := &UserL2Repository{
		BaseCRUDL2Repository: repository.NewBaseCRUDL2Repository[*domain.User, string](
			next,
			entityInfo,
			crudCache,
			2*time.Minute,
			log,
		),
		next: next,
	}

	return res
}

func (u2r *UserL2Repository) Find(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u2r *UserL2Repository) FindByName(ctx context.Context, login string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u2r *UserL2Repository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u2r *UserL2Repository) Create(ctx context.Context, entity *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u2r *UserL2Repository) Change(ctx context.Context, entity *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u2r *UserL2Repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
