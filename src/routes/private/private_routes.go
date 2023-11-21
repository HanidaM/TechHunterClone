package private

import (
	"TechHunterClone/src/handlers/jobseeker"
	"TechHunterClone/src/handlers/recruiter"

	"github.com/gin-gonic/gin"
)

func VacancyRoutes(r *gin.Engine) {

	vacancyRoutes := r.Group("/vacancies")
	{
		vacancyRoutes.POST("/", recruiter.CreateVacancy)
		vacancyRoutes.PUT("/:id", recruiter.UpdateVacancy)
		// vacancyRoutes.DELETE("/:id", recruiter.DeleteVacancy)
	}
}
func CompanyRoutes(r *gin.Engine) {
	companyRoutes := r.Group("/companies")
	{
		companyRoutes.POST("/", recruiter.CreateCompany)
		companyRoutes.PUT("/:id", recruiter.UpdateCompany)
		companyRoutes.DELETE("/:id", recruiter.DeleteCompany)

	}
}
func ResumeRoutes(r *gin.Engine) {
	resumeRoutes := r.Group("/resumes")
	{
		resumeRoutes.GET("/", jobseeker.GetResume)
		resumeRoutes.GET("/:id", jobseeker.GetResume)
		resumeRoutes.POST("/create", jobseeker.CreateResume)
		resumeRoutes.PUT("/update/:id", jobseeker.UpdateResume)
		resumeRoutes.DELETE("/:id", jobseeker.DeleteResume)
	}
}
