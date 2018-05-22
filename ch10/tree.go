package tree

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
