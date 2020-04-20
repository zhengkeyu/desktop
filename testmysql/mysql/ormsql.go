package ormsql

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	SQL_PASSWORD  = "abc123"
	DATABASE_NAME = "zheng"
)

type Table struct {
	Id    int
	Uid   int64
	Coins int
	Date  string
}

func (u *Table) TableName() string {
	return "TableInfo"
}

var Ormer orm.Ormer

func RunMysql(dbName,Password string) {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}
	err = orm.RegisterDataBase("default", "mysql", "root:"+Password+"@/"+dbName+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(Table))
	_ = orm.RunSyncdb("default", false, true)
	Ormer = orm.NewOrm()
	fmt.Println("mysql 初始化成功")
}

//是否有这个表
func IsHaveTableByName(tableName string) bool {
	r := Ormer.Raw("show tables")
	ps := &[]orm.Params{}
	_, err := r.Values(ps)
	if err != nil {
		panic(err)
	}
	key := "Tables_in_" + DATABASE_NAME
	for _, v := range *ps {
		if v[key] == tableName {
			return true
		}
	}
	return false
}
