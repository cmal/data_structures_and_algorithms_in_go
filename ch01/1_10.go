package main

import "fmt"

func isEven(value int) {
	switch {
	case value % 2 == 0:
		fmt.Println("is event")
	default:
		fmt.Println("is odd")
	}
}


func main() {
	isEven(2)
	isEven(3)
}
