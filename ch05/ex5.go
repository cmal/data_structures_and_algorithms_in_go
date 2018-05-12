// separate 0s, 1s and 2s
package main

import (
	"fmt"
)

func main() {
	arr := []int{2,2,0,1,1,0,1,1,1,1,2,1,2,0,1,0,0,1,0}

	countSep(arr)
	fmt.Println(arr)
}

func countSep(arr []int) {
	var zeros, ones int
	length := len(arr)
	for i := 0; i < length; i ++ {
		if arr[i] == 0 {
			zeros ++
		} else if arr[i] == 1{
			ones ++
		}
	}
	i := 0
	for ; i < zeros; i ++ {
		arr[i] = 0
	}
	for ; i < zeros + ones; i ++ {
		arr[i] = 1
	}
	for ; i < length; i ++ {
		arr[i] = 2
	}
}
