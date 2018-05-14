package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type List struct {
	head *Node
	count int
}

func main() {

	list := new(List)

	fmt.Println("is Empty: ", list.isEmpty())
	
	list.head = &Node{1, nil}
	list.count = 1

	list.traverse()

	fmt.Println("is Empty: ", list.isEmpty())

	list.addHead(3)
	list.addHead(4)
	fmt.Println(list.searchVal(3))
	fmt.Println(list.searchVal(1))
	fmt.Println(list.searchVal(4))

	list.traverse()
	list.reverse()
	fmt.Println("reversed: ")
	list.traverse()

	fmt.Println("Size: ", list.size())
	
	list.traverse()
	
	fmt.Println("delete 1: ", list.deleteVal(1))
	list.traverse()
	fmt.Println("delete 4: ", list.deleteVal(4))
	list.traverse()
	fmt.Println("delete 3: ", list.deleteVal(3))
	list.traverse()

	fmt.Println("is Empty: ", list.isEmpty())
	fmt.Println("Size: ", list.size())
	
}

func (list *List) addHead(val int) {
	list.head = &Node{val, list.head}
	list.count ++
}

func (list *List) searchVal(val int) (int, bool) {
	index := 0
	cur := list.head
	for cur != nil {
		if cur.value == val {
			return index, true
		}
		index ++
		cur = cur.next
	}
	return index, false
}

func (list *List) deleteVal(val int) bool {
	cur := list.head
	if cur == nil {
		return false
	}
	if cur.value == val {
		list.head = cur.next
		list.count --
		return true
	}

	for cur.next != nil {
		if cur.next.value == val {
			cur.next = cur.next.next
			list.count --
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *List) size() int {
	return list.count
}

func (list *List) traverse() {
	cur := list.head
	for cur != nil {
		fmt.Printf("%v -> ", cur.value)
		cur = cur.next
	}
	fmt.Printf("<nil>")
	fmt.Println()
}

func (list *List) isEmpty() bool {
	return list.count == 0
}

func (list *List) reverse() {
	if list.head == nil {
		return
	}

	cur := list.head
	next := cur.next
	var prev *Node

	for next != nil {
		cur.next = prev
		prev = cur
		cur = next
		next = next.next
	}
	cur.next = prev
	list.head = cur
}
