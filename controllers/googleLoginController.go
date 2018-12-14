package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/models"
	"fmt"
	"encoding/json"
	"log"
	"time"
	"github.com/EcommServer/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

func GoogleLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		SetHeaders(c)
		//getting credentials from headers

		var googleLoginUser models.GoogleUser
		_ = json.NewDecoder(c.Request.Body).Decode(&googleLoginUser)
		fmt.Println("SENDING IN BODY", googleLoginUser)

		loginUserToken := googleLoginUser.Token

		fmt.Println("FORM RETURNS ::", loginUserToken)
		//
		flag := -1                                                            //exit condition
		if loginUserToken != "" {
			//set claims
			Claims = UserClaims{
				models.User{},
				models.GoogleUser{Id:googleLoginUser.Id, Name:googleLoginUser.Name, Email: googleLoginUser.Email , SignedInSource: "google"},
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
			//http.SetCookie(c.Writer, &http.Cookie{
			//	Name:    "auth_token",
			//	Value:   ss,
			//	Expires: time.Now().Add(10 * time.Second),                  //expire this on any time logout
			//})

			c.Writer.Header().Set("status", "200")
			response := models.Token{Token:ss}
			helper.JsonResponse(response, "200", c.Writer)

			flag = 1

			//Save GoogleUser Records
			models.CreateGoogleUser(db , googleLoginUser)

			return
		} else {
			flag = -1
		}
		//}
		if flag == -1 {
			fmt.Println("Error Logging in")
			c.AbortWithStatusJSON(401, "Error Logging in")
		}
	}

}