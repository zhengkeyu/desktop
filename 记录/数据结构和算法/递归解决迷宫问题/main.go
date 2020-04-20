package main

import "fmt"

func main() {
	gameMap := newGameMap()
	changeMap(gameMap)
	findPath(gameMap,1,1)
	printMap(gameMap)

}
func newGameMap() [][]int {
	arr := make([][]int, 8)
	for k := range arr {
		arr[k] = make([]int, 8)
	}
	return arr
}
func changeMap(arr [][]int){
	//造迷宫围墙
	for k := range arr[0] {
		arr[0][k] = 1
	}
	for k := range arr[len(arr)-1] {
		arr[len(arr)-1][k] = 1
	}
	for _,v := range arr[1:len(arr)-1] {
		for j := range v {
			if j == 0 || j ==len(v)-1{
				v[j] = 1
			}
		}
	}
	//造迷宫内墙
	for i := 1;i<=2;i++{
		arr[3][i] = 1
	}
	for i := len(arr[5])-1;i>=5;i--{
		arr[5][i] = 1
	}
}
func printMap(arr [][]int) {
	for _, v := range arr {
		for _, i := range v {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
func findPath(arr [][]int,h,s int)bool{
	if arr[6][6] == 2{
		fmt.Println("找到出口了")
		return true
	}
	if arr[h][s] == 1{
		return false
	}else {
		arr[h][s] = 2
		if findPath(arr,h+1,s){
			return true
		}else if findPath(arr,h,s+1){
			return true
		}else if findPath(arr,h,s-1){
			return true
		}else if findPath(arr,h-1,s){
			return true
		}else{
			fmt.Println("死迷宫")
			return false
		}
	}
}