package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8}
	buildHeap(arr)
	fmt.Println(arr)
}

func buildHeap(arr []int) {
	size := len(arr)
	middle := (size - 1) / 2 // this is clever! and important!
	center := middle

	for ; middle >= 0 ; middle -- {
		// fmt.Println(arr)
		// fmt.Printf("middle: %v, arr[middle]: %v", middle, arr[middle])
		tmp := 2 * middle + 1 // this is important, 2i + 1 and 2i + 2
		tmpIndex := tmp + 1
		if tmpIndex < size && arr[tmp] < arr[tmpIndex] {
			tmp = tmpIndex
		}

		// fmt.Printf(", the value compared: %v", arr[tmp])
		if arr[middle] < arr[tmp] {
			// swap
			// fmt.Printf(", result: swapped")
			arr[middle], arr[tmp] = arr[tmp], arr[middle]
			if tmp <= center {
				middle = tmp + 1 // this is important!
			}
		}
		// fmt.Printf("\n")
		// going forward
	}
}
