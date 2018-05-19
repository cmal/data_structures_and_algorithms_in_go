package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type Node struct {
	val int
	links []*Node
}

func main() {
	graph := &Node{1,[]*Node{
		&Node{2, []*Node{
			&Node{4,[]*Node{
				&Node{5,[]*Node{}}}}}},
		&Node{3,[]*Node{
			&Node{6,[]*Node{}}}}}}
	fmt.Println(DFS(graph, 5))
	fmt.Println(DFS(graph, 7))
}

func DFS(node *Node, val int) bool {
	stk := stack.New()

	stk.Push(node)

	for stk.Len() != 0 {
		peek := stk.Peek().(*Node)
		if peek.val == val {
			return true
		}
		top := stk.Pop().(*Node)
		for _, item := range top.links {
			stk.Push(item)
		}
	}
	return false
}
