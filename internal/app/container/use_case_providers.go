package container

import (
	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/db"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

func (ucc *UseCaseContainer) providerTM() (any, error) {
	dbInst, err := container.GetInstance[db.DB](InstanceDB)
	if err != nil {
		return nil, errs.NewContainerError(ucc.GetName(), "provider: retrieve instance failed", err)
	}

	return db.NewTxManager(dbInst), nil
}

func (ucc *UseCaseContainer) providerChangeKeysUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerChangePasswordUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerLoginUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerLoginSimpleUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerProfileUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerRegisterUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerRoleAdminDeleteUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerRoleAdminGetUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerRoleAdminGetByNameUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerRoleAdminListUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerRoleAdminSaveUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerUserAdminDeleteUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerUserAdminGetUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerUserAdminGetByNameUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerUserAdminListUC() (any, error) {
	// ToDo: implement

	return nil, nil
}

func (ucc *UseCaseContainer) providerUserAdminSaveUC() (any, error) {
	// ToDo: implement

	return nil, nil
}
