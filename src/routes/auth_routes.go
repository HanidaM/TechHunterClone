// src/routes/public/auth_routes.go

package routes

import (
	"TechHunterClone/src/middlewares/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", auth.RegisterHandler)
		authRoutes.POST("/login", auth.LoginHandler)
		authRoutes.POST("/logout", auth.LogoutHandler)
	}
}
