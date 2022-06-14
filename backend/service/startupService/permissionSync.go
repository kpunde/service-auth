package startupService

import (
	"fmt"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/config"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/service"
	"serviceAuth/backend/utility"
	"strings"
)

func getPermissionSqlEntity(ipPermissions []config.PermissionConfigFormat) (map[string]*sql.Permission, []string) {
	var permMap = make(map[string]*sql.Permission)
	var permTitleList []string

	for _, perm := range ipPermissions {
		for _, verb := range perm.Verb {
			title := fmt.Sprintf("%v::%v::%v",
				strings.ToUpper(perm.Service),
				strings.ToUpper(perm.Function),
				strings.ToUpper(verb))

			permMap[title] = &sql.Permission{
				Service:  perm.Service,
				Function: perm.Function,
				Verb:     verb,
				Title:    title,
			}

			permTitleList = append(permTitleList, title)
		}
	}

	return permMap, permTitleList
}

func difference(slice1, slice2 []string) []string {
	mb := make(map[string]struct{}, len(slice2))
	for _, x := range slice2 {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range slice1 {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func PermissionSync() {
	permissionService := service.NewPermissionService()
	permissionFromConfigMap, permissionFromConfigList := getPermissionSqlEntity(config.GetPermissionConfig())

	for _, tenant := range adapter.Tenants {
		log.Info(fmt.Sprintf("Starting PermissionSync for %v", tenant))
		reqContext := &common.RequestContext{Tenant: tenant}
		permissionInDb := permissionService.FindPermissionByTitleList(reqContext, permissionFromConfigList)
		var permissionTitleInDB []string
		for _, perm := range permissionInDb {
			permissionTitleInDB = append(permissionTitleInDB, perm.Title)
		}

		permissionTitleToBeCreated := difference(permissionFromConfigList, permissionTitleInDB)
		for _, permission := range permissionTitleToBeCreated {
			_, err := permissionService.SaveFromModel(reqContext, permissionFromConfigMap[permission])
			utility.CheckError(err)
		}
	}
}
