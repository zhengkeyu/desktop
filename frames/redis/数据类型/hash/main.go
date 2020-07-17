package main

import (
	"fmt"
	. "redis/connect"
)

//hash
//类似map[string]interface{}
func main() {
	defer CloseRedis()

	var hashKey = "hashkey"

	//设置
	count, err := Db.HSet(hashKey, "name", "zheng", "age", "27").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hash表:%v，设置了%v个字段\n", hashKey, count)

	//字段不存在才会设置
	_,err = Db.HSetNX(hashKey,"newfile","keyu").Result()
	if err != nil {
		panic(err)
	}

	//查询
	//以map返回
	m, err := Db.HGetAll(hashKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("获取到的是map类型的hash表数据:", m)

	//指定要获取的多个字段,以slice返回
	s,err := Db.HMGet(hashKey,"name","age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("获取到的是slice类型的hash表数据:", s)

	//获取单个字段，以string返回
	//Db.HGet().Result()

	//删除
	//指定key和字段名来删除每个字段
	//Db.HDel().Result()
}
