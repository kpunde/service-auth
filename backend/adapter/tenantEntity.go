package adapter

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model
	IsActive  bool   `json:"is_active" gorm:"type:boolean;column:is_active"`
	Name      string `json:"name" gorm:"type:varchar(1024);column:name"`
	ShortName string `json:"short_name" gorm:"type:varchar(1024);column:short_name;UNIQUE"`
}

type NewTenantRequest struct {
	IsActive  bool   `json:"is_active" binding:"required"`
	Name      string `json:"name" binding:"required"`
	ShortName string `json:"short_name" binding:"required"`
}

func GetTenantSqlEntity(request *NewTenantRequest) *Tenant {
	return &Tenant{
		IsActive:  request.IsActive,
		Name:      request.Name,
		ShortName: request.ShortName,
	}
}
