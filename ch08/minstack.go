package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type MinStack struct {
	stk *stack.Stack
	min *stack.Stack
}

func main() {
	ms := new(MinStack)
	ms.Init()

	fmt.Println(ms.IsEmpty())
	ms.Push(1)
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Push(0)
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Push(1)
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Push(2)
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Push(-1)
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Pop()
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Pop()
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Pop()
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Pop()
	fmt.Println(ms.IsEmpty())
	fmt.Println(ms.Peek())
	fmt.Println(ms.Min())

	ms.Pop()
	fmt.Println(ms.IsEmpty())
}

func (ms *MinStack) Init() {
	ms.stk = stack.New()
	ms.min = stack.New()
}

func (ms *MinStack) IsEmpty() bool {
	return ms.stk.Len() == 0
}

func (ms *MinStack) Peek() int {
	return ms.stk.Peek().(int)
}

func (ms *MinStack) Pop() int {
	ms.min.Pop()
	return ms.stk.Pop().(int)
}

func (ms *MinStack) Push(value int) {
	if ms.min.Len() != 0 {
		peek := ms.min.Peek().(int)
		if peek < value {
			ms.min.Push(peek)
		} else {
			ms.min.Push(value)
		}
	} else {
		ms.min.Push(value)
	}
	ms.stk.Push(value)
}

func (ms *MinStack) Min() int {
	return ms.min.Peek().(int)
}
