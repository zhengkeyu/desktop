package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis客户端
var redisDb *redis.Client

func init() {
	redisDb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	_, err := redisDb.Ping().Result() //连接到redis服务器
	if err != nil {
		panic(err)
	}
}

func main() {
	s := redisDb.Get("zheng")
	fmt.Println(s.Result())
	a := redisDb.Keys("*")
	fmt.Println(a.Result())
}
