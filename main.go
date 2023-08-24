package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khalilullahalfaath/jwt-authentication-golang/controllers"
	"github.com/khalilullahalfaath/jwt-authentication-golang/initializers"
	"github.com/khalilullahalfaath/jwt-authentication-golang/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
