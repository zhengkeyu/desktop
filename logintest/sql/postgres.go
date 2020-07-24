package sql

import (
	utils "desktop/logintest/util"
	"errors"
	"github.com/go-pg/pg/v9"
	"time"
)

var Db *pg.DB

func init() {
	Db = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "abc123",
		Database: "mydb",
	})

	var n int
	_, err := Db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
}

type Register struct {
	UserName string `json:"username" binding:"min=4,max=12"`
	PassWord string `json:"password" binding:"min=6,max=12"`
}

type UserInfo struct {
	Id           int      `pg:"id"`
	UserName     string   `pg:"username"`
	PassWord     string   `pg:"password"`
	Perm         int      `pg:"perm"`
	Token2       string   `pg:"token2"`
	Token6       string   `pg:"token6"`
	RefreshTime2 int64    `pg:"refreshtime2"`
	RefreshTime6 int64    `pg:"refreshtime6"`
	tableName    struct{} `pg:"houtai_login"`
}

type ReqToken struct {
	Token string `json:"token"`
}

func CheckRegister(name string) (bool, error) {
	sqlData := UserInfo{}
	_, err := Db.Query(&sqlData, `select * from houtai_login where username = ?`, name)
	if err != nil {
		return false, err
	}

	if sqlData.Id == 0 {
		return true, nil
	}
	return false, nil //用户名被注册了
}

func InsertUser(u Register) error {
	err := Db.Insert(&UserInfo{
		UserName: u.UserName,
		PassWord: u.PassWord,
		Perm:     1,
	})
	if err != nil {
		return err
	}
	return nil
}

func CheckLogin(d Register) (bool, error) {
	sqlData := UserInfo{}
	_, err := Db.Query(&sqlData, `select * from houtai_login where username = ?`, d.UserName)
	if err != nil {
		return false, err
	}

	if d.PassWord != sqlData.PassWord {
		return false, errors.New("密码错误或用户不存在")
	}

	return true, nil
}

func NewToken(name string) (string, string, error) {
	m := UserInfo{
		Token2: utils.GetToken(),
	}
	time.Sleep(time.Second)
	m.Token6 = utils.GetToken()
	_, err := Db.Model(&m).Column("token2", "token6").Where("username = ?", name).Update()
	if err != nil {
		return "", "", err
	}

	return m.Token2, m.Token6, nil
}

func CheckRefreshTimeByToken6(token string, ) (bool, error) {
	u := UserInfo{}
	err := Db.Model(&u).Where("token6 = ?", token).Select()
	if err != nil {
		return false, err
	}
	if u.Id == 0 || u.RefreshTime2+3600*6 >= time.Now().Unix() { //过期
		return false, nil
	}

	return true, nil
}

func RefreshToken2(token6 string) error {
	u := UserInfo{
		Token2:       utils.GetToken(),
		RefreshTime2: time.Now().Unix(),
	}
	_, err := Db.Model(&u).Where("token6 = ?", token6).Column("token2", "refreshtime2").Update()
	if err != nil {
		return err
	}

	return nil
}
