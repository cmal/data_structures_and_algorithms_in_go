package main

import (
	"fmt"
)

func main() {
	list := new(DoublyLinkedList)
	list.AddHead(1)
	list.AddHead(2)
	list.AddTail(3)

	list.Traverse()

	list.Reverse()
	list.Traverse()

	fmt.Println("isEmpty, size: ", list.IsEmpty(), list.Size())

	fmt.Println(list.SearchVal(2))

	list.DeleteVal(1)
	list.Traverse()
	list.DeleteVal(2)
	list.Traverse()
	list.DeleteVal(3)
	list.Traverse()
	fmt.Println("isEmpty, size:", list.IsEmpty(), list.Size())
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	count int
}


type Node struct {
	value int
	prev *Node
	next *Node
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *DoublyLinkedList) Size() int {
	return list.count
}

func (list *DoublyLinkedList) AddHead(val int) {
	cur := &Node{val, nil, list.head}
	if list.tail == nil {
		list.tail = cur
	}
	list.head = cur
	if cur.next != nil {
		cur.next.prev = cur
	}
	list.count ++
}

func (list *DoublyLinkedList) AddTail(val int) {
	cur := &Node{val, list.tail, nil}
	if list.head == nil {
		list.head = cur
	}
	list.tail = cur
	if cur.prev != nil {
		cur.prev.next = cur
	}
	list.count ++
}

func (list *DoublyLinkedList) DeleteVal(val int) bool {
	cur := list.head
	for cur != nil {
		if cur.value != val {
			cur = cur.next
		} else {
			if cur.prev == nil && cur.next == nil {
				list.tail, list.head = nil, nil
			}
			if cur.prev == nil && cur.next != nil {
				list.head, cur.next.prev = nil, cur.next
			}
			if cur.next == nil && cur.prev != nil {
				list.tail, cur.prev.next = cur.prev, nil
			}
			if cur.prev != nil && cur.next != nil {
				cur.next.prev, cur.prev.next = cur.prev, cur.next
			}
			list.count --
			return true
		}
	}
	return false
}

func (list *DoublyLinkedList) SearchVal(val int) (int, bool) {
	cur := list.head
	index := 0
	for cur != nil {
		if cur.value != val {
			return index, true
		}
		index ++
		cur = cur.next
	}
	return 0, false
}

func (list *DoublyLinkedList) Traverse() {
	cur := list.head
	fmt.Printf("head ->")
	if cur == nil {
		fmt.Printf("\n")
		return
	}
	for cur != nil {
		fmt.Printf(" <- %v -> ", cur.value)
		cur = cur.next
	}
	fmt.Printf("<- tail\n")
}

func (list *DoublyLinkedList) Reverse() {

	cur := list.head

	for cur != nil {
		cur.prev, cur.next = cur.next, cur.prev
		cur = cur.prev
	}

	list.head, list.tail = list.tail, list.head
}
