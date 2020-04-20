package a

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type UserLogin struct {
	id     int
	uid    int
	data   string
	handle Login
}
type Login struct {
	Coins   int
	Token   int
	Uid     int
	Version string
}

func main() {
	db, err := sql.Open("mysql", "root:abc123@/zheng")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("SELECT * FROM userlogin WHERE id=?")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(100)
	if err != nil {
		panic(err)
	}
	rows.Next()
	u := UserLogin{}
	str := ""
	err = rows.Scan(&u.id, &u.uid, &u.data, &str)
	json.Unmarshal([]byte(str), &u.handle)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
