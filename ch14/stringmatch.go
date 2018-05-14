package main

import (
	"fmt"
)

func main() {
	text := "A long string that contains a short string which should be matched: bingo!"
	pattern := "bingo!"

	// matchFunc := BruteForce
	matchFunc := RobinKarp
	// matchFunc := KMP

	fmt.Println(matchFunc(text, pattern))
}

func BruteForce(text string, pattern string) int {

	n := len(text)
	m := len(pattern)

	for i := 0; i < n; i ++ {
		for j := 0; j < m; j ++ {
			if text[i + j] == pattern[j] {
				if j + 1 == m {
					return i
				}
			} else {
				break
			}
		}
	}
	return -1
}

func RobinKarp(text string, pattern string) int {
	n := len(text)
	m := len(pattern)

	powm := 1
	textHash := 0
	patternHash := 0

	prime := 101

	for i := 0; i < m; i ++ {
		powm = (powm << 1) % prime
	}

	for i := 0; i < m; i ++ {
		patternHash = (patternHash << 1 + int(pattern[i])) % prime
		textHash = (textHash << 1 + int(text[i])) % prime
	}

	for i := 0; i <= n - m; i ++ {
		if patternHash == textHash {
			// check collision
			j := 0
			for ; j < m; j ++ {
				if text[i + j] != pattern[j] {
					break;
				}
			}
			if j == m {
				return i
			}
		} else {
			textHash = (textHash << 1 - int(text[i]) * powm + int(text[i + m])) % prime
			if textHash < 0 {
				textHash = textHash + prime
			}
		}

	}

	return -1
}

func KMP(text string, pattern string) int {
	
}
