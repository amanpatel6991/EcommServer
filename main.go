package main

import (
	"github.com/EcommServer/controllers"
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/database"
)


func main() {
	db := database.InitDb("ecomm")
	defer db.Close()

	controllers.InitKeys()

	api := gin.Default()

	api.OPTIONS("/login", controllers.Cors)
	api.POST("/login", controllers.Login)

	api.OPTIONS("/googleLogin", controllers.Cors)
	api.POST("/googleLogin", controllers.GoogleLogin)

	api.OPTIONS("/api/v1/:routes", controllers.Cors)
	group:=api.Group("/api/v1/")

	group.Use(controllers.AuthMiddleWare())
	{
		group.GET("sample", controllers.GetSampleData(nil))
	}

	api.Run(":5000")

}

