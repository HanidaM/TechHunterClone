package main

import (
	"TechHunterClone/src/database"
	"TechHunterClone/src/routes/auth"
	"TechHunterClone/src/routes/private"
	"TechHunterClone/src/routes/public"
	"TechHunterClone/src/routes/web"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	// routes.MainRoutes(r)
	// routes.RecruiterRoutes(r)
	auth.AuthRoutes(r)
	private.CompanyRoutes(r)
	private.ResumeRoutes(r)
	private.VacancyRoutes(r)
	public.CompanyRoutes(r)
	public.VacancyRoutes(r)
	web.RegisterWebRoutes(r)

}
func main() {

	r := gin.Default()
	database.ConnectDB()
	SetUpRoutes(r)
	r.LoadHTMLGlob("src/templates/*.html")

	r.Run(":8080")
}
