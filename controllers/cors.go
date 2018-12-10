package controllers

import (
	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	SetHeaders(c)
}

func SetHeaders(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")

}