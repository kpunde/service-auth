package service

import (
	"serviceAuth/backend/entity"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/rest"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/repository"
)

type PermissionService interface {
	Save(requestContext *common.RequestContext, permissionReq *rest.NewPermissionRequest) (*sql.Permission, error)
	SaveFromModel(requestContext *common.RequestContext, permissionReq *sql.Permission) (*sql.Permission, error)
	Update(requestContext *common.RequestContext, permission *sql.Permission) (*sql.Permission, error)
	Delete(requestContext *common.RequestContext, permission *sql.Permission) error
	FindAll(requestContext *common.RequestContext) []sql.Permission
	FindPermissionByTitleList(requestContext *common.RequestContext, titles []string) []sql.Permission
}

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

func (p permissionService) Save(requestContext *common.RequestContext, permissionReq *rest.NewPermissionRequest) (*sql.Permission, error) {
	log.Debug(permissionReq)

	permission, err := p.permissionRepository.Save(requestContext, entity.GetPermissionSqlEntity(permissionReq))
	return permission, err
}

func (p permissionService) SaveFromModel(requestContext *common.RequestContext, permissionReq *sql.Permission) (*sql.Permission, error) {
	log.Debug(permissionReq)

	permission, err := p.permissionRepository.Save(requestContext, permissionReq)
	return permission, err
}

func (p permissionService) Update(requestContext *common.RequestContext, permission *sql.Permission) (*sql.Permission, error) {
	permission, err := p.permissionRepository.Update(requestContext, permission)
	return permission, err
}

func (p permissionService) Delete(requestContext *common.RequestContext, permission *sql.Permission) error {
	return p.permissionRepository.Delete(requestContext, permission)
}

func (p permissionService) FindAll(requestContext *common.RequestContext) []sql.Permission {
	return p.permissionRepository.FindAll(requestContext)
}

func (p permissionService) FindPermissionByTitleList(requestContext *common.RequestContext, titles []string) []sql.Permission {
	return p.permissionRepository.FindPermissionByTitleList(requestContext, titles)
}

func NewPermissionService() PermissionService {
	return &permissionService{
		permissionRepository: repository.NewPermissionRepository(),
	}
}
