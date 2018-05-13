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

	lengthText := len(text)
	lengthPattern := len(pattern)

	for i := 0; i < lengthText; i ++ {
		for j := 0; j < lengthPattern; j ++ {
			if text[i + j] == pattern[j] {
				if j + 1 == lengthPattern {
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
	lengthText := len(text)
	lengthPattern := len(pattern)

	powm := 1
	textHash := 0
	patternHash := 0

	prime := 101

	for i := 0; i < lengthPattern; i ++ {
		powm = (powm << 1) % prime
	}

	for i := 0; i < lengthPattern; i ++ {
		patternHash = (patternHash << 1 + int(pattern[i])) % prime
		textHash = (textHash << 1 + int(text[i])) % prime
	}

	for i := 0; i <= lengthText - lengthPattern; i ++ {
		if patternHash == textHash {
			// check collision
			j := 0
			for ; j < lengthPattern; j ++ {
				if text[i + j] != pattern[j] {
					break;
				}
			}
			if j == lengthPattern {
				return i
			}
		} else {
			textHash = (textHash << 1 - int(text[i]) * powm + int(text[i + lengthPattern])) % prime
			if textHash < 0 {
				textHash = textHash + prime
			}
		}

	}

	return -1
}
