package main

import "fmt"

type Node struct {
	left *Node
	right *Node
	val int
}

type BST struct {
	root *Node
}

func insertNode(n *Node, val int) {
	if val >= n.val {
		if n.right == nil {
			n.right = &Node{val: val}
		} else {
			insertNode(n.right, val)
		}
	} else {
		if n.left == nil {
			n.left = &Node{val: val}
		} else {
			insertNode(n.left, val)
		}
	}
}

func (t *BST) Insert(val int) {
	if t.root == nil {
		t.root = &Node{val: val}
	} else {
		insertNode(t.root, val)
	}
}

func printNode(n *Node, shift , level int) {
	// Print it
	for i := 0; i < shift || shift < 0; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%d (%d)\n", n.val, level)

	// Recurse
	if n.left != nil {
		printNode(n.left, shift - 1, level+1)
	}
	if n.right != nil {
		printNode(n.right, shift + 1, level+1)
	}
}

func (t *BST) Print() {
	printNode(t.root, 10, 0)
}

func main() {

	t := &BST{}
	t.Insert(5)
	t.Insert(0)
	t.Insert(3)
	t.Insert(1)
	t.Insert(1)
	t.Insert(1)
	t.Print()
}
