package main

func main() {
	//rand.Seed(time.Now.UnixNano()) //设置种子为当前时间的纳秒
	//rand.Float64()float64  //[0,10)的随机浮点数
	//rand.Intn(10)int      //[0,10)根据输入的值来生成随机整数

	//不用设定种子就可返回随机整数,相对上面的更随机，安全
	//b,_ := rand.Int(rand.Reader,big.NewInt(100000))
	//fmt.Println(b.Int64())

	//返回[0,100)的随机数数组
	//rand.Seed(time.Now().UnixNano())
	//arr := rand.Perm(100)
	//fmt.Println(arr)
}