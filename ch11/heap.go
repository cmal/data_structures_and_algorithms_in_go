package main

import (
	"fmt"
	"log"
)

func main() {
	arr := []int{1,2,3,4,5,6,7,8}
	heapMax := NewHeap(arr, false)
	fmt.Println(heapMax)
	heapMax.Insert(9)
	fmt.Println(heapMax)
	heapMax.DeleteRoot()
	fmt.Println(heapMax)
	// heapMax := NewHeap(arr, true)
	// fmt.Println(heapMax)
	arr1 := []int{1,2,3,4,5,6,7}
	heapMax1 := NewHeap(arr1, false)
	fmt.Println(heapMax1)
	heapMax1.Insert(8)
	fmt.Println(heapMax1)

	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	heapMax1.DeleteRoot()
	fmt.Println(heapMax1)
	// The following two lines should fatal
	// heapMax1.DeleteRoot()
	// fmt.Println(heapMax1)

	// test insert
	arr2 := []int{}
	heapMax2 := NewHeap(arr2, false)
	fmt.Println(heapMax2)

	heapMax2.Insert(1)
	fmt.Println(heapMax2)
	heapMax2.Insert(2)
	fmt.Println(heapMax2)
	heapMax2.Insert(5)
	fmt.Println(heapMax2)
	heapMax2.Insert(3)
	fmt.Println(heapMax2)
	heapMax2.Insert(6)
	fmt.Println(heapMax2)
	heapMax2.Insert(4)
	fmt.Println(heapMax2)
	heapMax2.Insert(7)
	fmt.Println(heapMax2)

}

type Heap struct {
	size int
	arr []int
	isMin bool
}


// if return true, means a swap should happen
// specifically, if i < j, and isMin,
// then h.arr[i] should < h.arr[j],
// if i < j and isMax, then h.arr[i] should > h.arr[j]
// otherwise, a swap should happen
func (h *Heap) compare(i, j int) bool {
	if i == j {
		log.Fatal("You are comparing two identical indices!")
	}

	var tmp bool

	if h.isMin {
		tmp = h.arr[i] > h.arr[j]
	} else {
		tmp = h.arr[i] < h.arr[j]
	}

	if i > j {
		return !tmp
	}
	return tmp
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func NewHeap(arr []int, isMin bool) *Heap {
	size := len(arr)
	heap := &Heap{size, arr, isMin}
	if size == 0 {
		return heap
	}
	center := size / 2 - 1

	// sink forward
	for i := center; i >= 0 ; i -- {
		heap.sink(i)
	}
	return heap
}

// use recursion
func (h *Heap) sink(i int) {
	tmp := 2 * i + 1
	tmp1 := tmp + 1

	if tmp1 < h.size && h.compare(tmp, tmp1) {
		tmp = tmp1
	}

	// if h.compare(i, tmp) {
	//	h.swap(i, tmp)
	//	if tmp <= center {
	//		i = tmp + 1 // here tmp + 1, think why!!!???
	//		// the 3rd for clause will do i--
	//		// so tmp + 1 will let i = tmp when
	//		// next loop start
	//	}
	// }

	if tmp < h.size && h.compare(i, tmp) {
		h.swap(i, tmp)
		h.sink(tmp)
	}
}

func (h *Heap) rise(low int) {
	high := (low - 1) / 2
	for low > 0 && h.compare(high, low) {
		h.swap(low, high)
		low = high
		high = (low - 1) / 2
	}
}

func (h *Heap) Insert(i int) {
	h.size ++
	h.arr = append(h.arr, i)

	// low := h.size - 1
	// high := (low - 1) / 2
	// rise
	h.rise(h.size - 1)
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

// this will be getMax when isMax
// and getMin when isMin
func (h *Heap) Peek() (int, bool) {
	if h.IsEmpty() {
		fmt.Println("heap is empty.")
		return 0, false
	}
	return h.arr[0], true
}

func (h *Heap) DeleteRoot() {
	if h.size == 0 {
		log.Fatal("Error: Heap is empty!")
	}

	maxIndex := h.size - 1
	h.swap(0, maxIndex)

	// remove last
	h.arr = h.arr[:maxIndex] // note: this will remove h.arr[maxIndex]
	h.size --

	// sink one
	// for i:= 0 ; i < maxIndex; {
	//	tmp := 2 * i + 1
	//	tmp1 := tmp + 1

	//	if tmp1 < maxIndex && h.compare(tmp, tmp1) {
	//		tmp = tmp1
	//	}

	//	if tmp < maxIndex && h.compare(i, tmp) {
	//		h.swap(i, tmp)
	//		i = tmp
	//	} else {
	//		break;
	//	}
	// }
	h.sink(0)
}
