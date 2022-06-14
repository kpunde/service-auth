package adapter

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"serviceAuth/backend/config"
	"serviceAuth/backend/log"
	"serviceAuth/backend/utility"
	"time"
)

var DB *gorm.DB
var _tenantDB map[string]*gorm.DB
var Tenants []string

func getSQLConnectorString(config *config.SqlConfigFormat) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
		config.Ssl,
		config.Timezone)
}

func InitDatabase() {
	var dbConfig config.SqlConfigFormat
	dbConfig = config.GetConfig().SqlConfig

	dsn := getSQLConnectorString(&dbConfig)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utility.CheckError(err)

	con, err := db.DB()
	utility.CheckError(err)

	con.SetMaxIdleConns(dbConfig.MaxIdleConn)
	con.SetMaxOpenConns(dbConfig.MaxOpenConn)
	con.SetConnMaxLifetime(time.Minute * time.Duration(dbConfig.ConnMaxLifetimeMins))
	DB = db
}

func InitTenantDatabase() {
	_tenantDB = make(map[string]*gorm.DB)
	var dbConfig config.SqlConfigFormat
	dbConfig = config.GetConfig().SqlConfig

	dsn := getSQLConnectorString(&dbConfig)

	tService := NewTenantService()
	tenantList, _ := tService.FindActive()

	for _, tenant := range tenantList {
		Tenants = append(Tenants, tenant.ShortName)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   fmt.Sprintf("%v.", tenant.ShortName),
				SingularTable: true,
			},
		})
		utility.CheckError(err)
		_tenantDB[tenant.ShortName] = db
	}
}

func IsSQLDbOkay() bool {
	db, err := DB.DB()
	utility.CheckError(err)

	if err = db.Ping(); err != nil {
		return false
	}

	return true
}

func GetTenantDb(tenantShortName string) *gorm.DB {
	if val, ok := _tenantDB[tenantShortName]; ok {
		return val
	} else {
		tService := NewTenantService()
		tenant, err := tService.FindByShortName(tenantShortName)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("Tenant not found")
			utility.CheckError(err)
			return nil
		}
		if tenant.IsActive {
			var dbConfig config.SqlConfigFormat
			dbConfig = config.GetConfig().SqlConfig

			dsn := getSQLConnectorString(&dbConfig)
			Tenants = append(Tenants, tenant.ShortName)

			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   fmt.Sprintf("%v.", tenant.ShortName),
					SingularTable: true,
				},
			})
			utility.CheckError(err)
			_tenantDB[tenant.ShortName] = db
			return db
		}
		return nil
	}
}
