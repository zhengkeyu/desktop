package connect

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

var Db *redis.Client

func init() {
	Db = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	str, err := Db.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis ping success ", str)
}

func CloseRedis() {
	err := Db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("redis close success")
	}
}
