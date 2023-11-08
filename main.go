package main

import (
	"ds_algo/stack"
	"fmt"
)

func main() {
	s := stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	s.Print()
	fmt.Println(s.Size())
}
