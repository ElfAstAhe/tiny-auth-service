package container

import (
    "fmt"

    "github.com/ElfAstAhe/go-service-template/pkg/container"
    "github.com/ElfAstAhe/go-service-template/pkg/db"
    "github.com/ElfAstAhe/go-service-template/pkg/errs"
    "github.com/ElfAstAhe/go-service-template/pkg/helper"
    "github.com/ElfAstAhe/go-service-template/pkg/utils"
    "github.com/ElfAstAhe/tiny-auth-service/internal/domain"
    "github.com/ElfAstAhe/tiny-auth-service/internal/repository/metrics"
    "github.com/ElfAstAhe/tiny-auth-service/internal/repository/postgres"
    "github.com/ElfAstAhe/tiny-auth-service/internal/repository/trace"
)

func (rc *RepositoryContainer) providerRoleRepo() (any, error) {
    dbInst, err := container.GetInstance[db.DB](InstanceDB)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    res, err := postgres.NewRolePgRepository(dbInst, dbInst)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceRoleRepo), err)
    }

    return res, nil
}

func (rc *RepositoryContainer) providerRoleMetricsRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.RoleRepository](InstanceRoleRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return metrics.NewRoleMetricsRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerRoleTraceRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.RoleRepository](InstanceRoleMetricsRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return trace.NewRoleTraceRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerRoleAdminRepo() (any, error) {
    dbInst, err := container.GetInstance[db.DB](InstanceDB)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    res, err := postgres.NewRoleAdminPgRepository(dbInst, dbInst)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceRoleAdminRepo), err)
    }

    return res, nil
}

func (rc *RepositoryContainer) providerRoleAdminMetricsRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return metrics.NewRoleAdminMetricsRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerRoleAdminTraceRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.RoleAdminRepository](InstanceRoleAdminMetricsRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return trace.NewRoleAdminTraceRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserRepo() (any, error) {
    hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }
    cipherHelperInst, err := container.GetInstance[helper.Cipher](InstanceDataCipherHelper)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }
    dbInst, err := container.GetInstance[db.DB](InstanceDB)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }
    userRolesRepoInst, err := container.GetInstance[domain.UserRolesRepository](InstanceUserRolesTraceRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    res, err := postgres.NewUserPgRepository(
        dbInst, dbInst,
        hashCipherInst, cipherHelperInst,
        userRolesRepoInst,
    )
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceUserRolesRepo), err)
    }

    return res, nil
}

func (rc *RepositoryContainer) providerUserMetricsRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserRepository](InstanceUserRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return metrics.NewUserMetricsRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserTraceRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserRepository](InstanceUserMetricsRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return trace.NewUserTraceRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserAdminRepo() (any, error) {
    hashCipherInst, err := container.GetInstance[utils.Cipher](InstanceHashCipher)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }
    cipherHelperInst, err := container.GetInstance[helper.Cipher](InstanceDataCipherHelper)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }
    dbInst, err := container.GetInstance[db.DB](InstanceDB)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }
    userRolesAdminRepoInst, err := container.GetInstance[domain.UserRolesAdminRepository](InstanceUserRolesAdminRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    res, err := postgres.NewUserAdminPgRepository(
        dbInst, dbInst,
        cipherHelperInst, hashCipherInst,
        userRolesAdminRepoInst,
    )
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceUserRolesAdminRepo), err)
    }

    return res, nil
}

func (rc *RepositoryContainer) providerUserAdminMetricsRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return metrics.NewUserAdminMetricsRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserAdminTraceRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserAdminRepository](InstanceUserAdminMetricsRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return trace.NewUserAdminTraceRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserRolesRepo() (any, error) {
    dbInst, err := container.GetInstance[db.DB](InstanceDB)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    res, err := postgres.NewUserRolesPgRepository(dbInst, dbInst)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceUserRolesRepo), err)
    }

    return res, nil
}

func (rc *RepositoryContainer) providerUserRolesMetricsRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserRolesRepository](InstanceUserRolesRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return metrics.NewUserRolesMetricsRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserRolesTraceRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserRolesRepository](InstanceUserRolesMetricsRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return trace.NewUserRolesTraceRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserRolesAdminRepo() (any, error) {
    dbInst, err := container.GetInstance[db.DB](InstanceDB)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    res, err := postgres.NewUserRolesAdminPgRepository(dbInst, dbInst)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), fmt.Sprintf("provider: create %s instance failed", InstanceUserRolesAdminRepo), err)
    }

    return res, nil
}

func (rc *RepositoryContainer) providerUserRolesAdminMetricsRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserRolesAdminRepository](InstanceUserRolesAdminRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return metrics.NewUserRolesAdminMetricsRepository(repoInst), nil
}

func (rc *RepositoryContainer) providerUserRolesAdminTraceRepo() (any, error) {
    repoInst, err := container.GetInstance[domain.UserRolesAdminRepository](InstanceUserRolesAdminMetricsRepo)
    if err != nil {
        return nil, errs.NewContainerError(rc.GetName(), "provider: retrieve instance failed", err)
    }

    return trace.NewUserRolesAdminTraceRepository(repoInst), nil
}
