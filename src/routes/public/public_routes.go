package public

import (
	"TechHunterClone/src/handlers/user"
	"github.com/gin-gonic/gin"
)

func VacancyRoutes(r *gin.Engine) {
	vacancyRoutes := r.Group("/vacancies")
	{
		vacancyRoutes.GET("/", user.GetVacancies)
		vacancyRoutes.GET("/:id", user.GetVacancy)
	}
}
func CompanyRoutes(r *gin.Engine) {
	companyRoutes := r.Group("/companies")
	{
		companyRoutes.GET("/", user.GetCompanies)
		companyRoutes.GET("/:id", user.GetCompany)

	}
}
