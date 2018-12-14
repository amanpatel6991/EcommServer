package controllers

import (
	"time"
	"github.com/EcommServer/models"
	"crypto/rsa"
	"io/ioutil"
	"fmt"
	"github.com/gin-gonic/gin"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/EcommServer/helper"
	"github.com/go-redis/redis"
)

type UserClaims struct {
	UserProfile       models.User           `json:"user_profile"`
	GoogleUserProfile models.GoogleUser           `json:"google_user_profile"`
	SecretKey         time.Time
	jwt.StandardClaims
}

const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)

var Claims UserClaims

func InitKeys() {
	var err error
	signBytes, err := ioutil.ReadFile(privKeyPath)
	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		fmt.Println("key not read")
		return
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		fmt.Println("key not read")
		return
	}
}

func AuthMiddleWare(redisDb *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {

		SetHeaders(context)
		requestToken := context.Request.Header.Get("token")

		isTokenBlacklisted := models.CheckForBlacklistToken(redisDb, requestToken)

		fmt.Println("isTokenBlacklisted :" , isTokenBlacklisted);

		if isTokenBlacklisted {
			helper.JsonResponse("Unauthorized Access due to Blacklisted Token", "401", context.Writer)
			context.Abort()
		}

		//validate token
		token, err := jwtreq.ParseFromRequestWithClaims(context.Request, jwtreq.AuthorizationHeaderExtractor, &Claims, func(token *jwt.Token) (interface{}, error) {
			return VerifyKey, nil
		})

		if err == nil && token.Valid {
			helper.JsonResponse("Access Granted", "200", context.Writer)
			context.Next()
		} else {
			helper.JsonResponse("Unauthorized Access", "401", context.Writer)
			context.Abort()
		}
	}
}
