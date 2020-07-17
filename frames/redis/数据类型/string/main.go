package main

import (
	"fmt"
	. "redis/connect"
)

//字符串
func main() {
	defer CloseRedis()

	var stringKey = "stringkey"
	//设置
	_, err := Db.Set(stringKey, "value", 0).Result()
	if err != nil {
		panic(err)
	}

	//获取
	res, err := Db.Get(stringKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("获取到%v的值为:%v\n", stringKey, res)

	//删除
	//可以传多个key
	//返回删除的数量和可能的错误
	count,err := Db.Del(stringKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("删除了%v条数据\n",count)


}
