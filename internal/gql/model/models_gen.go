// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"com.vitanexus/main/internal/orm/models"
)

type NewUser struct {
	models.Base
	FirstName string `json:"firstName" `
}

type UpdateUser struct {
	models.Base
	FirstName string `json:"firstName" `
}

type User struct {
	models.Base
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}
