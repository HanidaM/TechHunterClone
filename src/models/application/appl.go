package models

import "gorm.io/gorm"

// Resume represents a job seeker's resume.
type Resume struct {
	gorm.Model
	Title          string `json:"title" gorm:"type:varchar(100)"`
	UserID         uint   `json:"user_id" gorm:"not null;index"` // Foreign key
	FirstName      string `json:"first_name" gorm:"type:varchar(30)" validate:"required,min=2,max=30"`
	LastName       string `json:"last_name" gorm:"type:varchar(30)" validate:"required,min=2,max=30"`
	Email          string `json:"email" gorm:"type:varchar(100)" validate:"required,email"`
	Experience     string `json:"experience" gorm:"type:text"`
	Education      string `json:"education" gorm:"type:text"`
	Skills         string `json:"skills" gorm:"type:text"`
	Projects       string `json:"projects" gorm:"type:text"`
	SocialNetworks string `json:"social_networks" gorm:"type:text"` // JSON encoded string
}
