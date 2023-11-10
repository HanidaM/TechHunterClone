package private

import (
	"TechHunterClone/src/handlers/jobseeker"
	"TechHunterClone/src/handlers/recruiter"
	"github.com/gin-gonic/gin"
)

func VacancyRoutes(r *gin.Engine) {
	// Create routes for the Vacancy model

	vacancyRoutes := r.Group("/vacancies")
	{
		vacancyRoutes.POST("/", recruiter.CreateVacancy)
		vacancyRoutes.PUT("/:id", recruiter.UpdateVacancy)
		vacancyRoutes.DELETE("/:id", recruiter.DeleteVacancy)
	}
}
func CompanyRoutes(r *gin.Engine) {
	// Create routes for the Company model
	companyRoutes := r.Group("/companies")
	{
		companyRoutes.POST("/", recruiter.CreateCompany)
		companyRoutes.PUT("/:id", recruiter.UpdateCompany)
		companyRoutes.DELETE("/:id", recruiter.DeleteCompany)

	}
}
func ResumeRoutes(r *gin.Engine) {
	// Create routes for the Resume model
	resumeRoutes := r.Group("/resumes")
	{
		resumeRoutes.POST("/", jobseeker.CreateResume)
		resumeRoutes.PUT("/:id", jobseeker.UpdateResume)
		resumeRoutes.DELETE("/:id", jobseeker.DeleteResume)
	}
}
