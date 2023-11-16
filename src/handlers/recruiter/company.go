// src/handlers/recruiter/company.go

package recruiter

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/job"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCompany creates a new company profile
func CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, company)
}

// UpdateCompany updates an existing company profile
func UpdateCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := database.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&company)

	c.JSON(http.StatusOK, company)
}

// DeleteCompany deletes an existing company profile
func DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Company{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
