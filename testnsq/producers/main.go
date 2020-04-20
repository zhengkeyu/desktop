package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"math/rand"
	"os"
	"strconv"
	"time"
)
type Table struct {
	Id    int
	Uid   int64
	Coins int
	Date  string
}
// NSQ Producer Demo

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

//随机数据
func RandData(uid int64) Table {
	d := Table{
		Uid:   uid,
		Coins: rand.Intn(10000) + 1000,
		Date:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return d
}

var Uid string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1{
		Uid = "100"
	}else{
		Uid = os.Args[1]
	}
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed, err:%v\n", err)
		return
	}
	//回复
	go RunConsumer()
	//
	for {
		time.Sleep(time.Second)
		uid, _ := strconv.Atoi(Uid)
		d := RandData(int64(uid))
		by, _ := json.Marshal(d)
		err := producer.Publish("SaveInfo", by)
		if err != nil {
			fmt.Printf("publish msg to nsq failed, err:%v\n", err)
			continue
		}
		fmt.Println("生产数据: ",d)
	}
}

//消费者
type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	d := &Table{}
	err = json.Unmarshal(msg.Body, d)
	if err != nil {
		panic(err)
	}
	fmt.Println("生产数据回复: ",d)
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
func RunConsumer(){
	fmt.Println("hear...")
	err := initConsumer("SaveInfo-Reply-"+Uid, "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
}