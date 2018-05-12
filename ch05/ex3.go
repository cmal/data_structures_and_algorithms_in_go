// separate +ve -ve numbers
package main

import (
	"fmt"
)

func main() {
	arr := []int{1,-1,2,-2,3,-3,5,4,2,-5,-4,-2,-6,7}

	separate(arr)
	// fmt.Println("DONE!")
	fmt.Println(arr)
}

func separate(arr []int) {
	head := 0
	tail := len(arr) - 1
	for {
		for arr[head] > 0 {
			head ++
		}
		for arr[tail] < 0 {
			tail --
		}
		if (arr[head] < 0 && arr[tail] > 0 && head < tail) {
			swap(arr, head, tail)
			// fmt.Println(arr, head, tail)
		} else {
			break;
		}
	}
}

func swap(arr []int, p1 int, p2 int) {
	temp := arr[p1]
	arr[p1] = arr[p2]
	arr[p2] = temp
}
