package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khalilullahalfaath/jwt-authentication-golang/initializers"
	"github.com/khalilullahalfaath/jwt-authentication-golang/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the email and password from the request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing the password"})
		return
	}

	// create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating the user"})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

}
