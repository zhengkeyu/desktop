package main

import (
	"desktop/frames/redis/connect"
)

func main() {
	defer connect.CloseRedis()

	//err := connect.Db.Expire("b_kickoutinfo", time.Second*5).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//_,err := connect.Db.SAdd("b_kickoutinfo","u1").Result()
	//if err != nil {
	//	panic(err)
	//}
	//
	//_,err = connect.Db.SAdd("b_kickoutinfo","u1").Result()
	//if err != nil {
	//	panic(err)
	//}

	//err := connect.Db.ZIncrBy("k1",2,"z1").Err()
	//if err != nil {
	//	panic(err)
	//}

	//err := AddBlindBoxOpenHistory("1", "zky开盲盒获取到现金100000000万,21")
	//if err != nil {
	//	panic(err)
	//}

	//for i := 1; i <= 4; i++ {
	//	connect.Db.LPush("k1", strconv.Itoa(i))
	//	len, err := connect.Db.LLen("k1").Result()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println("len:",len)
	//}

	//err := connect.Db.IncrBy("incr111", 10).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//res,err := connect.Db.Get("incr111").Result()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(res)
	//err = connect.Db.IncrBy("incr111", 2).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//res,err = connect.Db.Get("incr111").Result()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(res)

}
