package repository

import (
	"fmt"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
)

type RoleRepository interface {
	Save(requestContext *common.RequestContext, role *sql.Role) (*sql.Role, error)
	Update(requestContext *common.RequestContext, role *sql.Role) (*sql.Role, error)
	Delete(requestContext *common.RequestContext, role *sql.Role) error
	FindAll(requestContext *common.RequestContext) []sql.Role
	FindById(requestContext *common.RequestContext, id int) (*sql.Role, error)
	FindPermissionsForRoleId(requestContext *common.RequestContext, roleIds []int) ([]sql.Permission, error)
}

type roleDb struct {
}

func (p roleDb) Save(requestContext *common.RequestContext, role *sql.Role) (*sql.Role, error) {
	tenant := requestContext.Tenant
	log.Debug(role)

	if dbc := adapter.GetTenantDb(tenant).Create(&role); dbc.Error != nil {
		return nil, dbc.Error
	}
	return role, nil
}

func (p roleDb) Update(requestContext *common.RequestContext, role *sql.Role) (*sql.Role, error) {
	tenant := requestContext.Tenant

	if dbc := adapter.GetTenantDb(tenant).Model(&role).Omit("Permissions").Updates(&role); dbc.Error != nil {
		return nil, dbc.Error
	}
	if len(role.Permissions) != 0 {
		tx := adapter.GetTenantDb(tenant).Begin()
		var lst []map[string]interface{}

		for _, perm := range role.Permissions {
			lst = append(lst, map[string]interface{}{"role_id": role.ID, "permission_id": perm.ID})
		}

		if err := tx.Exec(fmt.Sprintf("DELETE FROM %s.role_permissions WHERE role_id=?", tenant), role.ID).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Table(fmt.Sprintf("%s.role_permissions", tenant)).Create(lst).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	} else {
		err := adapter.GetTenantDb(tenant).Model(&role).Association("Permissions").Clear()
		if err != nil {
			return nil, err
		}
	}

	return p.FindById(requestContext, int(role.ID))
}

func (p roleDb) Delete(requestContext *common.RequestContext, role *sql.Role) error {
	tenant := requestContext.Tenant
	if dbc := adapter.GetTenantDb(tenant).Delete(&role); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (p roleDb) FindAll(requestContext *common.RequestContext) []sql.Role {
	tenant := requestContext.Tenant
	var roles []sql.Role
	adapter.GetTenantDb(tenant).Find(&roles)
	return roles
}

func (p roleDb) FindById(requestContext *common.RequestContext, id int) (*sql.Role, error) {
	tenant := requestContext.Tenant
	var role sql.Role
	if dbc := adapter.GetTenantDb(tenant).Preload("Permissions").Where("id =?", id).First(&role); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &role, nil
}

func (p roleDb) FindPermissionsForRoleId(requestContext *common.RequestContext, ids []int) ([]sql.Permission, error) {
	tenant := requestContext.Tenant
	var permissions []sql.Permission
	if dbc := adapter.GetTenantDb(tenant).
		Distinct("permission.title").
		Joins(fmt.Sprintf("INNER JOIN %s.role_permissions ON permission.id=role_permissions.permission_id", tenant)).
		Where("role_permissions.role_id IN ?", ids).
		Find(&permissions); dbc.Error != nil {
		return nil, dbc.Error
	}
	return permissions, nil
}

func NewRoleRepository() RoleRepository {
	return &roleDb{}
}
