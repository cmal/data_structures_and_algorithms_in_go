package main

import "fmt"

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}

func main() {
	i := 10
	fmt.Println("Value of i before increment is : ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is : ", i)
}
