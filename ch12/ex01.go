package main

import (
	"fmt"
	"math/rand"
	"log"
)

type Eight9Hash struct{
	size int
	arr [1E8]bool
}

func main() {
	eh := new(Eight9Hash)
	fmt.Println(eh.getNumber())
	// fmt.Println(eh.getNumber())
	// fmt.Println(eh.getNumber())
	// fmt.Println(eh.getNumber())
	fmt.Println(eh.requestNumber(19727887))
}

func (eh *Eight9Hash) getNumber() int {
	r := rand.Intn(1E8)
	limit := r - 1;
	if limit == -1 {
		limit = 99999999
	}
	for ; r != limit; {
		if !eh.arr[r] {
			eh.arr[r] = true
			return r
		} else {
			r ++
			if r == 1E8 {
				r = 0
			}
		}
	}
	log.Fatal("error: Full Hash set")
	return 0
}


func (eh *Eight9Hash) requestNumber(n int) bool {
	return eh.arr[n]
}
