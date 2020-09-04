package main

import (
	"encoding/json"
	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	"os"
	"time"
)

//var Pg = &pg.Options{
//	Addr:     "101.200.142.211:5432",
//	User:     "kong",
//	Password: "kong",
//	Database: "postgres",
//}

var Pg = &pg.Options{
	Addr:     "127.0.0.1:5432",
	User:     "postgres",
	Password: "abc123",
	Database: "mydb",
}

func main() {
	//setPg()
	//resetPg()

	//Compensation()

	//fmt.Println(time.Unix(1599127144, 0).Format("2006-01-02 15:04:05"))

	SendMsg()
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

func SendMsg() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	dstr1, err := json.Marshal(UserMsg{
		Title: "标题2",
		Txt:   "txt2",
		T:time.Now().Unix(),
	})

	if err != nil {
		panic(err)
	}

	err = client.RPush("msg.user."+"bbfc6af2-8d0a-4394-b068-8cc3377d0b86", dstr1).Err()
	if err != nil {
		panic(err)
	}

	dstr2, err := json.Marshal(UserMsg{
		Title: "标题2",
		Txt:   "txt2",
		T:time.Now().Unix(),
	})

	if err != nil {
		panic(err)
	}

	err = client.RPush("msg.user."+"bbfc6af2-8d0a-4394-b068-8cc3377d0b86", dstr2).Err()
	if err != nil {
		panic(err)
	}
}

type UserMsg struct {
	Title string `json:"title"`
	Txt   string `json:"txt"`
	T     int64  `json:"t"`
}
