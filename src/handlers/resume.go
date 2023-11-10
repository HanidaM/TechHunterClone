package handlers

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetResumes fetches all resumes
func GetResumes(c *gin.Context) {
	var resumes []models.Resume
	if err := database.DB.Find(&resumes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resumes)
}

// GetResume fetches a single resume by ID
func GetResume(c *gin.Context) {
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errer": "Resume not found"})
		return
	}
	c.JSON(http.StatusOK, resume)

}