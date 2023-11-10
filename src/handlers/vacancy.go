package handlers

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/job"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetVacancies fetches all job vacancies
func GetVacancies(c *gin.Context) {
	// Add database querying logic here to fill vacancies slice
	var vacancies []models.Vacancy
	if err := database.DB.Find(&vacancies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vacancies)
}

// GetVacancy fetches a single job vacancy by ID
func GetVacancy(c *gin.Context) {
	var vacancy models.Company
	id := c.Param("id")

	if err := database.DB.First(&vacancy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, vacancy)

}
