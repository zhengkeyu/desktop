package main

import (
	"github.com/go-pg/pg/v9"
	"time"
)

type mytable struct {
	Id        int      `sql:"id,type:serial"` //serial指定自增
	Name      string   `sql:"name,type:varchar(100)"`
	Age       int      `sql:"age,type:int"`
	Learn     []string `sql:"learn,type:text[]"`
	tableName struct{} `pg:"mytables"`
}

type project struct {
	Id        int      `sql:"id,type:int"`
	Uid       int      `sql:"uid,type:int"`
	PName     string   `sql:"pname,type:varchar(100)"`
	tableName struct{} `pg:"myproject"`
}

//type project2 struct {
//	Id        int      `sql:"id,type:int"`
//	Uid       int      `sql:"uid,type:int"`
//	PName     string   `sql:"pname,type:varchar(100)"`
//	tableName struct{} `pg:"mys"`
//}

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "abc123",
		Database: "mydb",
	})
	defer db.Close()

	////建表
	//err := db.CreateTable(&mytable{}, &orm.CreateTableOptions{
	//	IfNotExists:   true,  //设定彪不存在再创建
	//	FKConstraints: false, //外键
	//})
	//if err != nil {
	//	panic(err)
	//}

	////删表
	//_,err = db.Exec(`drop table if exists mytables`)
	//if err != nil{
	//	panic(err)
	//}

	////建表
	//err = db.CreateTable(&project{}, &orm.CreateTableOptions{
	//	IfNotExists:   true,  //设定彪不存在再创建
	//	FKConstraints: false, //外键
	//})
	//if err != nil {
	//	panic(err)
	//}

	////插入数据 //可放数组，批量插入
	//err = db.Insert(&mytable{
	//	Name:  "guo",
	//	Age:   23,
	//	Learn: []string{"python", "lua", "c#"},
	//})

	////查询
	//var data = []mytable{}
	//_, err = db.Query(&data, `select * from mytables`)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(data)

	////Exec()方法,没有返回，不适用查询
	//alter table mytables add Id int //添加列
	//ALTER  TABLE mytables ALTER COLUMN id SET NOT NULL 、、设置列不能为空
	//_, err =db.Exec(`ALTER TABLE mytables ADD CONSTRAINT id PRIMARY KEY(id)`) //设置列为主键
	//if err != nil {
	//	panic(err)
	//}

	////更新
	//err = db.Update(&mytable{
	//	Id:    1,
	//	Name:  "wang",
	//	Age: 100,
	//	Learn: []string{"ac", "teach"},
	//})
	//if err != nil {
	//	panic(err)
	//}

	////指定更新那一列
	//_, err = db.Model(&mytable{
	//	Id: 1,
	//	Age: 999,
	//}).Column("age").WherePK().Update()
	//if err != nil {
	//	panic(err)
	//}

	////条件查询
	//d := []mytable{}
	//err = db.Model(&mytable{}).Where(`age>22`).Order("age desc").Select(&d)
	//if err != nil {
	//	panic(err)
	//}
	////age>22:列age大于22，age desc:按列age降序排列
	//fmt.Println(d)

	////联表查询
	//s := `select t1.name,t1.age,t2.pname from public.mytables as t1 join public.myproject as t2 on t1.id=t2.id`
	//var r []struct {
	//	Name  string `sql:"name,type:varchar(100)"`
	//	Age   int    `sql:"age,type:int"`
	//	PName string `sql:"pname,type:varchar(100)"`
	//}
	//_, err = db.Query(&r, s)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(r)

	////占位符
	//d := new(mytable)
	//_, err = db.Query(d, `select * from mytables where age>? and id>?`, 22, 2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(d)

	////offset()和limit()
	////offset():偏移数，limit()：只取几条
	//d := make([]mytable,0)
	//err = db.Model(&mytable{}).Where(`id>0`).Offset(2).Select(&d) //该查询允许结构体和表字段不一致
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(d)

	////匹配字符
	//r := make([]mytable, 0)
	//_, err = db.Query(&r, `select * from mytables where learn[1] like '__go%'`) //该查询不允许结构体和表字段不一致
	//if err != nil {
	//	panic(err)
	//}
	////%匹配多个任意字符，_匹配一个任意字符
	//fmt.Println(r)

	////with的使用，执行后需再定义才有效，不是永久定义
	////with相当于封装了一段sql语句，用于后续的调用
	//var d []project
	//_,err = db.Query(&d,`with mys as (select uid,pname from myproject)select * from mys`)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(d)

	////returning的使用
	//o := new(mytable)
	//_, err = db.Query(o, `with s1 as (delete from mytables where name='liu' returning *)insert into m2 (select * from s1)`)
	////将mytables表中删除的数据插入到m2中
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(o)

	////能插入执行插入，不行执行更新Set(...)
	////约束的(id)里的id要唯一，括号内可以设置多个，需要都不重复才是插入否则是更新
	//d := User{
	//	Id: 3,
	//	Uid: 4,
	//	PassWord: "newpassword",
	//	Name: "newgogo",
	//}
	//_, err = db.Model(&d).OnConflict("(id) DO UPDATE"). //do nothing为不做操作
	//	Set("uid=EXCLUDED.uid,password=EXCLUDED.password,name=EXCLUDED.name").Insert()
	//if err != nil {
	//	panic(err)
	//}

	//d := 0
	//
	//_, err := db.Query(&d, `select max(startt) from prediction`)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(d)

	//_, err := db.Model(&o).OnConflict("(userid) do update").
	//	Set("totaltime = ?totaltime+totaltime").Insert()
	//if err != nil{
	//	panic(err)
	//}
	t := T1{
		Id:   2147483647,
		Name: "",
	}
	_,err := db.Model(&t).Set("name = ?",t.Name).WherePK().Update()
    if err != nil{
    	panic(err)
	}
}

