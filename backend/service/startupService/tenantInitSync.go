package startupService

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/config"
	"serviceAuth/backend/log"
)

func InitTenantSync() {
	tenantService := adapter.NewTenantService()
	tenantFromConfigMap := config.GetConfig().InitTenant
	_, err := tenantService.FindByShortName(tenantFromConfigMap.ShortName)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		tenant := &adapter.NewTenantRequest{
			IsActive:  tenantFromConfigMap.IsActive,
			Name:      tenantFromConfigMap.Name,
			ShortName: tenantFromConfigMap.ShortName,
		}
		_, err := tenantService.Save(tenant)
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create init tenants \n%v", err))
		}
	}
}
