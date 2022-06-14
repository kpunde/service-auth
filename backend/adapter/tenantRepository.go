package adapter

import (
	"fmt"
	"serviceAuth/backend/log"
)

type TenantRepository interface {
	Save(tenant *Tenant) (*Tenant, error)
	Update(tenant *Tenant) (*Tenant, error)
	Delete(tenant *Tenant) error
	FindAll() []Tenant
	FindById(id uint) (*Tenant, error)
	FindActive() ([]Tenant, error)
	FindByShortName(shortName string) (*Tenant, error)
}

type tenantRepository struct{}

func (t tenantRepository) Save(tenant *Tenant) (*Tenant, error) {
	log.Debug(tenant)

	if dbc := DB.Create(&tenant); dbc.Error != nil {
		return nil, dbc.Error
	}
	DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %v", tenant.ShortName))
	return tenant, nil
}

func (t tenantRepository) Update(tenant *Tenant) (*Tenant, error) {
	log.Debug(tenant)

	if dbc := DB.Updates(&tenant); dbc.Error != nil {
		return nil, dbc.Error
	}
	return tenant, nil
}

func (t tenantRepository) Delete(tenant *Tenant) error {
	if dbc := DB.Delete(&tenant); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (t tenantRepository) FindAll() []Tenant {
	var tenants []Tenant
	DB.Find(&tenants)
	return tenants
}

func (t tenantRepository) FindActive() ([]Tenant, error) {
	var tenants []Tenant
	if dbc := DB.Where("is_active =?", true).Find(&tenants); dbc.Error != nil {
		return nil, dbc.Error
	}
	return tenants, nil
}

func (t tenantRepository) FindById(id uint) (*Tenant, error) {
	var tenant Tenant
	if dbc := DB.Where("id =?", id).First(&tenant); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &tenant, nil
}

func (t tenantRepository) FindByShortName(shortName string) (*Tenant, error) {
	var tenant Tenant
	if dbc := DB.Where("short_name =?", shortName).First(&tenant); dbc.Error != nil {
		log.Debug("Error found")
		return nil, dbc.Error
	}
	return &tenant, nil
}

func NewTenantRepository() TenantRepository {
	return tenantRepository{}
}
