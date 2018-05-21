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

	fmt.Println("Exercise 1: ====>")
	list.traverse()
	fmt.Println(list.insertkth(1, 3))
	list.traverse()

	fmt.Println(list.insertkth(0, 3))
	list.traverse()

	fmt.Println(list.insertkth(2, 3))
	list.traverse()

	fmt.Println(list.insertkth(0, 4))
	list.traverse()

	fmt.Println(list.insertkth(1, 5))
	list.traverse()

	fmt.Println("Exercise 10: ====>")
	list.reversePrint()

	fmt.Println("Exercise 9: ====>")
	list.traverse()
	middle, err := list.findMiddle()
	fmt.Println("middle, err: ", middle, err)

	list.addHead(7)
	list.traverse()
	middle, err = list.findMiddle()
	fmt.Println("middle, err: ", middle, err)

	list.addHead(8)
	list.traverse()
	middle, err = list.findMiddle()
	fmt.Println("middle, err: ", middle, err)

	list.empty()
	middle, err = list.findMiddle()
	fmt.Println("middle, err: ", middle, err)

	list.addHead(1)
	middle, err = list.findMiddle()
	fmt.Println("middle, err: ", middle, err)

	list1 := new(List)
	list2 := new(List)
	// true
	fmt.Println("reverse Compare: ", list1.reverseCompare(list2))
	list1.addHead(1)
	// false
	fmt.Println("reverse Compare: ", list1.reverseCompare(list2))
	list1.empty()
	list2.addHead(1)
	// false
	fmt.Println("reverse Compare: ", list1.reverseCompare(list2))
	list2.empty()
	list1.addHead(1)
	list1.addHead(2)
	list1.addHead(3)
	list2.addHead(3)
	list2.addHead(2)
	list2.addHead(1)
	list1.traverse()
	list2.traverse()
	// true
	fmt.Println("reverse Compare: ", list1.reverseCompare(list2))
	list1.addHead(4)
	// false
	fmt.Println("reverse Compare: ", list1.reverseCompare(list2))
	list2.addHead(4)
	// false
	fmt.Println("reverse Compare: ", list1.reverseCompare(list2))
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

// exercise 1

func (list *List) insertkth(k int, val int) int {
	if k == 0 {
		list.addHead(val)
		return 1
	} else if list.isEmpty() {
		return -1
	} else {
		return insertkth(list, list.head, k - 1, val)
	}
}

func insertkth(list *List, node *Node, k int, val int) int {
	if node.next == nil {
		return -1
	}
	if k == 0 {
		node.next = &Node{val, node.next}
		list.count ++
		return 1
	} else {
		return insertkth(list, node.next, k - 1, val)
	}
}

// exercise 10

func (list *List) reversePrint() {
	reversePrintNode(list.head)
	fmt.Println()
}

func reversePrintNode(node *Node) {
	if node == nil {
		fmt.Printf("<nil>")
	} else {
		reversePrintNode(node.next)
		fmt.Printf(" <- %v", node.value)
	}
}

// exercise 9

func (list *List) findMiddle() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	
	fast := list.head
	slow := list.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow.value, true
}

// for test

func (list *List) empty() {
	list.head = nil
	list.count = 0
}


// exercise 8

func (list *List) reverseCompare(list1 *List) bool {
	if list.head == nil || list1.head == nil {
		if list.head == nil && list1.head == nil {
			return true
		} else {
			return false
		}
	}
	_, success := reverseCompareNode(list.head, list1.head)
	return success
}

func reverseCompareNode(node *Node, node1 *Node) (*Node, bool) {
	if node1.next == nil {
		return node.next, true
	}
	next, success := reverseCompareNode(node, node1.next)
	if !success {
		return nil, false
	} else if next.value != node1.value {
		return nil, false
	} else {
		return next.next, true
	}

}
