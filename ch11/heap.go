package main

import (
	"fmt"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8}
}

type Heap struct {
	size int
	arr []int
	isMin bool
}

func (h *Heap) compare(i, j int) bool {
	if h.isMin {
		return h[i] > h[j]
	} else {
		return h[i] < h[j]
	}
}

func NewHeap(arr []int, isMin bool) *Heap {

}
