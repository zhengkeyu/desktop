// nsq_consumer/main.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"desktop/testmysql/mysql"
)

// NSQ Consumer Demo

// MyHandler 是一个消费者类型
type MyHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	d := &ormsql.Table{}

	err = json.Unmarshal(msg.Body, d)
	if err != nil {
		panic(err)
	}
	_, err = ormsql.Ormer.Insert(d)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("写入数据库,Uid:%v,数据:%v\n", d.Uid, d)
		//发送保存数据成功
		str := "SaveInfo-Reply-" + strconv.Itoa(int(d.Uid))
		err := producer.Publish(str, msg.Body)
		if err != nil {
			fmt.Println("保存数据回复失败,err = ", err)
		}
	}
	return
}

// 初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &MyHandler{
		Title: "消费者",
	}
	c.AddHandler(consumer)
	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil

}

func main() {
	var dbName, password string
	if len(os.Args) != 3 {
		dbName = "zheng"
		password = "abc123"
	} else {
		dbName = os.Args[1]
		password = os.Args[2]
	}
	ormsql.RunMysql(dbName, password)
	err := initConsumer("SaveInfo", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	c := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) // 转发键盘中断信号到c
	<-c                              // 阻塞
}

func init() {
	err := initProducer("127.0.0.1:4150")
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}
}

var producer *nsq.Producer

// 初始化生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}
	return nil
}
