package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7}
	fmt.Println(BinSearch(arr, 6))
	found := BinSearch(arr, 7)
	fmt.Println("found 7: ", found)
	found = BinSearch(arr, 0)
	fmt.Println("found 0: ", found)
}


// search target in a sorted arr
func BinSearch(arr []int, target int) bool {
	l := len(arr)
	if (l == 1) {
		return target == arr[0]
	} else if (l > 1) {
		index := l / 2
		item := arr[index]
		if (item == target) {
			return true
		} else if (item < target) {
			return BinSearch(arr[index+1:], target)
		} else {
			return BinSearch(arr[:index-1], target)
		}
	} else {
		return false
	}
}

