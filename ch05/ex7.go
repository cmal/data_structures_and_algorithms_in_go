// given sorted list, find number, return its index if found,
// otherwise return index it would be at
package main

import (
	"fmt"
)

func main() {
	arr := []int{0,1,2,3,5,6}
	fmt.Println(findIndex(arr, 4))
	fmt.Println(findIndex(arr, 5))
}

func findIndex(arr []int, target int) (index int) {
	// use bin search
	// use interation

	if arr[index] == target {
		return index
	}

	start := 0
	end := len(arr) - 1

	for start < end {
		// fmt.Println("iteration", start, end, index)
		index = (start + end) / 2
		if arr[index] == target {
			break;
		}
		if arr[index] < target && target < arr[index + 1] {
			break;
		}

		if arr[index] >= target {
			end = index - 1
		} else if target >= arr[index + 1] {
			start = index + 1
		}
	}


	return index
}
