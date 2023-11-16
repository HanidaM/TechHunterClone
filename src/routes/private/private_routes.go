package private

import (
	"TechHunterClone/src/handlers"
	"TechHunterClone/src/handlers/jobseeker"
	"TechHunterClone/src/handlers/recruiter"

	"github.com/gin-gonic/gin"
)

func VacancyRoutes(r *gin.Engine) {
	// Routes for the Vacancy model

	vacancyRoutes := r.Group("/vacancies")
	{
		vacancyRoutes.POST("/", recruiter.CreateVacancy)
		vacancyRoutes.PUT("/:id", recruiter.UpdateVacancy)
		// vacancyRoutes.DELETE("/:id", recruiter.DeleteVacancy)
	}
}
func CompanyRoutes(r *gin.Engine) {
	// Routes for the Company model
	companyRoutes := r.Group("/companies")
	{
		companyRoutes.POST("/", recruiter.CreateCompany)
		companyRoutes.PUT("/:id", recruiter.UpdateCompany)
		companyRoutes.DELETE("/:id", recruiter.DeleteCompany)

	}
}
func ResumeRoutes(r *gin.Engine) {
	// Routes for the Resume model
	resumeRoutes := r.Group("/resumes")
	{
		resumeRoutes.GET("/", handlers.GetResumes)
		resumeRoutes.GET("/:id", handlers.GetResume)
		resumeRoutes.POST("/create", jobseeker.CreateResume)
		resumeRoutes.PUT("/update/:id", jobseeker.UpdateResume)
		resumeRoutes.DELETE("/:id", jobseeker.DeleteResume)
	}
}
