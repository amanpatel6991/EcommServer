package main

import (
	"github.com/EcommServer/database"
	"github.com/EcommServer/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
	db := database.InitDb("ecomm")
	defer db.Close()

	controllers.InitKeys()

	api := gin.Default()

	api.POST("/login", controllers.Login)

	group:=api.Group("/api/v1/")

	group.Use(controllers.AuthMiddleWare())
	{
		group.GET("sample", controllers.GetSampleData)
	}

	api.Run(":5000")

}
