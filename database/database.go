package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"fmt"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis"
)

func InitDb(dbName string) *gorm.DB{
	var db *gorm.DB
	db,err:=gorm.Open("mysql","root:password@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")  //todo see parameters in case of production ??

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

func InitRedisDb(Db int) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password", // no password set
		DB:       Db,  // use default DB
	})

	_, err := client.Ping().Result()
	if err!=nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("redis connected")

	return client
}