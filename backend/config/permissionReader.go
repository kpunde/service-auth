package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type PermissionConfigFormat struct {
	Service  string   `json:"service"`
	Function string   `json:"function"`
	Verb     []string `json:"verb"`
}

var permissionConfig []PermissionConfigFormat

func ReadPermissionConfig() {
	var fileName = "./config/configFiles/permissions.json"

	plan, err := ioutil.ReadFile(fileName)
	checkError(err)

	err = json.Unmarshal(plan, &permissionConfig)
	checkError(err)
	fmt.Println(permissionConfig)
}

func GetPermissionConfig() []PermissionConfigFormat {
	if len(permissionConfig) == 0 {
		lock.Lock()
		defer lock.Unlock()
		ReadPermissionConfig()
	}

	return permissionConfig
}
