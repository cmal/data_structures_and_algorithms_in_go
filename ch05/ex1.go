// find first repeated in a list

package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1,2,3,1,2,3}
	arr2 := []int{1,1,0,0}
	arr3 := []int{1,2,3,4,5}
	fmt.Println(findFirstRepeat(arr1))
	fmt.Println(findFirstRepeat(arr2))
	fmt.Println(findFirstRepeat(arr3))
}


func findFirstRepeat(arr []int) (int, bool) { // return index, found
	table := make(map[int]bool, len(arr))
	for index, item := range arr {
		if !table[item] {
			table[item] = true
		} else {
			return index, true
		}
	}
	return -1, false
}
