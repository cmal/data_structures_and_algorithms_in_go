package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6}
	fmt.Println(Rotate(arr, 3))
}

func Rotate(arr []int, pos int) []int {
	res := arr[pos:]
	head := arr[:pos]

	for _, item := range head {
		res = append(res, item)
	}
	return res
}
