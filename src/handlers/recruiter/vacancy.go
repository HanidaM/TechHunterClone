// src/handlers/recruiter/vacancy.go

package recruiter

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/job"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateVacancy creates a new job vacancy
func CreateVacancy(c *gin.Context) {
	// Implement the creation logic
	var vacancy models.Vacancy
	if err := c.BindJSON(&vacancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&vacancy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, vacancy)
}

// UpdateVacancy updates an existing job vacancy
func UpdateVacancy(c *gin.Context) {
	// Implement the update logic
	var vacancy models.Vacancy
	id := c.Param("id")

	if err := database.DB.First(&vacancy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	if err := c.BindJSON(&vacancy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&vacancy)

	c.JSON(http.StatusOK, vacancy)
}

// DeleteVacancy deletes an existing job vacancy
func DeleteVacancy(c *gin.Context) {
	// Implement the deletion logic
	id := c.Param("id")

	if err := database.DB.Delete(&models.Vacancy{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted"})

}