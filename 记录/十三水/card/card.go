package card

import (
	"math/rand"
	"time"
)

type Card struct {
	Value   []int
	Type    int
	Default [][]int
}

//初始化卡组
func InitCards() []int {
	cards := make([]int, 0)
	for i := 2; i <= 14; i++ {
		cards = append(cards, i+0x00, i+0x10, i+0x20, i+0x30)
	}
	return cards
}

//打乱卡组
func DisturbCards(cards []int) []int {
	newcards := make([]int, len(cards))
	rand.Seed(time.Now().UnixNano())
	randArr := rand.Perm(len(cards))
	for k, v := range randArr {
		newcards[k] = cards[v]
	}
	return newcards
}

//获取牌的数值
func GetCradValue(card int) int {
	return card % 16
}

//获取牌的花色
func GetCardColr(card int) int {
	return card / 16
}

//根据数值和花色排序
func CardSort(cards []int) {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < len(cards)-1-i; j++ {
			if GetCradValue(cards[j]) > GetCradValue(cards[j+1]) {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			} else if GetCradValue(cards[j]) == GetCradValue(cards[j+1]) {
				if GetCardColr(cards[j]) > GetCardColr(cards[j+1]) {
					cards[j], cards[j+1] = cards[j+1], cards[j]
				}
			}
		}
	}
}

//获取默认三墩的牌型
func (this *Card) GetDefault() {
	//cards := append([]int{}, this.Value...)
	//this.Default[2] = GetSimpleCards()
}
//获取牌型
//func GetSimpleCards(cards *[]int)[]int{
//	//switch  {
//	//case :
//	//
//	//}
//}
//判断铁支
func IsTieZhi(cards []int) []int {
	card := make([]int, 0)
	for i := 0; i < len(cards)-1; i++ {
		if GetCradValue(card[i]) == GetCradValue(card[i+1]) {
			if len(card) == 0 || len(card) == 4 {
				card = append(card, card[i], card[i+1])
			} else {
				card = append(card, card[i+1])
			}
		} else {
			if len(card) == 8 {
				return card
			}
			if len(card) < 4 {
				card = make([]int, 0)
			} else {
				card = card[:4]
			}
		}
	}
	return card
}

//是不是特殊牌
func (this *Card) IsSpecial() (int, string) {
	if v1, str1 := this.IsLongType(); v1 != -1 {
		//this.Type = v1
		return v1, str1
	}
	if this.IsSanTongHuaShun() {
		return SAN_TONG_HUA_SHUN, "三同花顺"
	}
	if this.IsSanFenTianXia() {
		return SAN_FEN_TIAN_XIA, "三分天下"
	}
	if this.IsSiTaoSanTiao() {
		return SI_TAO_SAN_TIAO, "四套三条"
	}
	if this.IsLiuDuiBan() {
		return LIU_DUI_BAN, "六对半"
	}
	if this.IsSanShunZi() {
		return SAN_SHUN_ZI, "三顺子"
	}
	if this.IsSanTongHua() {
		return SAN_TONG_HUA, "三同花"
	}
	return -1, ""
}

//是不是青龙或一条龙
func (this *Card) IsLongType() (int, string) {
	b1 := true
	for i := 0; i < 12; i++ {
		if GetCradValue(this.Value[i]) == GetCradValue(this.Value[i+1]) {
			return -1, ""
		}
	}
	for i := 0; i < 12; i++ {
		if GetCardColr(this.Value[i]) != GetCardColr(this.Value[i+1]) {
			b1 = false
		}
	}
	if b1 {
		return QING_LONG, "青龙"
	} else {
		return YI_TIAO_LONG, "一条龙"
	}
}

//获取手牌里指定的顺子
func GetShunZiArr(card *[]int, c int) []int {
	arr := make([]int, 0)
	count := 044
	for i := 1; i < len(*card); i++ {
		if GetCradValue((*card)[i])-GetCradValue((*card)[count]) == 1 &&
			GetCardColr((*card)[i]) == GetCardColr((*card)[count]) {
			if len(arr) == 0 {
				arr = append(arr, (*card)[count], (*card)[i])
			} else {
				arr = append(arr, (*card)[i])
				if len(arr) == c {
					res := make([]int, 0)
				hear:
					for _, v := range *card {
						for _, l := range arr {
							if v == l {
								continue hear
							}
						}
						res = append(res, v)
					}
					*card = res
					return arr
				}
			}
			count = i
		} else if GetCradValue((*card)[i])-GetCradValue((*card)[count]) == 0 {
			continue
		} else {
			arr = make([]int, 0)
			count = i
		}
	}
	return arr
}

