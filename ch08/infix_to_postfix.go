package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {

}

func InfixToPostfix(expn string) string {
	fmt.Println(expn)
	stk := new(Stack)
	resStk := new(Stack)
	output := ""

	for _, ch := range expn {
		if ch >= '0' && ch <= '9' {
			stk.Push(ch)
		} else {
			switch ch {
			case '+', '-', '*', '/', '%', '^':
				for !stk.IsEmpty() {
					head := stk.Pop()
					if precedence(ch) >= precedence(head) && head != '(' {
						resStk.Push(head)
					} else {
						stk.Push(head)
					}
				}
				stk.push(ch)
			case '(':
				stk.Push(ch)
			case ')':
				for !stk.IsEmpty() {
					head := stk.Pop()
					if head == '(' {
						break;
					} else {
						resStk.Push(head)
					}
				}
			}
		}
		for stk.IsEmpty() == false {
			head := stk.Pop()
			resStk.Push(head)
		}
	}

	s := make([]rune, resStk.Len())
	i := 0
	for resStk.IsEmpty() {
		s[i++] = resStk.Pop().(rune)
	}
	for _, item = range s {
		output += item + " "
	}
	return output
}

func precedence(x rune) int {
	switch x {
	case '(':
		return 0
	case '+', '-':
		return 1
	case '*', '/', '%':
		return 2
	case '^':
		return 3
	default:
		return 4
	}
}
