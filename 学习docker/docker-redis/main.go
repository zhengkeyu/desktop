package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
)

//redis客户端
var redisDb *redis.Client

func init() {
	redisDb = redis.NewClient(&redis.Options{
		Addr: "172.17.0.2:6379",
	})
	_, err := redisDb.Ping().Result() //连接到redis服务器
	if err != nil {
		panic(err)
	}
}

func main() {
	//s := redisDb.Get("zheng")
	//fmt.Println(s.Result())
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		paths := strings.Split(path, "/")
		if len(paths) < 3 || paths[1] != "get" && paths[1] != "set" {
			writer.Write([]byte("commit error"))
			return
		}
		if paths[1] == "get" {
			dbmsg := redisDb.Get(paths[2])
			resp := "get result:"
			s, err := dbmsg.Result()
			if err != nil {
				resp += paths[2] + ",no find"
			} else {
				resp += "key = " + s
			}
			writer.Write([]byte(resp))
		} else if paths[1] == "set" && len(paths) == 4 {
			redisDb.Set(paths[2], paths[3], 0)
			writer.Write([]byte(fmt.Sprintf("set success,key = %v,value = %v", paths[2], paths[3])))
		} else {
			writer.Write([]byte("commit error"))
		}
	})
	http.ListenAndServe(":8500", nil)
}
