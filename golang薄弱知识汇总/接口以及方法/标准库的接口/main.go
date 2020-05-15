package main

import "fmt"

type Player struct {
	Name  string
	Uid   int64
	Card  []int
	Other map[int]string
}

func (p *Player) String() string { //实现这个函数，打印*Player时会按这个函数的实现来打印
	return fmt.Sprintf("Name: %v\nUid: %v\nCard: %v\nOther: %v\n",
		p.Name, p.Uid, p.Card, p.Other)
}
func main() {
	p := &Player{
		Name:  "zheng",
		Uid:   1024,
		Card:  []int{1, 2, 3, 4},
		Other: map[int]string{1: "三国演义", 2: "水浒传"},
	}
	fmt.Println(p)
}
