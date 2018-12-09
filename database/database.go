package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb(dbName string) *gorm.DB{
	var db *gorm.DB
	db,err:=gorm.Open("mysql","root:password@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")

	if err!=nil {
		log.Fatal(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}else{
		fmt.Println("connected")
		return db
	}

}