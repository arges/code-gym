package main
import (
	"fmt"
)

type Node struct {
	Next *Node
	Prev *Node
	Value int
}

type LinkedList struct {
	Head *Node
	Tail *Node
}


func (l *LinkedList) Insert(node *Node) {
	// If nothing has been inserted yet, setup the first element.
	if l.Head == nil && l.Tail == nil {
		l.Head = node
		l.Tail = node
		return
	}

	// Otherwise extend the tail
	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node
}


func (l *LinkedList) Remove(node *Node) {
	var n *Node
	for n = l.Head; n.Next != nil; n = n.Next {
		if n == node {
			p := n.Prev
			p.Next = n.Next
			n.Prev = p
		}
	}
	if n == node {
		p := n.Prev
		p.Next = n.Next
		n.Prev = p
	}

}


func (l *LinkedList) Print() {
	var n *Node
	for n = l.Head; n != nil && n.Next != nil; n = n.Next {
		fmt.Println(n)
	}
	fmt.Println(n)
}


func main() {
	l := LinkedList{}
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	n4 := &Node{Value: 4}
	l.Insert(n1)
	l.Insert(n2)
	l.Insert(n3)
	l.Insert(n4)
	l.Print()
	fmt.Println("---")
	l.Remove(n1)
	l.Remove(n2)
	l.Remove(n3)
	l.Remove(n4)
	l.Print()
}
