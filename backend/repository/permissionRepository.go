package repository

import (
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
)

type PermissionRepository interface {
	Save(requestContext *common.RequestContext, permission *sql.Permission) (*sql.Permission, error)
	Update(requestContext *common.RequestContext, permission *sql.Permission) (*sql.Permission, error)
	Delete(requestContext *common.RequestContext, permission *sql.Permission) error
	FindAll(requestContext *common.RequestContext) []sql.Permission
	FindPermissionByTitleList(requestContext *common.RequestContext, titles []string) []sql.Permission
}

type permissionDb struct {
}

func (p permissionDb) Save(requestContext *common.RequestContext, permission *sql.Permission) (*sql.Permission, error) {
	tenant := requestContext.Tenant
	log.Debug(permission)

	if dbc := adapter.GetTenantDb(tenant).Create(&permission); dbc.Error != nil {
		return nil, dbc.Error
	}
	return permission, nil
}

func (p permissionDb) Update(requestContext *common.RequestContext, permission *sql.Permission) (*sql.Permission, error) {
	tenant := requestContext.Tenant
	if dbc := adapter.GetTenantDb(tenant).Save(&permission); dbc.Error != nil {
		return nil, dbc.Error
	}
	return permission, nil
}

func (p permissionDb) Delete(requestContext *common.RequestContext, permission *sql.Permission) error {
	tenant := requestContext.Tenant
	if dbc := adapter.GetTenantDb(tenant).Delete(&permission); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (p permissionDb) FindAll(requestContext *common.RequestContext) []sql.Permission {
	tenant := requestContext.Tenant
	var permissions []sql.Permission
	adapter.GetTenantDb(tenant).Find(&permissions)
	return permissions
}

func (p permissionDb) FindPermissionByTitleList(requestContext *common.RequestContext, titles []string) []sql.Permission {
	tenant := requestContext.Tenant
	var permissions []sql.Permission
	adapter.GetTenantDb(tenant).Where("title IN ?", titles).Find(&permissions)
	return permissions
}

func NewPermissionRepository() PermissionRepository {
	return &permissionDb{}
}
