package entity

import (
	"serviceAuth/backend/entity/rest"
	"serviceAuth/backend/entity/sql"
)

func GetPersonSqlEntity(newPersonRestReq *rest.NewPersonRequest) *sql.Person {
	return &sql.Person{
		Email:    newPersonRestReq.Email,
		IsActive: newPersonRestReq.IsActive,
		PasswordInfo: sql.PasswordInfo{
			PasswordHash: newPersonRestReq.Password,
		},
	}
}
func GetPermissionSqlEntity(newPermissionRestReq *rest.NewPermissionRequest) *sql.Permission {
	return &sql.Permission{
		Title:    newPermissionRestReq.Title,
		Service:  newPermissionRestReq.Service,
		Function: newPermissionRestReq.Function,
		Verb:     newPermissionRestReq.Verb,
	}
}

func GetRoleSqlEntity(newRoleRestReq *rest.NewRoleRequest) *sql.Role {
	return &sql.Role{
		Title: newRoleRestReq.Title,
	}
}

func GetTenantSqlEntity(request *rest.NewTenantRequest) *sql.Tenant {
	return &sql.Tenant{
		IsActive:  request.IsActive,
		Name:      request.Name,
		ShortName: request.ShortName,
	}
}
