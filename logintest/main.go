package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	r := md5.Sum([]byte("zky"))
	str := fmt.Sprintf("%x",r)
	fmt.Println(len(str))
}
