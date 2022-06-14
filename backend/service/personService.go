package service

import (
	"github.com/google/uuid"
	"serviceAuth/backend/entity"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/rest"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
	"serviceAuth/backend/repository"
	"serviceAuth/backend/utility"
)

type PersonService interface {
	Save(requestContext *common.RequestContext, personReq *rest.NewPersonRequest) (*sql.Person, error)
	Update(requestContext *common.RequestContext, person *sql.Person) (*sql.Person, error)
	Delete(requestContext *common.RequestContext, person *sql.Person) error
	FindAll(requestContext *common.RequestContext) []sql.Person
	FindById(requestContext *common.RequestContext, id string) (*sql.Person, error)
	FindByEmail(requestContext *common.RequestContext, email string) (*sql.Person, error)
	FindPasswordServices(requestContext *common.RequestContext, person *sql.Person) (*sql.PasswordInfo, error)
	FindRolesForPersonId(requestContext *common.RequestContext, id string) ([]sql.Role, error)
	UpdatePassword(requestContext *common.RequestContext, id string, password *sql.PasswordInfo) error
	UpdateRolesForPersonId(requestContext *common.RequestContext, id string, roles []sql.Role) (*sql.Person, error)
}

type personService struct {
	personRepository repository.PersonRepository
}

func (i *personService) FindRolesForPersonId(requestContext *common.RequestContext, id string) ([]sql.Role, error) {
	uuidId, err := uuid.Parse(id)
	utility.CheckError(err)

	return i.personRepository.FindRolesForPersonId(requestContext, uuidId)
}

func (i *personService) UpdateRolesForPersonId(requestContext *common.RequestContext, id string, roles []sql.Role) (*sql.Person, error) {
	uuidId, err := uuid.Parse(id)
	utility.CheckError(err)

	return i.personRepository.UpdateRolesForPersonId(requestContext, uuidId, roles)
}

func (i *personService) Update(requestContext *common.RequestContext, person *sql.Person) (*sql.Person, error) {
	person, err := i.personRepository.Update(requestContext, person)
	return person, err
}

func (i *personService) Delete(requestContext *common.RequestContext, person *sql.Person) error {
	return i.personRepository.Delete(requestContext, person)
}

func (i *personService) Save(requestContext *common.RequestContext, personRequest *rest.NewPersonRequest) (*sql.Person, error) {
	log.Debug(personRequest)

	personRequest.Password = utility.HashAndSalt(personRequest.Password)

	person, err := i.personRepository.Save(requestContext, entity.GetPersonSqlEntity(personRequest))
	return person, err
}

func (i *personService) FindAll(requestContext *common.RequestContext) []sql.Person {
	return i.personRepository.FindAll(requestContext)
}

func (i *personService) FindById(requestContext *common.RequestContext, id string) (*sql.Person, error) {
	uuidId, err := uuid.Parse(id)
	utility.CheckError(err)
	return i.personRepository.FindById(requestContext, uuidId)
}

func (i *personService) FindByEmail(requestContext *common.RequestContext, email string) (*sql.Person, error) {
	return i.personRepository.FindByEmail(requestContext, email)
}

func (i *personService) FindPasswordServices(requestContext *common.RequestContext, user *sql.Person) (*sql.PasswordInfo, error) {
	return i.personRepository.GetPersonPassword(requestContext, user)
}

func (i *personService) UpdatePassword(requestContext *common.RequestContext, id string, password *sql.PasswordInfo) error {
	pwd := utility.HashAndSalt(password.PasswordHash)
	uuidId, err := uuid.Parse(id)

	person, err := i.personRepository.FindById(requestContext, uuidId)
	if err != nil {
		return err
	}

	person.PasswordInfo.PasswordHash = pwd
	return i.personRepository.UpdatePassword(requestContext, person)
}

func NewPersonService() PersonService {
	return &personService{
		personRepository: repository.NewPersonRepository(),
	}
}
