package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v9"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//var Pg = &pg.Options{
//	Addr:     "101.200.142.211:5432",
//	User:     "kong",
//	Password: "kong",
//	Database: "postgres",
//}

var Pg = &pg.Options{
	Addr:     "123.207.85.242:5433",
	User:     "postgres",
	Password: "2932615qian",
	Database: "postgres",
}

func main() {
	//setPg()
	//resetPg()

	//Compensation()

	//fmt.Println(time.Now().Unix())

	//printVip()

	//db := pg.Connect(Pg)
	//defer db.Close()
	//
	//huatis := make([]HuaTi, 0)
	//err := db.Model(&huatis).Limit(50).Select()
	//if err != nil {
	//	panic(err)
	//}
	//
	//b,err := json.Marshal(huatis)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(len(b))


}

func testNami(){
	callurl := fmt.Sprintf("%v/%v?user=%v&secret=%v", `http://open.sportnanoapi.com`,
		"api/sports/football/match/live",
		"test1218",
		"86c15012a5a70f8907d")

	resp, err := http.Get(callurl)

	if err != nil { //请求失败
		panic(err)
	}

	if resp != nil && resp.StatusCode != 200 {
		fmt.Println("调用第三方错误1", resp)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	live := NMLive{}
	err = json.Unmarshal(body, &live)
	if err != nil {
		fmt.Println("json解析错误", err)
		return
	}

	for _,v := range live{
		fmt.Println(v)
	}

	s := make(chan os.Signal)
	signal.Notify(s,os.Interrupt)
	d := <- s
	fmt.Println(d)
}

type NMLive [][]interface{}
type HuaTi struct {
	Id        int64
	Uid       string
	Username  string
	Avatar    string
	Lv        int
	Title     string
	Txt       string
	Img       string
	Comment   int
	Jing      int
	Createt   int64
	Commentt  int64
	Ding      int
	Dingt     int64
	Jingt     int64
	Vip       int
	Clock     int
	tableName struct{} `pg:"k_huati"`
}

type User struct {
	Id         string
	Username   string
	Pwd        string
	Viplevel   int
	Gold       int64
	Gender     int16
	Iconurl    string
	Nickname   string
	Mobile     string
	Wxid       string
	Tokenuuid  string
	Realname   string
	Idnumber   string
	Email      string
	Exp        int
	Packet     [][]int
	Vip        int
	Vipst      int64
	Yqcode     string //自己的邀请码
	Usedyqcode string //已使用的邀请码
	IsStar     bool   `pg:"is_star"`
	IsShowing  bool   `pg:"is_showing"`
	Uchannel   int
	tableName  struct{} `pg:"k_usrs"`
}

func setPg() {
	db := pg.Connect(Pg) //连接数据库
	defer db.Close()

	users := make([]User, 0)
	err := db.Model(&users).Where("vip > 0").Order("vip").Select()
	if err != nil {
		panic(err)
	}

	file, err := os.Create("./sqldata.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	datastr, err := json.MarshalIndent(users, "  ", " ")
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(datastr))
	if err != nil {
		panic(err)
	}

	nowt := time.Now().Unix()
	for _, v := range users {
		exp := time.Unix(v.Vipst, 0).AddDate(0, v.Vip, 0).Unix()
		if exp <= nowt { //过期
			v.Vip = 0
			v.Vipst = 0
		} else {
			v.Vip *= 31
		}
		err := updateVip(db, v)
		if err != nil {
			panic(err)
		}
	}
}

func updateVip(db *pg.DB, u User) error {
	_, err := db.Model(&u).Set("vip = ?,vipst = ?", u.Vip, u.Vipst).WherePK().Update()
	return err
}

//重置
func resetPg() {
	file, err := os.Open("./sqldata.json")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	sqlData := make([]User, 0)
	err = json.Unmarshal(b, &sqlData)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(Pg) //连接数据库
	defer db.Close()

	for _, v := range sqlData {
		err = updateVip(db, v)
		if err != nil {
			panic(err)
		}
	}
}

//版本补偿
func Compensation() {
	db := pg.Connect(Pg) //连接数据库
	defer db.Close()

	users := make([]User, 0)
	err := db.Model(&users).Column("id", "packet").Select()
	if err != nil {
		panic(err)
	}

	for _, v := range users {
		newpack := AddProp(v.Packet, [][]int{{9, 300}})
		v2 := v
		_, err := db.Model(&v2).Set("packet=?", newpack).WherePK().Update()
		if err != nil {
			panic(err)
		}
	}
}

func AddProp(old [][]int, p [][]int) [][]int {
	for _, v := range p {
		if len(v) != 2 {
			continue
		}
		ex := false
		for i, vv := range old {
			if len(vv) != 2 {
				continue
			}
			if vv[0] == v[0] {
				old[i][1] += v[1]
				ex = true
				break
			}
		}
		if !ex {
			old = append(old, v)
		}
	}
	return old
}

//var NotSame = errors.New("不一样")
//func Check() {
//	db := pg.Connect(Pg) //连接数据库
//	defer db.Close()
//
//	users := make([]User, 0)
//	err := db.Model(&users).Where("vip > 0").Order("vip").Select()
//	if err != nil {
//		panic(err)
//	}
//
//	file, err := os.Open("./sqldata.json")
//	if err != nil {
//		panic(err)
//	}
//
//	b, err := ioutil.ReadAll(file)
//	if err != nil {
//		panic(err)
//	}
//
//	sqlData := make([]User, 0)
//	err = json.Unmarshal(b, &sqlData)
//	if err != nil {
//		panic(err)
//	}
//
//	if len(users) != len(sqlData) {
//		panic(NotSame)
//	}
//
//outs:
//	for _, v := range users {
//		for _, v2 := range sqlData {
//			if v.Id == v2.Id && v.Vip == v2.Vip && v.Vipst == v2.Vipst {
//				continue outs
//			}
//		}
//		panic(NotSame)
//	}
//}

func printVip() {
	db := pg.Connect(Pg) //连接数据库
	defer db.Close()

	users := make([]User, 0)
	err := db.Model(&users).Where("vip > 0").Order("vip desc").Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(len(users))
	for _, v := range users {
		date := time.Unix(v.Vipst, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("days:%v,time:%v,code:%v\n", v.Vip, date, v.Yqcode)
	}
}

type UserMsg struct {
	Title string `json:"title"`
	Txt   string `json:"txt"`
	T     int64  `json:"t"`
}
