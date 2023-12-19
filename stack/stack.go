package stack

import "fmt"

type Stack struct {
	top  *node
	size int
}

type node struct {
	data int
	next *node
}

func create_node(value int) *node {
	t := node{
		data: value,
		next: nil,
	}
	return &t
}

func (s *Stack) Push(value int) bool {
	if s.top == nil {
		s.top = create_node(value)
		s.size++
		return true
	}

	new_node := create_node(value)
	new_node.next = s.top
	s.top = new_node
	s.size++
	return true
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return -1
	}

	top := s.top
	s.top = s.top.next
	s.size--
	return top.data
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.data
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Print() {
	if s.top == nil {
		return
	}

	for cur := s.top; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.data)
	}
	fmt.Println()
}
