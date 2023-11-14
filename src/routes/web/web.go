// src/routes/web/web.go

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(r *gin.Engine) {

	r.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	r.GET("/auth/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/auth/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

}
