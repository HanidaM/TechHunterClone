package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type Vacancy struct {
	gorm.Model
	JobTitle   string     `json:"job_title" gorm:"index;type:varchar(100)"`
	Company    string     `json:"company" gorm:"type:varchar(100)"`
	Salary     string     `json:"salary" gorm:"type:varchar(50)"`
	Location   string     `json:"location" gorm:"type:varchar(100)"`
	TechStacks StringList `json:"tech_stacks" gorm:"type:text"`
	VacancyURL string     `json:"vacancy_url" gorm:"type:text"`
	Source     string     `json:"source" gorm:"type:text"`
}

type Company struct {
	gorm.Model
	Name        string     `json:"company_name" gorm:"type:varchar(100)"`
	Description string     `json:"company_description" gorm:"type:text"`
	Headcount   int        `json:"headcount"`
	Type        string     `json:"company_type" gorm:"type:varchar(50)"`
	Industry    string     `json:"company_industry" gorm:"type:varchar(50)"`
	TechStacks  StringList `json:"tech_stacks" gorm:"type:text"`
	LogoURL     string     `json:"logo_url" gorm:"type:text"`
	WebsiteURL  string     `json:"website_url" gorm:"type:text"`
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
