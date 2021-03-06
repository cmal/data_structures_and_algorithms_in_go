package main

import (
	"fmt"
)

func main() {
	text := "A long string that contains a short string which should be matched: bingo!"
	pattern := "bingo!"
	// pattern := "bingo!!"

	// matchFunc := BruteForce
	// matchFunc := RobinKarp
	// matchFunc := Kmp
	matchFunc := Sunday

	fmt.Println(matchFunc(text, pattern))


	// for test

	// p := "ABCDABD"
	// p := "ABCDABCE"
	// p := "ABCDABDE"
	// p := "DABCDABDE"
	// p := "GCAGAGAG"
	// next := make([]int, len(p) + 1)
	// KMPcompile(p, next)

	
}

func BruteForce(text string, pattern string) int {

	n := len(text)
	m := len(pattern)

	for i := 0; i <= n - m; i ++ {
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
			var tmp int
			if i == n {
				tmp = 0
			} else {
				tmp = int(text[i + m])
			}
			textHash = (textHash << 1 - int(text[i]) * powm + tmp) % prime
			if textHash < 0 {
				textHash = textHash + prime
			}
		}

	}

	return -1
}

func Kmp(text string, pattern string) int {

	n := len(text)
	m := len(pattern)

	shiftArr := make([]int, m + 1)
	KmpCompile(pattern, shiftArr, m)

	i := 0
	j := 0

	for j < n {
		for i > -1 && pattern[i] != text[j] {
			i = shiftArr[i]
		}
		i ++
		j ++
		if i >= m {
			return j - i
			// i = shiftArr[i] // if you want more than one search
		}
	}
	
	return -1
}

func KmpCompile(pattern string, shiftArr []int, m int) {
	shiftArr[0] = -1
	i := 0
	j := -1

	// fmt.Printf(" TABLE:| i| j| shiftArr\n")
	for i < m {
		// fmt.Printf("-------+--+--+----------------------\n")
		// fmt.Printf("before:|%2d|%2d|%v\n", i, j, shiftArr)
		// fmt.Printf("j >= 0 : %v ", j >= 0)
		// if j > -1 {
		// 	fmt.Printf("pattern[i] != pattern[j] : %v\n", pattern[i] != pattern[j])
		// } else {
		// 	fmt.Printf("\n")
		// }
		for j > -1 && pattern[i] != pattern[j] {
			// fmt.Printf("j <- shiftArr[%d]\n", j)
			j = shiftArr[j]
			// fmt.Println("j changed to: ", j)
		}
		i ++
		j ++
		if i < len(pattern) && pattern[i] == pattern[j] {
			// fmt.Printf("i++, j++, shiftArr[i]=shiftArr[j]\n")
			shiftArr[i] = shiftArr[j]
		} else {
			// fmt.Printf("i++, j++, shiftArr[i]=j\n")
			shiftArr[i] = j
		}
		// fmt.Printf(" after:|%2d|%2d|%v\n", i, j, shiftArr)
	}
}


func SundayCompile(pattern string, m int, occ map[rune]int) {
	for _, ch := range pattern {
		occ[ch] = -1;
	}

	for i := 0; i < m; i ++ {
		ch := rune(pattern[i])
		occ[ch] = i
	}
}


func Sunday(text string, pattern string) int {

	n := len(text)
	m := len(pattern)

	occ := make(map[rune]int, m)

	SundayCompile(pattern, m, occ)
	
	for i := 0; i <= n - m; {
		j := 0
		for ; j < m; j ++ {
			// fmt.Println(i + j, j)
			if text[i + j] != pattern[j] {
				break
			}
		}

		if j == m {
			return i
		}

		i += m;
		if i < n {
			i -= occ[rune(text[i])]
		}
		
	}
	return -1
}
