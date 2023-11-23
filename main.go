package main

import (
	_ "TechHunterClone/docs"
	"TechHunterClone/src/database"
	"TechHunterClone/src/routes/auth"
	"TechHunterClone/src/routes/private"
	"TechHunterClone/src/routes/public"
	"TechHunterClone/src/routes/web"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRoutes(r *gin.Engine) {

	auth.AuthRoutes(r)
	private.CompanyRoutes(r)
	private.ResumeRoutes(r)
	private.VacancyRoutes(r)
	public.CompanyRoutes(r)
	public.VacancyRoutes(r)
	public.ResumeRoutes(r)

	web.RegisterWebRoutes(r)

}
func main() {

	r := gin.Default()
	database.ConnectDB()
	SetUpRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.LoadHTMLGlob("src/templates/*.html")

	r.Run(":8080")

}
