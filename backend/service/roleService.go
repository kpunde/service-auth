package service

import (
	"serviceAuth/backend/entity"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/rest"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/repository"
)

type RoleService interface {
	Save(requestContext *common.RequestContext, roleReq *rest.NewRoleRequest) (*sql.Role, error)
	Update(requestContext *common.RequestContext, role *sql.Role) (*sql.Role, error)
	Delete(requestContext *common.RequestContext, role *sql.Role) error
	FindAll(requestContext *common.RequestContext) []sql.Role
	FindById(requestContext *common.RequestContext, id int) (*sql.Role, error)
	FindPermissionsForRoleId(requestContext *common.RequestContext, roleIds []int) ([]sql.Permission, error)
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func (r roleService) Save(requestContext *common.RequestContext, roleReq *rest.NewRoleRequest) (*sql.Role, error) {
	log.Debug(roleReq)
	role, err := r.roleRepository.Save(requestContext, entity.GetRoleSqlEntity(roleReq))
	return role, err
}

func (r roleService) Update(requestContext *common.RequestContext, role *sql.Role) (*sql.Role, error) {
	role, err := r.roleRepository.Update(requestContext, role)
	return role, err
}

func (r roleService) Delete(requestContext *common.RequestContext, role *sql.Role) error {
	return r.roleRepository.Delete(requestContext, role)
}

func (r roleService) FindAll(requestContext *common.RequestContext) []sql.Role {
	return r.roleRepository.FindAll(requestContext)
}

func (r roleService) FindById(requestContext *common.RequestContext, id int) (*sql.Role, error) {
	return r.roleRepository.FindById(requestContext, id)
}

func (r roleService) FindPermissionsForRoleId(requestContext *common.RequestContext, roleIds []int) ([]sql.Permission, error) {
	return r.roleRepository.FindPermissionsForRoleId(requestContext, roleIds)
}

func NewRoleService() RoleService {
	return roleService{roleRepository: repository.NewRoleRepository()}
}
