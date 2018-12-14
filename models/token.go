package models

import (
	"github.com/go-redis/redis"
	"time"
	"fmt"
)

type Token struct {
	Token string                      `json:"token"`
}

//Interacting with Redis

func BlacklistToken(client *redis.Client, token string) error{
	err := client.Set(token, "B", time.Second * 1000).Err()     //B -> Blacklisted , delete token after its expiration time 1000seconds
	if err != nil {
		fmt.Println("BlacklistToken Method :", err)
	} else {
		fmt.Println("REDIS dB :", token, " is BLACKLISTED !")
	}

	return err
}

//Returns true if token is blacklisted
func CheckForBlacklistToken(client *redis.Client, token string) bool {
	_ , err := client.Get(token).Result()                         //B -> Blacklisted
	if err != nil {  // some error occurs , token not found in redis db
		fmt.Println("CheckForBlacklistToken Method :", err)
		return false
	}

	fmt.Println("CheckForBlacklistToken Method : token not blacklisted !!")
	return true
}