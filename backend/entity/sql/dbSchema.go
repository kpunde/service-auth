package sql

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Add uuid-ossp extension for postgres database, e.g
//CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Tenant struct {
	gorm.Model
	IsActive  bool   `json:"is_active" gorm:"type:boolean;column:is_active"`
	Name      string `json:"name" gorm:"type:varchar(1024);column:name"`
	ShortName string `json:"short_name" gorm:"type:varchar(1024);column:short_name;UNIQUE"`
}

//Add uuid-ossp extension for postgres database, e.g
//CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Person struct {
	Id           uuid.UUID    `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Email        string       `json:"email" binding:"required,email" gorm:"type:varchar(320);UNIQUE;column:email"`
	IsActive     bool         `json:"is_active" gorm:"type:boolean;column:is_active"`
	PasswordInfo PasswordInfo `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Roles        []Role       `json:"roles" gorm:"many2many:person_roles;constraint:OnDelete:CASCADE"`
	Updated      int64        `gorm:"autoUpdateTime"`
	Created      int64        `gorm:"autoCreateTime"`
}

//func (Person) TableName() string {
//	return "person"
//}

func (person *Person) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	id := uuid.New()
	tx.Statement.SetColumn("Id", id)
	return
}

type PasswordInfo struct {
	PersonId     uuid.UUID `json:"person_id" gorm:"type:uuid;column:person_id;primary_key"`
	PasswordHash string    `json:"password" gorm:"type:varchar(1024);column:password_hash;not null"`
}

//func (PasswordInfo) TableName() string {
//	return "password_info"
//}

type Permission struct {
	gorm.Model
	Service  string `json:"service" gorm:"type:varchar(1024);not null"`
	Function string `json:"function" gorm:"type:varchar(1024);not null"`
	Verb     string `json:"verb" gorm:"type:varchar(1024);not null"`
	Title    string `json:"title" gorm:"type:varchar(1024);not null;UNIQUE"`
}

//func (Permission) TableName() string {
//	return "permission"
//}

type Role struct {
	gorm.Model
	Title       string       `json:"title" gorm:"type:varchar(1024);not null"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE"`
}

//func (Role) TableName() string {
//	return "role"
//}
