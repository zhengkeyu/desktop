package main

import (
	"Desktop/记录/十三水/card"
	"fmt"
	"strconv"
)

func main() {
	//var d = &desk.Desk{
	//	PlayerCards:[]*card.Card{new(card.Card),new(card.Card)},
	//}
	//d.SendCards()
	//fmt.Println(d.PlayerCards[0],d.PlayerCards[1])
	var card = &card.Card{}
	strArr := []string{"黑3","梅3","梅5","方1","方5","方7","梅7","红8","黑8","方10","黑10","方12","红12"}
	card.Value = GiveMeCards(strArr)
	fmt.Println("MyCards:",card.Value)
	_,Type := card.IsSpecial()
	fmt.Println(Type)
}
func GiveMeCards(strArr []string)[]int{
	cards := make([]int,0)
	for _,v := range strArr{
		num,_ := strconv.Atoi(v[3:])
		if num == 1{
			num = 14
		}
		switch v[:3] {
		case "黑":
			cards = append(cards,0x30+num)
		case "红":
			cards = append(cards,0x20+num)
		case "梅":
			cards = append(cards,0x10+num)
		case "方":
			cards = append(cards,0x00+num)
		}
	}
	card.CardSort(cards)
	return cards
}