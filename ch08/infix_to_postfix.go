package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
)

func main() {
	fmt.Println(InfixToPostfix("1+2*3"))
	fmt.Println(InfixToPostfix("(1+2)*3"))
	fmt.Println(InfixToPostfix("(1+2)*(3+4)"))
	fmt.Println(InfixToPostfix("(1*2)+(3*4)"))
	fmt.Println(InfixToPostfix("1^2*3+4"))
	fmt.Println(InfixToPostfix("1+2*3^4"))
	fmt.Println(InfixToPostfix("1+((3))*5/(6-4)"))
	fmt.Println(InfixToPostfix("10+((3))*5/(163-4)"))

	fmt.Println("---- Infix to Prefix ----")
	fmt.Println(InfixToPrefix("1+2"))
	fmt.Println(InfixToPrefix("1+(2*3)"))
	fmt.Println(InfixToPrefix("(1+2)*3"))
	fmt.Println(InfixToPrefix("10+((3))*5/(163-4)"))
}

func InfixToPostfix(expn string) string {
	fmt.Printf("%v    =>    ", expn)
	stk := stack.New()
	resQ := queue.New()
	output := ""

	for _, ch := range expn {
		if ch >= '0' && ch <= '9' {
			// fmt.Println("Enqueue " + string(ch))
			resQ.Enqueue(ch)
		} else {
			switch ch {
			case '+', '-', '*', '/', '%', '^':
				// fmt.Println("stack len: ", stk.Len())
				for stk.Len() != 0 {
					peek := stk.Peek().(rune)
					// fmt.Println("Compare ", string(ch), " to ", string(peek))
					if precedence(ch) > precedence(peek) {
						// fmt.Println(">")
						break;
					// } else {
						// fmt.Println("<=")
					}
					// fmt.Println("Pop")
					head := stk.Pop().(rune)
					if head == '(' {
						break
					}
					// fmt.Println("Enqueue " + string(head))
					resQ.Enqueue(head)
				}
				// fmt.Println("Push " + string(ch))
				stk.Push(ch)
				resQ.Enqueue(' ')
			case '(':
				// fmt.Println("Push " + string(ch))
				stk.Push(ch)
			case ')':
				for stk.Len() != 0 {
					// fmt.Println("Pop")
					head := stk.Pop().(rune)
					if head == '(' {
						break
					}
					// fmt.Println("Enqueue " + string(head))
					resQ.Enqueue(head)
				}
			}
		}
	}
	for stk.Len() != 0 {
		// fmt.Println("Pop ")
		head := stk.Pop().(rune)
		// fmt.Println("Enqueue " + string(head))
		resQ.Enqueue(head)
	}

	for resQ.Len() != 0 {
		head := resQ.Dequeue().(rune)
		output += string(head)
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

func reverseString(expn string) string {
	length := len(expn)
	reversed := make([]rune, length)
	for index, char := range expn {
		reversed[length - 1 - index] = char
	}
	return string(reversed)
}

func swapParenthese(expn string) string {
	runes := []rune(expn)
	for index, char := range expn {
		switch char {
		case '(':
			runes[index] = ')'
		case ')':
			runes[index] = '('
		}
	}
	return string(runes)
}

func InfixToPrefix(expn string) string {
	fmt.Printf("%v    =>    ", expn)
	return reverseString(InfixToPostfix(swapParenthese(reverseString(expn))))
}
