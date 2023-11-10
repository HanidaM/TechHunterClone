// src/routes/web/web.go

package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWebRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mainpage.html", gin.H{})
	})

}
