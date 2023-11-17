// src/handlers/jobseeker/resume.go

package jobseeker

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/user"
	services "TechHunterClone/src/services/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetResumes gets all resumes
func GetResumes(c *gin.Context) {
	var resumes []models.Resume
	if err := database.DB.Find(&resumes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resumes)
}

// GetResume gets a single resume by ID
func GetResume(c *gin.Context) {
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	c.JSON(http.StatusOK, resume)
}

// CreateResume creates a new resume

func CreateResume(c *gin.Context) {
	var resume models.Resume

	tokenString, err := c.Cookie("jwt-token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid token"})
		return
	}

	userID, err := services.GetID(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - invalid token"})
		return
	}

	if err := c.BindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resume.UserID = userID

	if err := database.DB.Create(&resume).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resume)
}

// UpdateResume updates an existing resume
func UpdateResume(c *gin.Context) {
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
	id := c.Param("id")

	if err := database.DB.Delete(&models.Resume{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted"})
}