type T1 struct {
	Id   int
	Name string
	tableName struct{} `pg:"t1"`
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
	Packet     string
	Vip        int
	Vipst      int64
	Yqcode     string
	Usedyqcode string
	Usource    int
	Uchannel   int
	Mark       int
	Status     int
	Lastlogin  time.Time
	Createt    time.Time `pg:"create_at"`
	IsStar     bool      `pg:"is_star"`
	IsShowing  bool      `pg:"is_showing"`
	tableName  struct{}  `pg:"k_usrs"`
}

type InviteTask struct {
	Id        int
	Name      string
	Head      string
	Gifts     []string
	Uid       string
	tableName struct{} `pg:"invitetask"`
}

type OnLineTime struct {
	Id        int
	UserId    string   `pg:"userid"`
	TotalTime int      `pg:"totaltime"`
	tableName struct{} `pg:"onlinetime"`
}
type Test struct {
	Id        int
	UserName  string `pg:"username"`
	Bpwd      string
	Pwd       string
	tableName struct{} `pg:"test"`
}
type Prediction struct {
	Id        int      `pg:"id"`
	NaMiId    int      `pg:"namiid"`
	Name      string   `pg:"name"`
	Startt    int64    `pg:"startt"`
	Companyid int      `pg:"companyid"`
	RqResult  int      `pg:"rqresult"`
	SfpResult int      `pg:"sfpresult"`
	DxResult  int      `pg:"dxresult"`
	tableName struct{} `pg:"prediction"`
}

type LiaoTianShiCount struct {
	Id        int
	NamiId    int64 `pg:"namiid"`
	Startt    int64
	Name      []string
	UserId    string   `pg:"userid"`
	tableName struct{} `pg:"liaotianshicount"`
}

func RunRemote() {
	db := pg.Connect(&pg.Options{
		Addr:     "123.207.85.242:5433",
		User:     "postgres",
		Password: "2932615qian",
		Database: "postgres",
	})
	defer db.Close()
}

type NaMi struct {
	Namiid     int   `pg:"namiid"`
	Startt     int64 `pg:"startt"`
	Startballt int64 `pg:"startballt"`
}
