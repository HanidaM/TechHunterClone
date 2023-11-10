
# TechHunterClone

A brief description of the project

This project is a job searching platform that connects job seekers with recruiters. It allows job seekers to find and apply for jobs, while recruiters can post job vacancies and manage applications.



## Tech Stack

- **Programming Language:** Go
- **Web Framework:** Gin
- **ORM:** Gorm
- **Database:** PostgreSQL
- **Caching:** Redis
- **API Documentation:** Swagger


## Project Structure

Below is a high-level overview of the project structure:
```
├───src
│   ├───config
│   │   # Configuration files for the application
│   ├───database
│   │   ├───migrations
│   │   │   # Database migration scripts for evolving the database schema
│   │   └───recruiter
│   │       # Database-related code and scripts specific to the recruiter module
│   ├───handlers
│   │   ├───jobseeker
│   │   │   # Handlers and controllers specific to jobseeker actions
│   │   └───recruiter
│   │       # Handlers and controllers specific to recruiter actions
│   ├───middlewares
│   │   ├───auth
│   │   │   # Middleware code for authentication (e.g., JWT, sessions)
│   │   └───roles
│   │       # Middleware code for role-based access control
│   ├───models
│   │   ├───application
│   │   │   # Models and data structures related to job applications
│   │   ├───job
│   │   │   # Models and data structures related to job postings
│   │   └───user
│   │       # Models and data structures related to users, including jobseekers and recruiters
│   ├───routes
│   │   ├───public
│   │   │   # Publicly accessible routes, such as job listings and company profiles
│   │   ├───private
│   │   │   # Protected routes for authenticated users
│   │   └───web
│   │       # Web-specific routes for rendering HTML templates
│   ├───services
│   │   ├───auth_service
│   │   │   # Authentication-related services and functions
│   │   ├───notification
│   │   │   # Services for sending notifications or alerts
│   │   └───search
│   │       # Services related to searching for jobs and resumes
│   └───templates
│       # Templates for rendering HTML pages
└───utils
    ├───response
    │   # Utilities for formatting standardized response data
    └───validators
        # Utilities for input validation


```
