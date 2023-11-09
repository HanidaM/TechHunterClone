package models

import (
	"gorm.io/gorm"
	"time"
)

// Vacancy represents a job listing.
type Vacancy struct {
	gorm.Model
	Title    string `json:"title" gorm:"index;type:varchar(100)"`
	Salary   string `json:"salary" gorm:"type:varchar(50)"`
	Company  string `json:"company" gorm:"type:varchar(100)"`
	Location string `json:"location" gorm:"type:varchar(100)"`
	UserID   uint   `json:"user_id" gorm:"not null"` // Foreign key
}

// VacancyDescription represents the details of a job vacancy.
type VacancyDescription struct {
	gorm.Model
	Description string `json:"description" gorm:"type:text"`
	Keywords    string `json:"keywords" gorm:"type:text"`   // JSON encoded string
	TechStack   string `json:"tech_stack" gorm:"type:text"` // JSON encoded string
	URL         string `json:"url" gorm:"type:text"`
	VacancyID   uint   `json:"vacancy_id" gorm:"not null"` // Foreign key
}

// Company represents an employer's company information.
type Company struct {
	gorm.Model
	Name        string `json:"company_name" gorm:"type:varchar(100)"`
	Description string `json:"company_description" gorm:"type:text"`
	Headcount   int    `json:"headcount"`
	Type        string `json:"company_type" gorm:"type:varchar(50)"`
	Industry    string `json:"company_industry" gorm:"type:varchar(50)"`
	TechStacks  string `json:"tech_stacks" gorm:"type:text"` // JSON encoded string
	LogoURL     string `json:"logo_url" gorm:"type:text"`
	WebsiteURL  string `json:"website_url" gorm:"type:text"`
	ReviewedAt  *time.Time
	UserID      uint `json:"user_id" gorm:"not null"` // Foreign key
}
