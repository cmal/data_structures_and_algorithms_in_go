// some sort algorithms
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// sortFunc := bubbleSort
	// sortFunc := insertSort
	// sortFunc := selectionSort
	// sortFunc := selectionSort1
	// sortFunc := mergeSort
	// sortFunc := quickSort

	// compFunc := less
	// compFunc := more


	// ### test for func merge
	// arr1 := []int{4,5,6,1,2,3}
	// merge(arr1, 0, 5, 2, compFunc)

	// arr2 := []int{1,2,3,4,5,6}
	// merge(arr2, 0, 5, 2, compFunc)

	// arr3 := []int{1,2}
	// merge(arr3, 0, 1, 0, compFunc)
	// fmt.Println(arr1, arr2, arr3)
	// ### test for func merge end


	arr := []int{1,2,3,6,7,5,4,0}

	// sortFunc(arr, compFunc)
	bucketSort(arr)
	fmt.Println(arr)


	// compFunc = less

	// arr = []int{1,2,3,6,7,5,4,0}

	// sortFunc(arr, compFunc)
	// fmt.Println(arr)
}

func bubbleSort(arr []int, comp func(int, int) bool) {
	length := len(arr)

	// the nth iteration
	n := 1
	// should compare length - n times
	// n = 1..(length-1)

	// swapped in this iteration
	swapped := false
	// if at the end of some iteration, swapped is false, then algorithm finished

	for ; n < length; n ++ {
		swapped = false
		for i := 0; i < length - n; i ++ {
			if (comp(arr[i], arr[i+1])) {
				swap(arr, i, i + 1)
				swapped = true
			}
		}
		if !swapped {
			break
		}

	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func less(a, b int) bool {
	return a < b
}

func more(a, b int) bool {
	return a > b
}

func insertSort(arr []int, comp func(int, int) bool) {
	length := len(arr)

	for i := 1; i < length; i ++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0; j -- {
			if comp(arr[j], temp) {
				swap(arr, j, j + 1)
			} else {
				break
			}
		}
	}
}


func selectionSort(arr []int, comp func(int, int) bool) {
	// tail := len(arr) - 1 // the selected largest element should be swapped with this
	// decrease it at every iteration
	// you can also use head

	for tail := len(arr) - 1; tail > 0; tail -- {
		// go through index 0~tail

		index := 0 // save the largest element's index for swap

		for i := 1; i < tail; i ++ {
			if comp(arr[i], arr[index]) {
				index = i
			}
		}
		swap(arr, tail, index)
	}
}

func selectionSort1(arr []int, comp func(int, int) bool) {
	for i := 0; i < len(arr) - 1; i ++ {
		index := i
		for j := i + 1; j < len(arr); j ++ {
			if comp(arr[j], arr[index]) {
				index = j
			}
		}
		swap(arr, i, index)
	}
}


func mergeSort(arr []int, comp func(int, int) bool) {
	mergeSortPart(arr, 0, len(arr) - 1, comp)
}

func mergeSortPart(arr []int, start int, end int, comp func(int, int) bool) {
	// fmt.Println(len(arr))
	// fmt.Println("mergeSortPart: start, end: ", start, end)
	// recursively
	if start == end {
		return;
	}
	if start + 1 == end {
		if comp(arr[start], arr[end]) {
			swap(arr, start, end)
		}
	}

	// need function merge

	mid := (start + end) / 2
	mergeSortPart(arr, start, mid, comp)
	mergeSortPart(arr, mid + 1, end, comp)

	merge(arr, start, end, mid, comp)
}

func merge(arr []int, start int, end int, mid int, comp func(int, int) bool) {
	p1 := start
	p2 := mid + 1

	size1 := mid + 1 - start
	size2 := end + 1 - (mid + 1)

	curSize1 := 0
	curSize2 := 0

	tempArr := make([]int, end + 1 - start)

	// fmt.Println(size1, size2)
	// fmt.Println(len(arr), mid, start, end)
	for cursor, index := 0, 0; index < end + 1 - start; index ++ {
		if curSize2 == size2 {
			cursor = p1
			p1 ++
			curSize1 ++
		} else if curSize1 == size1 {
			cursor = p2
			p2 ++
			curSize2 ++
		} else if comp(arr[p1], arr[p2]) {
			cursor = p2
			p2 ++
			curSize2 ++
		} else {
			cursor = p1
			p1 ++
			curSize1 ++
		}

		tempArr[index] = arr[cursor]
		// fmt.Println("p1=>", p1, "p2=>", p2, "curSize1=>", curSize1, "curSize2=>", curSize2, "index=>", index, "cursor=>", cursor)
		// fmt.Println(tempArr)
	}

	// copy to the original array
	for i := start; i <= end; i ++ {
		arr[i] = tempArr[i - start]
	}
}

func quickSort(arr []int, comp func(int, int) bool) {
	length := len(arr)
	rand.Shuffle(length, func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	tempArr := make([]int, length)
	quickSortPart(arr, tempArr, 0, length-1, comp)
}

func quickSortPart(arr []int, tempArr []int, start int, end int, comp func(int, int) bool) {
	// fmt.Println(arr, "start, end:", start, end)
	if (start >= end) {
		return
	}
	pivot := arr[start]

	// fmt.Println(start, end)
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
	tempArr[head] = pivot
	for i := start; i <= end; i ++ {
		arr[i] = tempArr[i]
	}
	// fmt.Println(head, tail)
	quickSortPart(arr, tempArr, start, head - 1, comp)
	quickSortPart(arr, tempArr, tail + 1, end, comp)
}

var bucketLowerBound = 0
var bucketUpperBound = 7

func bucketSort(arr []int) {
	length := len(arr)
	bucketRange := bucketUpperBound - bucketLowerBound + 1
	tempArr := make([]int, bucketRange)

	for i := 0; i < length; i ++ {
		tempArr[arr[i]] += 1
	}

	cnt := 0
	for index, item := range tempArr {
		for i := 0; i < item; i ++ {
			arr[cnt] = index + bucketLowerBound
			cnt ++
		}
	}
}
