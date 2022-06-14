package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"serviceAuth/backend/adapter"
	"serviceAuth/backend/entity/common"
	"serviceAuth/backend/entity/sql"
	"serviceAuth/backend/log"
)

//var db = adapter.DB

type PersonRepository interface {
	Save(requestContext *common.RequestContext, person *sql.Person) (*sql.Person, error)
	Update(requestContext *common.RequestContext, person *sql.Person) (*sql.Person, error)
	UpdatePassword(requestContext *common.RequestContext, person *sql.Person) error
	UpdateRolesForPersonId(requestContext *common.RequestContext, personId uuid.UUID, roles []sql.Role) (*sql.Person, error)
	GetPersonPassword(requestContext *common.RequestContext, person *sql.Person) (*sql.PasswordInfo, error)
	Delete(requestContext *common.RequestContext, person *sql.Person) error
	FindAll(requestContext *common.RequestContext) []sql.Person
	FindById(requestContext *common.RequestContext, id uuid.UUID) (*sql.Person, error)
	FindByEmail(requestContext *common.RequestContext, email string) (*sql.Person, error)
	FindRolesForPersonId(requestContext *common.RequestContext, id uuid.UUID) ([]sql.Role, error)
}

type personDb struct {
}

func (d personDb) FindByEmail(requestContext *common.RequestContext, email string) (*sql.Person, error) {
	tenant := requestContext.Tenant
	var person sql.Person
	if dbc := adapter.GetTenantDb(tenant).Where("email =?", email).First(&person); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &person, nil
}

func (d personDb) Save(requestContext *common.RequestContext, person *sql.Person) (*sql.Person, error) {
	tenant := requestContext.Tenant
	log.Debug(person)

	if dbc := adapter.GetTenantDb(tenant).Create(&person); dbc.Error != nil {
		return nil, dbc.Error
	}
	return person, nil
}

func (d personDb) Update(requestContext *common.RequestContext, person *sql.Person) (*sql.Person, error) {
	tenant := requestContext.Tenant
	if dbc := adapter.GetTenantDb(tenant).Model(&person).Omit("Roles").Updates(&person); dbc.Error != nil {
		return nil, dbc.Error
	}
	if len(person.Roles) != 0 {
		tx := adapter.GetTenantDb(tenant).Begin()
		var lst []map[string]interface{}

		for _, role := range person.Roles {
			lst = append(lst, map[string]interface{}{"person_id": person.Id, "role_id": role.ID})
		}

		if err := tx.Exec(fmt.Sprintf("DELETE FROM %s.person_roles WHERE person_id=?", tenant), person.Id).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Table(fmt.Sprintf("%s.person_roles", tenant)).Create(lst).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
	} else {
		err := adapter.GetTenantDb(tenant).Model(&person).Association("Roles").Clear()
		if err != nil {
			return nil, err
		}
	}

	return person, nil
}

func (d personDb) UpdatePassword(requestContext *common.RequestContext, person *sql.Person) error {
	tenant := requestContext.Tenant
	if dbc := adapter.GetTenantDb(tenant).Session(&gorm.Session{FullSaveAssociations: true}).Save(person); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (d personDb) GetPersonPassword(requestContext *common.RequestContext, person *sql.Person) (*sql.PasswordInfo, error) {
	tenant := requestContext.Tenant
	var pwd sql.PasswordInfo
	if dbc := adapter.GetTenantDb(tenant).Where("person_id =?", person.Id).First(&pwd); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &pwd, nil
}

func (d personDb) Delete(requestContext *common.RequestContext, person *sql.Person) error {
	tenant := requestContext.Tenant
	if dbc := adapter.GetTenantDb(tenant).Delete(&person); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (d personDb) FindAll(requestContext *common.RequestContext) []sql.Person {
	tenant := requestContext.Tenant
	var persons []sql.Person
	adapter.GetTenantDb(tenant).Table(fmt.Sprintf("%s.person", tenant)).Select("id", "email", "is_active").Scan(&persons)
	return persons
}

func (d personDb) FindById(requestContext *common.RequestContext, id uuid.UUID) (*sql.Person, error) {
	tenant := requestContext.Tenant
	var person sql.Person
	if dbc := adapter.GetTenantDb(tenant).Preload("Roles").Where("id =?", id).First(&person); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &person, nil
}

func (p personDb) FindRolesForPersonId(requestContext *common.RequestContext, id uuid.UUID) ([]sql.Role, error) {
	tenant := requestContext.Tenant
	var roles []sql.Role
	if dbc := adapter.GetTenantDb(tenant).
		Joins(fmt.Sprintf("INNER JOIN %s.person_roles ON role.id=person_roles.role_id", tenant)).
		Where("person_roles.person_id = ?", id).
		Find(&roles); dbc.Error != nil {
		return nil, dbc.Error
	}
	return roles, nil
}

func (p personDb) UpdateRolesForPersonId(requestContext *common.RequestContext, personId uuid.UUID, roles []sql.Role) (*sql.Person, error) {
	tenant := requestContext.Tenant
	person, err := p.FindById(requestContext, personId)
	if err != nil {
		return nil, err
	}
	person.Roles = roles
	if dbc := adapter.GetTenantDb(tenant).Model(&person).Updates(&person); dbc.Error != nil {
		return nil, dbc.Error
	}

	return person, nil
}

func NewPersonRepository() PersonRepository {
	return &personDb{}
}
