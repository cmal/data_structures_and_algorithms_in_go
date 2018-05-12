// find the kth smallest element in the union of
// two sorted lists.

package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1,3,5,6,7,8}
	arr2 := []int{0}
	arr3:= []int{-8,-7,-5,-4}

	fmt.Println(findKth(arr1, arr2, 3))
	fmt.Println(findKth(arr1, arr3, 3))
}


func findKth(arr1 []int, arr2 []int, k int) int {
	i1 := 0
	i2 := 0
	l1 := len(arr1)
	l2 := len(arr2)

	a1_ex := false
	a2_ex := false

	for {
		fmt.Println(i1, i2, a1_ex, a2_ex)
		if i1 + i2 + 1 == k {
			if a2_ex || arr1[i1] < arr2[i2] {
				return arr1[i1]
			} else {
				return arr2[i2]
			}
		}
		if (a2_ex || arr1[i1] < arr2[i2]) {
			if i1 == l1 - 1 {
				a1_ex = true
			}
			i1 ++
		} else if (a1_ex || arr1[i1] > arr2[i2]){
			if i2 == l2 - 1 {
				a2_ex = true
			}
			i2 ++
		} else {
			fmt.Println("error")
			return -1
		}

		if a1_ex && a2_ex {
			fmt.Println("error2")
			return -1
		}
	}
}
