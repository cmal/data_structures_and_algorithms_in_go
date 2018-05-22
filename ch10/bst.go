package main

import (
	"fmt"
)

func main() {
	tree := CreateBST([]int{1,2,3,7,6,5,4});
	tree.InOrderTraversePrint()
}

type Node struct {
	value int
	left,right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) InOrderTraversePrint() {
	printInOrder(t.root)
	fmt.Println()
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Printf(" %v ", n.value)
	printInOrder(n.right)
}

func CreateBST(arr []int) *Tree {
	tree := new(Tree)
	for _, item := range arr {
		tree.Add(item)
	}
	return tree
}

func (tree *Tree) Add(val int) {
	if tree.root == nil {
		tree.root = &Node{val,nil,nil}
	}
	if tree.root > val {
		add(tree.left, val)
	} else {
		add(tree.right, val)
	}
}

func (tree *Tree) add(n *Node, val int) {
	if n == nil {
		n.value = val
	}
	if n.value > val {
		add(n.left, val)
	} else {
		add(tree.right, val)
	}
}

func (tree *Tree) Search(val int) bool {
	
}

func (tree *Tree) Min() int {
	
}

func (tree *Tree) Max() int {
	
}

func (tree *Tree) IsBST() bool {
	
}

func (tree *Tree) IsBST1() bool {
	
}

func (tree *Tree) IsBST2() bool {
	
}

func (tree *Tree) Remove(val int) bool {
	
}

