package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8}
	buildHeap(&arr)
	fmt.Println(arr)
}

func buildHeap(arr *[]int) {
	size := len(arr)
	middle := size / 2 // this is clever!

	for ; middle > 0 ; middle -- {
		tmp := 2 * middle
		if arr[2 * middle] < arr[2 * middle + 1] {
			tmp += 1
		}

		if arr[middle] < arr[tmp] {
			// swap
			arr[middle], arr[tmp] = arr[tmp], arr[middle]
		}
		// going forward
	}
}
