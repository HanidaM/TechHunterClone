// src/handlers/jobseeker/resume.go

package jobseeker

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateResume creates a new resume
func CreateResume(c *gin.Context) {
	// Implement the creation logic
	var resume models.Resume
	if err := c.BindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&resume).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resume)
}

// UpdateResume updates an existing resume
func UpdateResume(c *gin.Context) {
	// Implement the update logic
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	if err := c.BindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&resume)

	c.JSON(http.StatusOK, resume)
}

// DeleteResume deletes an existing resume
func DeleteResume(c *gin.Context) {
	// Implement the deletion logic
	id := c.Param("id")

	if err := database.DB.Delete(&models.Resume{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted"})
}
