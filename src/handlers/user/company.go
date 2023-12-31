package user

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCompanies(c *gin.Context) {
	var companies []models.Company
	if err := database.DB.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)

}

func GetCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := database.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)

}
