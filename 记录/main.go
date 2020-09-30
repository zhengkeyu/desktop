package main

import (
	"fmt"
	"time"
)

func main() {
	//loc,_ := time.LoadLocation("Asia/Shanghai")
	//t1,_ := time.ParseInLocation("2006-01-02 15:04:05","2020-09-27 00:00:00",loc)
	//t2,_ := time.ParseInLocation("2006-01-02 15:04:05","2020-09-27 23:59:59",loc)
	//fmt.Println(t1.Unix(),t2.Unix())


   fmt.Println(time.Unix(1596702145,0).Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Now().Unix())
}

type YouKesLogin struct {
	Mac         string
	FirstLogint int64 `pg:"firstlogint"`
	Logint      int64
	From        string
	Pac         string
	Brand       string //手机品牌
	PhoneType   string `pg:"phonetype"` //手机型号
	Total       int64
	tableName   struct{} `pg:"k_youke"`
}
