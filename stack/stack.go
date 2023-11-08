package stack

import "fmt"

// stack ds
type Stack struct {
	top  *node
	size int
}

// node ds
type node struct {
	data int
	next *node
}

// create a new node
func create_node(value int) *node {
	t := node{
		data: value,
		next: nil,
	}
	return &t
}

// push a new value to the stack
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

// pop the top value from the stack
func (s *Stack) Pop() int {
	if s.top == nil {
		return -1
	}

	top := s.top
	s.top = s.top.next
	s.size--
	return top.data
}

// peek the top value from the stack
func (s *Stack) Peek() int {
	if s.top == nil {
		return -1
	}
	return s.top.data
}

// check if the stack is empty
func (s *Stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

// get the size of the stack
func (s *Stack) Size() int {
	return s.size
}

// print the stack
func (s *Stack) Print() {
	if s.top == nil {
		return
	}

	for cur := s.top; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.data)
	}
	fmt.Println()
}
