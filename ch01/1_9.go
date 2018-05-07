package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 1, 2, 3:
		fmt.Println("one, two or three")
	default:
		fmt.Println("something else")
	}
}
