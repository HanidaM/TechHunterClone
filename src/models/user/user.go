package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleRecruiter Role = "recruiter"
)

type User struct {
	ID        uint `gorm:"primary"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FirstName string         `json:"first_name" gorm:"index;type:varchar(30)" validate:"required,min=2,max=30"`
	LastName  string         `json:"last_name" gorm:"index;type:varchar(30)" validate:"required,min=2,max=30"`
	Email     string         `json:"email" gorm:"uniqueIndex;type:varchar(100)" validate:"required,email"`
	Password  string         `json:"-"`
	Role      Role           `json:"role" gorm:"index;type:varchar(20)"`
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

type Resume struct {
	gorm.Model
	Title          string     `json:"title" gorm:"type:varchar(100)"`
	UserID         uint       `json:"user_id" gorm:"not null;index"` // Foreign key
	Experience     StringList `json:"experience" gorm:"type:text"`
	Grading        string     `json:"grading" gorm:"type:text"`
	Education      StringList `json:"education" gorm:"type:text"`
	Skills         StringList `json:"skills" gorm:"type:text"`
	Projects       StringList `json:"projects" gorm:"type:text"`
	SocialNetworks StringList `json:"social_networks" gorm:"type:text"`
}

type StringList []string

func (list *StringList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSON value")
	}

	return json.Unmarshal(bytes, list)
}

func (list StringList) Value() (driver.Value, error) {
	return json.Marshal(list)
}
