package main

import (
	"fmt"
)

type Node struct {
	Value int
}

type Stack []*Node

func (q *Stack) Push(val int) {
	*q = append(*q, &Node{val})
}

func (q *Stack) Pop() int {
	x := q.Len() - 1
	n := (*q)[x]
	*q = (*q)[:x]
	return n.Value
}
func (q *Stack) Len() int {
	return len(*q)
}

func (q *Stack) Peek() int {
	return (*q)[q.Len() -1].Value
}

type MinStack struct {
	main Stack
	temp Stack
}

func (s *MinStack) Push(val int) {
	fmt.Printf("Push %d\n", val)
	if s.main.Len() == 0 {
		s.main.Push(val)
	} else if s.main.Peek() < val {
		s.temp.Push(s.main.Pop())
		s.Push(val)
		s.main.Push(s.temp.Pop())
	} else {
		s.main.Push(val)
	}
}

func (s *MinStack) Print() {
	for _, k := range s.main {
		fmt.Printf("%d ", k.Value)
	}
	fmt.Printf("\n")
}

func main() {

	m := MinStack{}
	m.Push(3)
	m.Print()
	m.Push(5)
	m.Print()
	m.Push(1)
	m.Print()
	m.Push(4)
	m.Print()
}
