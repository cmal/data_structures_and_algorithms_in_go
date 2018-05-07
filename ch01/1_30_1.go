package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7}
	fmt.Println("Found 3", BinSearch(arr, 3))
	fmt.Println("Found 5", BinSearch(arr, 5))
	fmt.Println("found 7", BinSearch(arr, 7))
	fmt.Println("Found 0", BinSearch(arr, 0))
}

func BinSearch(arr []int, target int) bool {
	var low, high int
	high = len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if (arr[mid] == target) {
			return true
		} else if (arr[mid] < target) {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return false
}
