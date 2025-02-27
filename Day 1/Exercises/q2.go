package main

import (
	"fmt"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.val, " ")
	preOrder(n.left)
	preOrder(n.right)
}

func postOrder(n *Node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Print(n.val, " ")
}

func inOrder(n *Node) {
	if n == nil {
		return
	}
	inOrder(n.left)
	fmt.Print(n.val, " ")
	inOrder(n.right)
}

func main() {
	node_1 := Node{
		val: "a",
	}
	node_2 := Node{
		val: "b",
	}
	node_3 := Node{
		val: "c",
	}
	node_4 := Node{
		val:   "-",
		left:  &node_2,
		right: &node_3,
	}
	root := Node{
		val:   "+",
		left:  &node_1,
		right: &node_4,
	}

	fmt.Println("Preorder:")
	preOrder(&root)
	fmt.Println("\nPostorder:")
	postOrder(&root)
	fmt.Println("\nInorder:")
	inOrder(&root)

}
