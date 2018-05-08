// gcd
package main

import "fmt"

func GCD(m int, n int) int {
	if (m == n) {
		return m
	} else if (m < n) {
		return GCD(n - m, m)
	} else {
		return GCD(m - n, m)
		
	}
}

func main() {
	fmt.Println(GCD(8, 12))
	fmt.Println(GCD(12, 9))
	fmt.Println(GCD(1, 3))
	fmt.Println(GCD(3, 1))
}
