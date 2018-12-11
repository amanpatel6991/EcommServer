package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetSampleData(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context){
		SetHeaders(c)
		fmt.Println("access !! :" , Claims)
		fmt.Println("access db !! :" , db)

	}
}