//获取手牌里指定的顺子
func GetShunZiArrTwo(card *[]int, c int) []int {
	arr := make([]int, 0)
	count := 0
	for i := 1; i < len(*card); i++ {
		if GetCradValue((*card)[i])-GetCradValue((*card)[count]) == 1 {
			if len(arr) == 0 {
				arr = append(arr, (*card)[count], (*card)[i])
			} else {
				arr = append(arr, (*card)[i])
				if len(arr) == c {
					res := make([]int, 0)
				hear:
					for _, v := range *card {
						for _, l := range arr {
							if v == l {
								continue hear
							}
						}
						res = append(res, v)
					}
					*card = res
					return arr
				}
			}
			count = i
		} else if GetCradValue((*card)[i])-GetCradValue((*card)[count]) == 0 {
			continue
		} else {
			arr = make([]int, 0)
			count = i
		}
	}
	return arr
}

//是不是三同花顺
func (this *Card) IsSanTongHuaShun() bool {
	cards := append([]int{}, this.Value...)
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	arr3 := make([]int, 0)
	arr1 = GetShunZiArr(&cards, 5)
	arr2 = GetShunZiArr(&cards, 5)
	arr3 = GetShunZiArr(&cards, 3)
	if len(arr1) == 5 &&
		len(arr2) == 5 &&
		len(arr3) == 3 {
		if IsSameColor(arr1) &&
			IsSameColor(arr2) &&
			IsSameColor(arr3) {
			return true
		} else {
			return false
		}
	}
	//
	cards = append([]int{}, this.Value...)
	for k, v := range cards {
		if GetCradValue(v) == 14 {
			cards[k] -= 13
		}
	}
	CardSort(cards)
	arr1 = GetShunZiArr(&cards, 5)
	arr2 = GetShunZiArr(&cards, 5)
	arr3 = GetShunZiArr(&cards, 3)
	if len(arr1) == 5 &&
		len(arr2) == 5 &&
		len(arr3) == 3 {
		if IsSameColor(arr1) &&
			IsSameColor(arr2) &&
			IsSameColor(arr3) {
			return true
		} else {
			return false
		}
	}
	return false
}

//判断牌是否相同花色
func IsSameColor(cards []int) bool {
	for i := 1; i < len(cards); i++ {
		if GetCardColr(cards[i]) != GetCardColr(cards[0]) {
			return false
		}
	}
	return true
}

//是不是三分天下
func (this *Card) IsSanFenTianXia() bool {
	accumulate := 0
	count := 0
	for i := 0; i < len(this.Value)-1; i++ {
		if GetCradValue(this.Value[i]) == GetCradValue(this.Value[i+1]) {
			count++
			if count == 3 {
				accumulate++
				count = 0
			}
		} else {
			if count != 0 {
				return false
			}
		}
	}
	if accumulate == 3 {
		return true
	}
	return false
}

//是不是四套三条
func (this *Card) IsSiTaoSanTiao() bool {
	accumulate := 0
	count := 0
	for i := 0; i < len(this.Value)-1; i++ {
		if GetCradValue(this.Value[i]) == GetCradValue(this.Value[i+1]) {
			count++
			if count == 2 {
				accumulate++
				count = 0
			}
		} else if count != 0 {
			return false
		}
	}
	if accumulate == 4 {
		return true
	}
	return false
}

//是不是六对半
func (this *Card) IsLiuDuiBan() bool {
	count := 0
	for i := 0; i < len(this.Value)-2; i++ {
		if GetCradValue(this.Value[i]) == GetCradValue(this.Value[i+1]) &&
			GetCradValue(this.Value[i+1]) != GetCradValue(this.Value[i+2]) {
			count++
		}
	}
	if count == 6 {
		return true
	}
	return false
}

//是不是三顺子
func (this *Card) IsSanShunZi() bool {
	cards := append([]int{}, this.Value...)
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	arr3 := make([]int, 0)
	arr1 = GetShunZiArrTwo(&cards, 5)
	arr2 = GetShunZiArrTwo(&cards, 5)
	arr3 = GetShunZiArrTwo(&cards, 3)
	if len(arr1) == 5 &&
		len(arr2) == 5 &&
		len(arr3) == 3 {
		return true
	}
	//
	cards = append([]int{}, this.Value...)
	for k, v := range cards {
		if GetCradValue(v) == 14 {
			cards[k] -= 13
		}
	}
	CardSort(cards)
	arr1 = GetShunZiArrTwo(&cards, 5)
	arr2 = GetShunZiArrTwo(&cards, 5)
	arr3 = GetShunZiArrTwo(&cards, 3)
	if len(arr1) == 5 &&
		len(arr2) == 5 &&
		len(arr3) == 3 {
		return true
	}
	return false
}

//是不是三同花
func (this *Card) IsSanTongHua() bool {
	color := make([]int, 4)
	for i := 0; i < len(this.Value); i++ {
		switch GetCardColr(this.Value[i]) {
		case 0:
			color[0]++
		case 1:
			color[1]++
		case 2:
			color[2]++
		case 3:
			color[3]++
		}
	}
	for _, v := range color {
		if !(v == 3 || v == 0 || v == 13 || v == 5 || v == 10 || v == 8) {
			return false
		}
	}
	return true
}
