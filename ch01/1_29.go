package main

import "fmt"

func main() {

	arr := []int{1,2,3,4,5}

	found := SeqSearch(arr, 3)
	fmt.Println("Found 3: ", found)

	found = SeqSearch(arr, 8)
	fmt.Println("Found 8: ", found)
}

func SeqSearch(arr []int, target int) bool {
	for _, item := range arr {
		if target == item {
			return true
		}
	}
	return false
}
