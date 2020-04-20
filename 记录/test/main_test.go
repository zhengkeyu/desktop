package test

import (
	"math/rand"
	"sort"
	"testing"
	"time"
	"fmt"
)

func TestFun(t *testing.T){
	var arr = make([]int,1000000)
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(arr);i++{
		arr[i] = rand.Intn(10000000)
	}
	sort.Ints(arr)
	fmt.Println(arr)
}
