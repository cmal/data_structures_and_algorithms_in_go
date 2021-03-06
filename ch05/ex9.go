// find min in sorted rotated list
package main

import (
	"fmt"
)

func main() {
	// assume increasing array
	// and members >= 3
	arr1 := []int{9,0,1,2,3,5,6,7,8}
	fmt.Println(findMin(arr1))
	arr2 := []int{8,9,0,1,2,3,5,6,7}
	fmt.Println(findMin(arr2))
	arr3 := []int{1,2,3,5,6,7,8,9,0}
	fmt.Println(findMin(arr3))
	arr4 := []int{2,3,5,6,7,8,9,0,1}
	fmt.Println(findMin(arr4))
}

func findMin(arr []int) (index int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		index = (start + end) / 2
		if start + 1 == end {
			if arr[index] > arr[index + 1] {
				index ++
			}
			break
		}
		if arr[index - 1] > arr[index] && arr[index] < arr[index + 1] {
			break
		}

		if arr[index] > arr[end] { // on first group
			start = index + 1
		} else { // on second group
			end = index - 1
		}
	}
	return
}
