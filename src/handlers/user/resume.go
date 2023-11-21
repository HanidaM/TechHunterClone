package user

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/user"
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
		c.JSON(http.StatusNotFound, gin.H{"errer": "Resume not found"})
		return
	}
	c.JSON(http.StatusOK, resume)

}
