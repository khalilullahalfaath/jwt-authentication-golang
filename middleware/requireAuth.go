package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/khalilullahalfaath/jwt-authentication-golang/initializers"
	"github.com/khalilullahalfaath/jwt-authentication-golang/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("in middleware")

	// get the token from the header cookie
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// decode and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check if the signing method is HMAC
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacsamplesecret is the secret key
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user
		var user models.User
		result := initializers.DB.Where("id = ?", claims["sub"]).First(&user)

		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// set the user in the context
		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
	}

}
