package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	fmt.Println(IsPalindrome("ABA"))
	fmt.Println(IsPalindrome("ABBA"))
	fmt.Println(IsPalindrome(""))
	fmt.Println(IsPalindrome("A"))
	fmt.Println(IsPalindrome("AA"))
	fmt.Println(IsPalindrome("ABC"))
	fmt.Println(IsPalindrome("AAABBBCCCBBBAAA"))
}

func IsPalindrome(s string) bool {
	length := len(s)
	halflen := length/2
	even := length % 2 == 0
	stk := stack.New()

	for index, char := range s {
		if !even && index == halflen {
			continue
		} else if index >= halflen {
			top := stk.Pop().(rune)
			if top != char {
				return false
			}
		} else {
			stk.Push(char)
		}
	}
	return true
}
