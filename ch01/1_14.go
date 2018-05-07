package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum is :: ", sum)
}
