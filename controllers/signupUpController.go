package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/models"
	"fmt"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/EcommServer/helper"
)

//Redirect To Login Page after signup
func Signup(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		SetHeaders(c)
		//getting credentials from headers

		var signupUser models.User
		_ = json.NewDecoder(c.Request.Body).Decode(&signupUser)
		fmt.Println("SENDING IN BODY", signupUser)

		signupEmail := signupUser.Email
		signupPassword := signupUser.Password

		fmt.Println("FORM RETURNS ::", signupEmail, signupPassword)

		//Save New User

		if signupUser.Email == "" || signupUser.Password == "" {
			c.AbortWithStatusJSON(401, "Cannot Signup with blank Email or/and Password")
		}else {
			data, _ := models.CreateUser(db, signupUser)
			if data.Id != 0 {
				helper.JsonResponse(data, "200", c.Writer)
			} else {
				fmt.Println("Error Signing Up . Please try after sometime")
				c.AbortWithStatusJSON(401, "Error Signing Up . Please try after sometime")
			}
		}
	}

}