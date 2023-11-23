package public

import (
	"TechHunterClone/src/handlers/user"

	"github.com/gin-gonic/gin"
)

func VacancyRoutes(r *gin.Engine) {
	vacancyRoutes := r.Group("/vacancies")
	{
		vacancyRoutes.GET("/", user.GetVacancies)
	}
}
func CompanyRoutes(r *gin.Engine) {
	companyRoutes := r.Group("/companies")
	{
		companyRoutes.GET("/", user.GetCompanies)
		companyRoutes.GET("/:id", user.GetCompany)

	}
}
func ResumeRoutes(r *gin.Engine) {
	resumeRoutes := r.Group("/resumes")
	{
		resumeRoutes.GET("/", user.GetResumes)
		resumeRoutes.GET("/:id", user.GetResume)
	}
}
