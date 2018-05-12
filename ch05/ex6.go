// find the pos the list becomes positive
package main

import (
	"fmt"
)

func main() {
	arr := []int{-6, -5, -4, -3, -1, 2, 3, 9}
	fmt.Println(findFirstPositive(arr))
	arr = []int{-1, 2}
	fmt.Println(findFirstPositive(arr))
}

func findFirstPositive(arr []int) int {
	return findFirstPositive1(arr, 0, len(arr))
}

func findFirstPositive1(arr []int, start int, end int) int {
	// use bin search
	// use recursive

	if start == end {
		return arr[start + 1]
	}

	mid := (start + end) / 2

	if arr[mid] < 0 && arr[mid + 1] > 0 {
		return arr[mid + 1]
	}

	if arr[mid] > 0 {
		return findFirstPositive1(arr, start, mid - 1)
	} else { // arr[mid] < 0 && arr[mid + 1] < 0
		return findFirstPositive1(arr, mid + 1, end)
	}


}
