// quick select kth

package main

import (
	"fmt"
)

func main() {
	arr := []int{9,1,8,2,7,3,6,4,5}
	compFunc := more
	fmt.Println(quickSelect(arr, 5, compFunc))

	arr1 := []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 0, compFunc))
	arr1 = []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 1, compFunc))
	arr1 = []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 2, compFunc))
	arr1 = []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 3, compFunc))
	arr1 = []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 4, compFunc))
	arr1 = []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 5, compFunc))
	arr1 = []int{1,2,3,6,7,5,4,0}
	fmt.Println(quickSelect(arr1, 6, compFunc))
}

func less(a, b int) bool {
	return a < b
}

func more(a, b int) bool {
	return a > b
}

func quickSelect(arr []int, k int, comp func(int, int) bool) int {
	length := len(arr)
	tempArr := make([]int, length)

	return quickSelectPart(arr, tempArr, 0, length - 1, k, comp)
}

func quickSelectPart(arr []int, tempArr []int, start int, end int, k int, comp func(int, int) bool) int {
	if start >= end {
		return arr[start]
	}
	// fmt.Println(start, end, k)
	// fmt.Println(arr, tempArr)

	pivot := arr[start]
	head := start
	tail := end

	for i := start + 1; i <= end; i ++ {
		if comp(pivot, arr[i]) {
			tempArr[head] = arr[i]
			head ++
		} else {
			tempArr[tail] = arr[i]
			tail --
		}
	}
	if head == k {
		return pivot
	} else {
		if head > k {
			for i := start; i < head; i ++ {
				arr[i] = tempArr[i]
			}
			return quickSelectPart(arr, tempArr, start, head - 1, k, comp)
		} else {
			for i := tail + 1; i <= end; i ++ {
				arr[i] = tempArr[i]
			}
			return quickSelectPart(arr, tempArr, tail + 1, end, k, comp)
		}

	}
}
