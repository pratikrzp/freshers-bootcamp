package main

import (
	"fmt"
)

func Express() {
	expression := "a+b-c"
	expTree := ExpressionTree{[]*Node{nil}}
	for _, c := range expression {
		str := string(c)
		switch str {
		case "+", "-":
			fmt.Println("operator")
			left := expTree.Pop()
			right := expTree.Pop()
			newNode := Node{left: left, right: right, value: str}
			expTree.Push(&newNode)
		default:
			fmt.Println("operand")
			newNode := &Node{left: nil, right: nil, value: str}
			expTree.Push(newNode)
		}
	}
}

type ExpressionTree struct {
	tree []*Node
}

func (e *ExpressionTree) Pop() (node *Node) {
	treeSize := len(e.tree)
	if treeSize == 0 {
		return
	}

	newTree := ExpressionTree{}
	lastNode := e.tree[treeSize-1]
	for _, n := range e.tree {
		if lastNode != n {
			newTree.Push(n)
		}
	}
	e.tree = newTree.tree
	return
}
func (e *ExpressionTree) Push(node *Node) {
	treeSize := len(e.tree)
	if treeSize != 0 {
		lastNode := e.tree[treeSize-1]
		lastNode.right = node
		node.left = lastNode
	}
	e.tree = append(e.tree, node)
}

type Node struct {
	value string
	left  *Node
	right *Node
}
