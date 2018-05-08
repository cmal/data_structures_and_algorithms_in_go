// factorial
package main

import "fmt"

func main() {
	fmt.Println("factorial 5 is :: ", Factorial(5))
	
}

func Factorial(i int) int {
	if (i == 1) {
		return i
	}
	return Factorial(i - 1) * i
}
