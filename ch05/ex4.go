// sep 0s and 1s
package main

import (
	"fmt"
)

func main() {
	arr := []int{0,1,1,0,1,1,1,1,1,0,1,0,0,1,0}
	sep(arr)
	fmt.Println(arr)
}

func sep(arr []int) {
	head := 0
	tail := len(arr) - 1

	for {
		for arr[head] == 0 {
			head ++
		}
		for arr[tail] == 1 {
			tail --
		}
		if head < tail && arr[head] == 1 && arr[tail] == 0 {
			// fmt.Println(arr, head, tail)
			swap(arr, head, tail)
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
