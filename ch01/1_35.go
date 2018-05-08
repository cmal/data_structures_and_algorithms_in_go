// Hanoi

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	fmt.Println(num)
	TowersOfHanoi(num)
}

func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are:")
	TOHUtil(num, "A", "C", "B")
}

func TOHUtil(num int, from string, to string, temp string) {
	if num == 0 {
		return
	}
	TOHUtil(num - 1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	TOHUtil(num - 1, temp, to, from)
}
