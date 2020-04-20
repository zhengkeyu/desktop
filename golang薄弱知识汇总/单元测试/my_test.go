package cal

import (
	"strconv"
	"testing"
)

func TestThree(t *testing.T){
	_,err := strconv.Atoi("111")
	if err != nil{
		t.Fatalf("strconv.Atoi()执行错误%v",err)
	}
	t.Logf("success")
}
