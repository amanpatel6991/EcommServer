package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/database"
	"github.com/EcommServer/models"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
	"log"
	"github.com/EcommServer/responseFormatter"
	"fmt"
)

func Login(c *gin.Context) {

	//getting credentials from headers
	name1 := c.Request.Header.Get("username")
	pwd1 := c.Request.Header.Get("password")

	fmt.Println("FORM RETURNS ::", name1, pwd1)

	db := database.InitDb("ecomm")
	defer db.Close()

	//validate user credentials
	var user []models.User
	db.Model(models.User{}).Find(&user)

	flag := -1                                                            //exit condition
	for _, v := range user {

		if strings.ToLower(name1) == v.Email && strings.ToLower(pwd1) == v.Password {
			//set claims
			Claims = UserClaims{
				models.User{Id:v.Id, FirstName:v.FirstName, LastName:v.LastName , SignedInSource: v.SignedInSource},
				time.Now(), //to generate unique token everytime  //rand.Intn(10000),
				jwt.StandardClaims{
					Issuer: "testing_administrator",
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodRS256, Claims)
			ss, err := token.SignedString(SignKey)

			if err != nil {
				c.String(404, "Not Found")
				fmt.Fprintln(c.Writer, err)
				log.Printf("err: %+v\n", err)
				return
			}

			c.Writer.Header().Set("status", "200")
			response := Token{Token:ss}
			responseFormatter.JsonResponse(response, "200 OK", c.Writer)
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

