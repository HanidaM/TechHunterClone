// src/routes/web/web.go

package web

import (
	"TechHunterClone/src/database"
	models "TechHunterClone/src/models/user"
	services "TechHunterClone/src/services/auth_service"
	"errors"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(r *gin.Engine) {
	r.GET("/auth/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/auth/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	r.GET("/main", func(c *gin.Context) {
		token, err := c.Cookie("Authorise")
		var loggedIn bool
		var role string
		if err == nil {
			user, err := services.VerifyToken(token)
			if err == nil {
				loggedIn = true
				role = string(user.Role)
			}
		}

		c.HTML(http.StatusOK, "main.html", gin.H{
			"LoggedIn": loggedIn,
			"Role":     role,
		})
	})

	r.GET("/main/profile/user", func(c *gin.Context) {
		token, err := c.Cookie("Authorise")
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "token not found"})
			return

		}
		user, err := services.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized - invalid token"})
			return
		}

		var resume models.Resume
		resumeExists := true
		if err := database.DB.Where("user_id = ?", user.ID).First(&resume).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				resumeExists = false
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				return
			}
		}

		c.HTML(http.StatusOK, "userpage.html", gin.H{
			"ResumeExists":   resumeExists,
			"Title":          resume.Title,
			"Experience":     resume.Experience,
			"Grading":        resume.Grading,
			"Education":      resume.Education,
			"Skills":         resume.Skills,
			"Projects":       resume.Projects,
			"SocialNetworks": resume.SocialNetworks,
		})
	})

}
