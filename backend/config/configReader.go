package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var config Format
var lock = &sync.Mutex{}

type SqlConfigFormat struct {
	Host                string `json:"host"`
	Port                string `json:"port"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	Database            string `json:"database"`
	Ssl                 string `json:"ssl"`
	Timezone            string `json:"timezone"`
	MaxIdleConn         int    `json:"max_idle_conn"`
	MaxOpenConn         int    `json:"max_open_conn"`
	ConnMaxLifetimeMins int    `json:"conn_max_lifetime_mins"`
}

type LogConfigFormat struct {
	Level string `json:"level"`
}

type InitTenantFormat struct {
	ShortName string `json:"short_name"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
}

type Format struct {
	SqlConfig  SqlConfigFormat  `json:"sql"`
	LogConfig  LogConfigFormat  `json:"log"`
	InitTenant InitTenantFormat `json:"init_tenant"`
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func ReadConfig() {
	selectedEnv := os.Getenv("ENV")
	var fileName string

	switch selectedEnv {
	case "dev":
		fileName = "dev_config.json"
	case "stage":
		fileName = "stage_config.json"
	case "prod":
		fileName = "prod_config.json"
	}

	plan, err := ioutil.ReadFile(fileName)
	checkError(err)

	err = json.Unmarshal(plan, &config)
	checkError(err)
}

func GetConfig() *Format {
	if config == (Format{}) {
		lock.Lock()
		defer lock.Unlock()
		ReadConfig()
	}

	return &config
}
