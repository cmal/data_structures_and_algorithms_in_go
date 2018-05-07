package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7}
	Rotate(arr, 3)
	fmt.Println(arr)
}

// why this will change arr?
func Rotate(arr []int, pos int) {
	l := len(arr)
	Reverse(arr, 0, pos - 1)
	Reverse(arr, pos, l - 1)
	Reverse(arr, 0, l - 1)
}

// why this will change arr?
func Reverse(arr []int, start int, end int) {
	i := start
	j := end
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i ++
		j --
	}
}
