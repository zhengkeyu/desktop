package main

import (
	"fmt"
	"strconv"
)

//栈结构
type Stack struct {
	NumIndex    int
	NumArr      []int
	SymbolIndex int
	SymbolArr   []string
}

//栈是先进后出，进栈出栈操作的都是栈尾
func main() {
	//用栈的特点实现的计算器
	//calculate := "1+3*6/2+3+9+2*6-3*2"
	calculate := "21+2*2+14/2-2"
	fmt.Println(Run(calculate))
}

func Run(str string) int {
	stack := Stack{
		NumIndex:    -1,
		NumArr:      make([]int, len(str)/2+1),
		SymbolIndex: -1,
		SymbolArr:   make([]string, len(str)/2),
	}
	b := false
	for k := range str {
		by := fmt.Sprintf("%c", str[k])
		switch by {
		case "+", "-", "*", "/":
			if stack.SymbolIndex == -1 {
				stack.SymbolArr[0] = by
				stack.SymbolIndex++
			} else {
				switch by {
				case "+":
					if stack.SymbolArr[stack.SymbolIndex] == "*" || stack.SymbolArr[stack.SymbolIndex] == "/" {
						ChangeStack(&stack, by) //如果之前的运算符比当前的优先就调用此函数
					} else {
						stack.SymbolArr[stack.SymbolIndex+1] = by
						stack.SymbolIndex++
					}
				case "-":
					if stack.SymbolArr[stack.SymbolIndex] == "*" || stack.SymbolArr[stack.SymbolIndex] == "/" {
						ChangeStack(&stack, by)
					} else {
						stack.SymbolArr[stack.SymbolIndex+1] = by
						stack.SymbolIndex++
					}
				case "*":
					if stack.SymbolArr[stack.SymbolIndex] == "*" || stack.SymbolArr[stack.SymbolIndex] == "/" {
						ChangeStack(&stack, by)
					} else {
						stack.SymbolArr[stack.SymbolIndex+1] = by
						stack.SymbolIndex++
					}
				case "/":
					if stack.SymbolArr[stack.SymbolIndex] == "*" || stack.SymbolArr[stack.SymbolIndex] == "/" {
						ChangeStack(&stack, by)
					} else {
						stack.SymbolArr[stack.SymbolIndex+1] = by
						stack.SymbolIndex++
					}
				}
			}
			b = false
		default:
			if b == true { //用于处理用于计算的数是1位以上的情况
				newnum, _ := strconv.Atoi(strconv.Itoa(stack.NumArr[stack.NumIndex]) + by)
				stack.NumArr[stack.NumIndex] = newnum
			} else {
				num, _ := strconv.Atoi(by)
				if stack.NumIndex == -1 {
					stack.NumArr[0] = num
					stack.NumIndex++
					b = true
				} else {
					stack.NumArr[stack.SymbolIndex+1] = num
					stack.NumIndex++
					b = true
				}
			}
		}
	}
	return stack.GetFinalResult()
}

//出栈压栈计算当前运算符
func ChangeStack(stack *Stack, str string) {
	switch stack.SymbolArr[stack.SymbolIndex] {
	case "*":
		stack.NumArr[stack.NumIndex-1] = stack.NumArr[stack.NumIndex-1] * stack.NumArr[stack.NumIndex]
		stack.NumIndex--
	case "/":
		stack.NumArr[stack.NumIndex-1] = stack.NumArr[stack.NumIndex-1] / stack.NumArr[stack.NumIndex]
		stack.NumIndex--
	}
	stack.SymbolArr[stack.SymbolIndex] = str
}

//遍历运算符栈，计算结果
func (this Stack) GetFinalResult() int {
	for i := this.SymbolIndex; i != -1; i-- {
		switch this.SymbolArr[i] {
		case "*":
			this.NumArr[this.NumIndex-1] = this.NumArr[this.NumIndex-1] * this.NumArr[this.NumIndex]
			this.NumIndex--
		case "/":
			this.NumArr[this.NumIndex-1] = this.NumArr[this.NumIndex-1] * this.NumArr[this.NumIndex]
			this.NumIndex--
		case "+":
			this.NumArr[this.NumIndex-1] = this.NumArr[this.NumIndex-1] + this.NumArr[this.NumIndex]
			this.NumIndex--
		case "-":
			this.NumArr[this.NumIndex-1] = this.NumArr[this.NumIndex-1] - this.NumArr[this.NumIndex]
			this.NumIndex--
		}
	}
	return this.NumArr[this.NumIndex]
}
