package user

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/job"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func GetVacancies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var vacancies []models.Vacancy

	if err := database.DB.Offset(offset).Limit(pageSize).Find(&vacancies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var totalCount int64
	database.DB.Model(&models.Vacancy{}).Count(&totalCount)
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	c.JSON(http.StatusOK, gin.H{
		"vacancies":   vacancies,
		"totalPages":  totalPages,
		"currentPage": page,
	})
}

func GetVacancy(c *gin.Context) {
	var vacancy models.Company
	id := c.Param("id")

	if err := database.DB.First(&vacancy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, vacancy)

}
