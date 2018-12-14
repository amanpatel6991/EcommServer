package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/EcommServer/helper"
	"github.com/go-redis/redis"
	"fmt"
	"github.com/EcommServer/models"
)

func SignOut(redisDb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		SetHeaders(c)

		requestToken := c.Request.Header.Get("token")
		fmt.Println("inLogout :", requestToken)

		err := models.BlacklistToken(redisDb, requestToken)

		if err == nil {
			c.Writer.Header().Set("status", "200")
			helper.JsonResponse("User Token is now blacklisted !", "200", c.Writer)
		} else {
			fmt.Println("Error Logging out (error in blacklisting token)")
			c.AbortWithStatusJSON(401, "Error Logging out , in blacklisting token")
		}

	}

}