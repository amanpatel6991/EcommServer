package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/models"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
	"log"
	"fmt"
	"github.com/EcommServer/helper"
	"net/http"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c * gin.Context){

		SetHeaders(c)
		//getting credentials from headers

		var loginUser models.LoginInfo
		_ = json.NewDecoder(c.Request.Body).Decode(&loginUser)
		fmt.Println("SENDING IN BODY", loginUser.Email, loginUser.Password)

		loginEmail := loginUser.Email
		loginPassword := loginUser.Password

		fmt.Println("FORM RETURNS ::", loginEmail, loginPassword)

		//validate user credentials
		var user []models.User
		db.Model(models.User{}).Find(&user)

		flag := -1                                                            //exit condition
		for _, v := range user {

			if strings.ToLower(loginEmail) == v.Email && strings.ToLower(loginPassword) == v.Password {
				//if strings.ToLower(loginUsername) == "aman@gmail.com" && strings.ToLower(loginPassword) == "password" {
				//set claims
				Claims = UserClaims{
					models.User{Id:v.Id, FirstName:v.FirstName, LastName:v.LastName, Email:v.Email, SignedInSource: v.SignedInSource},
					//models.User{Id: 1, FirstName: "Aman", LastName: "Patel",Email: "aman@gmail.com", SignedInSource: "manual"},
					models.GoogleUser{},
					time.Now(), //to generate unique token everytime  //rand.Intn(10000),
					jwt.StandardClaims{
						Issuer: "testing_administrator",
						ExpiresAt: time.Now().Add(time.Second * 1000).Unix(), //set this to large value (15-20 hrs) after testing
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims)
				ss, err := token.SignedString(SignKey)

				if err != nil {
					c.String(404, "Not Found")
					log.Printf("err: %+v\n", err)
					return
				}

				//set session cookie with token info
				http.SetCookie(c.Writer, &http.Cookie{
					Name:    "auth_token",
					Value:   ss,
					Expires: time.Now().Add(10 * time.Second), //expire this on any time logout
				})

				c.Writer.Header().Set("status", "200")
				response := models.Token{Token:ss}
				helper.JsonResponse(response, "200", c.Writer)

				flag = 1
				return
			} else {
				flag = -1
			}
		}
		if flag == -1 {
			fmt.Println("Error Logging in")
			c.AbortWithStatusJSON(401, "Error Logging in")
		}
	}

}