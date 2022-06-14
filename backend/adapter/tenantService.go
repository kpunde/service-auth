package adapter

import (
	"serviceAuth/backend/log"
)

type TenantService interface {
	Save(tenantReq *NewTenantRequest) (*Tenant, error)
	Update(tenant *Tenant) (*Tenant, error)
	Delete(tenant *Tenant) error
	FindAll() []Tenant
	FindActive() ([]Tenant, error)
	FindById(id uint) (*Tenant, error)
	FindByShortName(shortName string) (*Tenant, error)
}

type tenantService struct {
	tenantRepository TenantRepository
}

func (r tenantService) Save(tenantReq *NewTenantRequest) (*Tenant, error) {
	log.Debug(tenantReq)
	tenant, err := r.tenantRepository.Save(GetTenantSqlEntity(tenantReq))
	return tenant, err
}

func (r tenantService) Update(tenant *Tenant) (*Tenant, error) {
	tenant, err := r.tenantRepository.Update(tenant)
	return tenant, err
}

func (r tenantService) Delete(tenant *Tenant) error {
	return r.tenantRepository.Delete(tenant)
}

func (r tenantService) FindAll() []Tenant {
	return r.tenantRepository.FindAll()
}

func (r tenantService) FindActive() ([]Tenant, error) {
	return r.tenantRepository.FindActive()
}

func (r tenantService) FindById(id uint) (*Tenant, error) {
	return r.tenantRepository.FindById(id)
}

func (r tenantService) FindByShortName(shortName string) (*Tenant, error) {
	return r.tenantRepository.FindByShortName(shortName)
}

func NewTenantService() TenantService {
	return tenantService{tenantRepository: NewTenantRepository()}
}
