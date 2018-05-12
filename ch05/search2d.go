// 5.37
package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1,11,21,31},
		[]int{2,12,22,32},
		[]int{3,13,23,33},
		[]int{4,14,24,34}}

	fmt.Println(find2d(matrix, -1))
	fmt.Println(find2d(matrix, 1))
	fmt.Println(find2d(matrix, 31))
	fmt.Println(find2d(matrix, 2))
	fmt.Println(find2d(matrix, 13))
	fmt.Println(find2d(matrix, 4))
	fmt.Println(find2d(matrix, 34))
	fmt.Println(find2d(matrix, 100))
}

func find2d(matrix[][]int, target int) (row, col int) {
	col = len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		this := matrix[row][col]
		if this == target {
			return
		} else if this < target {
			row ++
		} else {
			col --
		}
	}
	return -1, -1
}
