package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	sum := SumArray(arr)
	fmt.Println("sum: ", sum)
}

func SumArray(arr []int) int {
	var sum int
	for _, item := range arr {
		sum += item
	}
	return sum
}
