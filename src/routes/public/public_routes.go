package public

import (
	"TechHunterClone/src/handlers"
	"github.com/gin-gonic/gin"
)

func VacancyRoutes(r *gin.Engine) {
	vacancyRoutes := r.Group("/vacancies")
	{
		vacancyRoutes.GET("/", handlers.GetVacancies)
		vacancyRoutes.GET("/:id", handlers.GetVacancy)
	}
}
func CompanyRoutes(r *gin.Engine) {
	companyRoutes := r.Group("/companies")
	{
		companyRoutes.GET("/", handlers.GetCompanies)
		companyRoutes.GET("/:id", handlers.GetCompany)

	}
}
