package main

import (
	utils "desktop/logintest/util"
	"fmt"
	"testing"
	"time"
)

func TestToken(t *testing.T){
	for i:=1;i<=10;i++{
		time.Sleep(time.Second)
		fmt.Println(utils.GetToken())
	}

}
