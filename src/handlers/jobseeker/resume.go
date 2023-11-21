// src/handlers/jobseeker/resume.go

package jobseeker

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/user"
	services "TechHunterClone/src/services/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResumes(c *gin.Context) {
	var resumes []models.Resume
	if err := database.DB.Find(&resumes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resumes)
}

func GetResume(c *gin.Context) {
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	c.JSON(http.StatusOK, resume)
}

func CreateResume(c *gin.Context) {
	var resume models.Resume

	if err := c.BindJSON(&resume); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resume data", "details": err.Error()})
		return
	}

	tokenString, err := c.Cookie("Authorise")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required: No token found"})
		return
	}

	userID, err := services.GetID(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed: Invalid token"})
		return
	}
	resume.UserID = userID

	if err := database.DB.Create(&resume).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create resume", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume created successfully", "resumeID": resume.ID})
}

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

func DeleteResume(c *gin.Context) {
	id := c.Param("id")

	if err := database.DB.Delete(&models.Resume{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted"})
}
