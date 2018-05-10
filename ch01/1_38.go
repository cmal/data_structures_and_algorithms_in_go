// all permutations of an int list
package main

import "fmt"

func Permutation(data []int, i int, length int) {
	
}

func swap(data []int, x int, y int) {
	data[x], data[y] = data[y], data[x]
}

func main() {
	data := [5]int{1,2,3,4,5}
	Permutation(data[:], 0, 5)
}
