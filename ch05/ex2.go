package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1,5,6,7,2,8,3,4,9}

	fmt.Println(find3ForSum(arr, 13))
	fmt.Println(find3ForSum1(arr, 13))
}

func find3ForSum(arr []int, sum int) (int, int, int, bool) {
	// brute force, O(n^3)
	length := len(arr)

	for i := 0; i < length - 2; i ++ {
		for j := i + 1; j < length - 1; j ++ {
			for k := j + 1; k < length; k ++ {
				if arr[i] + arr[j] + arr[k] == sum {
					return arr[i], arr[j], arr[k], true
				}
			}
		}
	}
	return -1, -1, -1, false
}

func find3ForSum1(arr []int, sum int) (int, int, int, bool) {
	sort.Slice(arr, func (i, j int) bool {
		return arr[i] < arr[j]
	})
	// fmt.Println(arr)

	length := len(arr)

	for i := 0; i < length - 2; i ++ {
		j := i + 1
		k := length - 1
		for ; j < length - 1 && k > j; {
			temp := arr[i] + arr[j] + arr[k]
			if temp == sum {
				return arr[i], arr[j], arr[k], true
			} else if temp > sum {
				k --
			} else {
				j ++
			}
		}
	}
	return -1, -1, -1, false
}
