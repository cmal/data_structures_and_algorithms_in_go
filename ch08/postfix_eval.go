package main

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/golang-collections/collections/stack"
)

func main() {
	expn := "6 14 2 3 + 8 * + 3 + *"
	value := postfixEval(expn)
	fmt.Println("Given Postfix Expn: ", expn)
	fmt.Println("Result after Eval: ", value)
}

func postfixEval(expn string) int {
	stk := stack.New()
	str := strings.Split(expn, " ")

	for _, item := range str {
		switch item {
		case "+":
			num1 := stk.Pop().(int)
			num2 := stk.Pop().(int)
			stk.Push(num1 + num2)
		case "-":
			num1 := stk.Pop().(int)
			num2 := stk.Pop().(int)
			stk.Push(num1 - num2)
		case "*":
			num1 := stk.Pop().(int)
			num2 := stk.Pop().(int)
			stk.Push(num1 * num2)
		case "/":
			num1 := stk.Pop().(int)
			num2 := stk.Pop().(int)
			stk.Push(num1 / num2)
		default:
			num, _ := strconv.Atoi(item)
			stk.Push(num)
		}
	}
	return stk.Pop().(int)
}
