package main

import (
	"github.com/go-pg/pg/v9"
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

type User struct {
	Id        int      `pg:"id"`
	Uid       int      `pg:"uid"`
	PassWord  string   `pg:"password"`
	Name      string   `pg:"name"`
	tableName struct{} `pg:"user"`
}

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

	////插入数据
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
	//err = db.Model(&mytable{}).Where(`id>0`).Offset(2).Select(&d)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(d)

	////匹配字符
	//r := make([]mytable, 0)
	//_, err = db.Query(&r, `select * from mytables where learn[1] like '__go%'`)
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
	//_, err = db.Model(&d).OnConflict("(id) DO UPDATE").
	//	Set("uid=EXCLUDED.uid,password=EXCLUDED.password,name=EXCLUDED.name").Insert()
	//if err != nil {
	//	panic(err)
	//}

}
