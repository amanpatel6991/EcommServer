package main

import (
	"github.com/EcommServer/controllers"
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/database"
	"github.com/EcommServer/models"
)

func main() {
	db := database.InitDb("ecomm")
	defer db.Close()

	//db.DropTable(&models.User{} ,&models.GoogleUser{}, &models.Address{})

	db.Debug().AutoMigrate(&models.User{}, &models.GoogleUser{}, &models.Address{})

	//Insert sample user FOR TESTING DB
	//add1:=models.Address{PhoneNumber: "9811399095" , AddressLine1:"line 1" , AddressLine2: "line 2"}
	//add2 := models.Address{PhoneNumber: "8368269946", AddressLine1:"line 11", AddressLine2: "line 21"}
	//add3 := models.Address{PhoneNumber: "9999777765", AddressLine1:"line 11", AddressLine2: "line 21"}
	//create:=db.Debug().Create(
	//	&models.User{
	//		FirstName:"sdAman",LastName:"Patelsad",Email:"aman@gmail.com",Password:"password",SignedInSource:"manual",
	//		Addresses: []models.Address{}}).RowsAffected
	//
	//fmt.Println("created ::" , create)
	//
	//var userDataSample []models.User
	//query:=db.Debug().Preload("Addresses").Find(&userDataSample).RowsAffected
	//fmt.Println("test Db data ::" , query , userDataSample)

	//DELETE
	//var user models.User
	//db.Debug().Preload("Addresses").First(&user , 7)
	//a:=db.Debug().Delete(&user , 7).RowsAffected
	//fmt.Println("test delete " , a , user)

	//UPDATE
	//user:=models.User{Email: "samplenew@gmail.com"}
	//update:=db.Debug().Model(&models.User{}).Where("id=?",8).Update(&user).RowsAffected
	//db.Find(&user , 8)
	//fmt.Println("test updte :" , update , user)

	//Test model functions
	//1. Create
	//user:=models.User{
	//	FirstName:"Shubham",LastName:"Garg",Email:"sample@gmail.com",Password:"password",SignedInSource:"manual",
	//	Addresses: []models.Address{add3}}
	//
	//user , respMsg := models.CreateUser(db , user)
	//fmt.Println("test db create :" , user , respMsg)
	//2. Update
	//user := models.User{
	//	FirstName:"Shubham1", LastName:"garg", Email:"sample@gmail.com", Password:"password", SignedInSource:"manual",
	//	//Addresses: []models.Address{
	//	//	{PhoneNumber: "9999777765", AddressLine1:"line 11", AddressLine2: "line 21"},
	//	//},
	//}
	//user, respMsg := models.UpdateUserById(db, user , 14)
	//fmt.Println("test db update :", user, respMsg)
	//3. Query
	//var  user []models.User
	//user, respMsg := models.GetAllUserswithAssociations(db)
	//fmt.Println("test db query :", user, respMsg)
	//4. DELETE
	//user, respMsg := models.DeleteUserById(db ,10)
	//fmt.Println("test db delete :", user, respMsg)

	controllers.InitKeys()

	api := gin.Default()

	api.OPTIONS("/login", controllers.Cors)
	api.POST("/login", controllers.Login)

	api.OPTIONS("/googleLogin", controllers.Cors)
	api.POST("/googleLogin", controllers.GoogleLogin)

	api.OPTIONS("/api/v1/:routes", controllers.Cors)
	group := api.Group("/api/v1/")

	group.Use(controllers.AuthMiddleWare())
	{
		group.GET("sample", controllers.GetSampleData(nil))
	}

	api.Run(":5000")

}

