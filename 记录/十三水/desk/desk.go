package desk

import (
	"Desktop/记录/十三水/card"
)

const PLAYERCOUNT = 2

type Desk struct {
	DeskCards   []int
	PlayerCards []*card.Card
}

//发牌
func (this *Desk) SendCards() {
	this.DeskCards = card.DisturbCards(card.InitCards())
	for i := 0; i < PLAYERCOUNT; i++ {
		this.PlayerCards[i].Value = append(this.PlayerCards[i].Value, this.DeskCards[:13]...)//发牌
		card.CardSort(this.PlayerCards[i].Value) //排序
		this.PlayerCards[i].IsSpecial() //检查是否有特殊牌
		this.DeskCards = this.DeskCards[13:]
	}
}
