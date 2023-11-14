package models

import (
	"errors"
	"gorm.io/gorm"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleRecruiter Role = "recruiter"
)

// User represents a system user with roles.
type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"index;type:varchar(30)" validate:"required,min=2,max=30"`
	LastName  string `json:"last_name" gorm:"index;type:varchar(30)" validate:"required,min=2,max=30"`
	Email     string `json:"email" gorm:"uniqueIndex;type:varchar(100)" validate:"required,email"`
	Password  string `json:"-"`
	Role      Role   `json:"role" gorm:"index;type:varchar(20)"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email required")
	}

	if u.Password == "" {
		return errors.New("password required")
	}

	return nil
}
