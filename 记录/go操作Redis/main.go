package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//先确保本地有安装redis,打开redis-server.exe
	//先下载第三方库 go get github.com/garyburd/redigo/redis

	//连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis失败:", err)
	} else {
		fmt.Println("连接成功")
	}
	defer conn.Close()


	//添加数据,(key-val)
	_, err = conn.Do("Set", "name", "张三")
	//_, err = conn.Do("MSet", "name", "张三","age",10) //也可以一次添加多个
	if err != nil {
		fmt.Println("添加失败")
	}else{
		fmt.Println("添加成功")
	}
	//读取数据
	data, err := redis.String(conn.Do("Get", "name"))//因为是interface{},所以用redis内置的接口转化成string
	//data, err := redis.Strings(conn.Do("MGet", "name","age"))//也可以一次获取多个,返回多个时要用redis.Strings()
	if err != nil {
		fmt.Println("读取失败",err)
	}else{
		fmt.Println("读取数据:",data)
	}

	//redis还可以操作列表，哈希等数据结构
}
