
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
├───src
│   ├───assets
│   │   └───emails          # Templates for emails to be sent by the application
│   ├───config
│   ├───database
│   │   ├───migrations
│   │   └───seeds           # Initial data seeding for roles, admin users, etc.
│   ├───handlers
│   │   ├───jobseeker       # Handlers specific to jobseeker actions
│   │   └───recruiter       # Handlers specific to recruiter actions
│   ├───middlewares
│   │   ├───auth            # Authentication middleware (e.g., JWT, sessions)
│   │   └───roles           # Role-based access control middleware
│   ├───models
│   │   ├───user            # User model, interfaces for jobseekers and recruiters
│   │   ├───job             # Job posting model
│   │   └───application     # Job application model
│   ├───routes
│   │   ├───public          # Publicly accessible routes (e.g., job listings, company profiles)
│   │   ├───private         # Protected routes for authenticated users
│   │   └───admin           # Administrative routes for platform management
│   ├───services
│   │   ├───search          # Services for job and resume search functionalities
│   │   └───notification    # Services for sending notifications or alerts
│   └───utils
│       ├───validator       # Input validation utilities
│       └───response        # Standardized response formatting utilities
