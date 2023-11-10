package auth

import (
	models "TechHunterClone/src/models/user"
	services "TechHunterClone/src/services/auth_service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func errorResponse(c *gin.Context, status int, err error) {
	c.AbortWithError(status, err)
}

func RegisterHandler(c *gin.Context) {
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")
	roleStr := c.PostForm("role")

	var role models.Role
	switch roleStr {
	case "user":
		role = models.RoleUser
	case "recruiter":
		role = models.RoleRecruiter
	}

	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Role:      role,
	}
	err := services.ValidateUser(user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	if password != confirmPassword {
		errorResponse(c, http.StatusBadRequest, errors.New("passwords do not match"))
		return
	}

	hashedPassword, err := services.HashPassword(password) // Hash the password
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, errors.New("failed to hash password"))
		return
	}

	user.Password = hashedPassword // Set the hashed password

	db, ok := c.Get("db")
	if !ok {
		errorResponse(c, http.StatusInternalServerError, errors.New("database connection not found"))
		return
	}

	dbInstance, ok := db.(*gorm.DB)
	if !ok {
		errorResponse(c, http.StatusInternalServerError, errors.New("invalid database connection"))
		return
	}

	err = dbInstance.Create(user).Error
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/login")
}

func LoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	if email == "" || password == "" {
		errorResponse(c, http.StatusBadRequest, errors.New("email and password are required"))
		return
	}

	db, ok := c.Get("db")
	if !ok {
		errorResponse(c, http.StatusInternalServerError, errors.New("database connection not found in context"))
		return
	}

	dbInstance, ok := db.(*gorm.DB)
	if !ok {
		errorResponse(c, http.StatusInternalServerError, errors.New("invalid database connection in context"))
		return
	}

	var user models.User
	if err := dbInstance.Where("email = ?", email).First(&user).Error; err != nil {
		errorResponse(c, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	if err := services.CheckPassword(user.Password, password); err != nil {
		errorResponse(c, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	accessToken, err := services.CreateToken(&user)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("token", accessToken, int(time.Hour.Seconds()*48), "/", "", false, true)

	c.Redirect(http.StatusFound, "/main")
}

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}
