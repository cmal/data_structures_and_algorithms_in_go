package main

import (
	"fmt"
)

func main() {
	tree := CreateBST([]int{1,2,3,7,6,5,4});
	tree.InOrderTraversePrint()
	fmt.Println("search 3: ", tree.Search(3))
	fmt.Println("search 8: ", tree.Search(8))
	fmt.Println("min: ", tree.Min())
	fmt.Println("max: ", tree.Max())
	fmt.Println("isBST: ", tree.IsBST())
	tree.Remove(4)
	tree.InOrderTraversePrint()
	tree.Remove(7)
	tree.InOrderTraversePrint()
	tree.Remove(1)
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
	tree.root = add(tree.root, val)
}

func add(n *Node, val int) *Node {
	if n == nil {
		return &Node{val, nil, nil}
	} else if n.value > val {
		n.left = add(n.left, val)
	} else {
		n.right = add(n.right, val)
	}
	return n
}

func (tree *Tree) Search(val int) bool {
	if tree.root == nil {
		return false
	}
	return search(tree.root, val)
}

func search(n *Node, val int) bool {
	if n == nil {
		return false
	}
	if n.value == val {
		return true
	}
	return search(n.left, val) || search(n.right, val)
}

func (tree *Tree) Min() int {
	node := tree.root
	return min(node)
}

func min(node *Node) int{
	for node.left != nil {
		node = node.left
	}
	return node.value
}

func (tree *Tree) Max() int {
	node := tree.root
	return max(node)
}

func max(node *Node) int {
	for node.right != nil {
		node = node.right
	}
	return node.value
}

func (tree *Tree) IsBST() bool {
	return isBST(tree.root)
}

func isBST(n *Node) bool {
	if n == nil {
		return true
	}

	if n.left == nil && n.right == nil {
		return true
	}
	
	// assume no equals
	if n.left == nil {
		return n.value < n.right.value && isBST(n.right)
	}

	if n.right == nil {
		return n.left.value < n.value && isBST(n.left)
	}

	return n.left.value < n.value && n.value < n.right.value && isBST(n.left) && isBST(n.right)
}

func (tree *Tree) IsBST1() bool {
	// use infix traverse
	return false
}

func (tree *Tree) Remove(val int) {
	node := tree.root
	tree.root = remove(node, val)
}

func remove(node *Node, val int) *Node {
	if node.value == val {
		// fmt.Printf("%v == %v\n", node.value, val)
		left := node.left
		right := node.right

		if left == nil && right == nil {
			// fmt.Println("return nil")
			node = nil
		} else if left == nil {
			root := min(node.right)
			// fmt.Printf("--> remove right %v\n", root)
			node.right = remove(node.right, root)
			node.value = root
		} else {
			root := max(node.left)
			// fmt.Printf("--> remove left %v\n", root)
			node.left = remove(node.left, root)
			node.value = root
		}
	} else if node.value < val {
		// fmt.Printf("%v < %v\n", node.value, val)
		node.right = remove(node.right, val)
	} else {
		// fmt.Printf("%v > %v\n", node.value, val)
		node.left = remove(node.left, val)
	}
	return node
}


