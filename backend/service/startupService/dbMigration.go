package startupService

import (
	"fmt"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
)

func MigrateDB() {

	err := adapter.DB.AutoMigrate(
		&sql.Tenant{},
	)
	if err != nil {
		panic(err)
	}

	for _, tenant := range adapter.Tenants {
		log.Info(fmt.Sprintf("Migrating for tenant: %v", tenant))
		adapter.DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %v", tenant))

		err := adapter.GetTenantDb(tenant).AutoMigrate(
			&sql.Person{},
			&sql.PasswordInfo{},
			&sql.Permission{},
			&sql.Role{},
		)
		if err != nil {
			panic(err)
		}
	}
}

func MigrateTenantDB() {
	err := adapter.DB.AutoMigrate(
		&sql.Tenant{},
	)
	if err != nil {
		panic(err)
	}
}
