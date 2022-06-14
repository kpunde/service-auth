package service

import (
	"fmt"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/utility"
)

type LoginService interface {
	Login(requestContext *common.RequestContext, email string, password string) (*sql.Person, bool)
}

type loginService struct {
	personService PersonService
}

func (l loginService) Login(requestContext *common.RequestContext, email string, password string) (*sql.Person, bool) {
	user, err := l.personService.FindByEmail(requestContext, email)
	if err != nil {
		return nil, false
	}

	pwd, err := l.personService.FindPasswordServices(requestContext, user)
	if err != nil {
		return nil, false
	}

	log.Debug(fmt.Sprintf("Password from database is : %v", pwd.PasswordHash))
	log.Debug(fmt.Sprintf("Password from request is : %v", password))

	authResult := utility.IsPasswordMatch(pwd.PasswordHash, password)
	log.Debug(fmt.Sprintf("Auth result is: %v", authResult))
	return user, authResult
}

func NewLoginService() LoginService {
	return &loginService{personService: NewPersonService()}
}
