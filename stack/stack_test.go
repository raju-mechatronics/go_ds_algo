package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if s.Pop() != 5 {
		t.Error("Expected 5")
	}
	if s.Pop() != 4 {
		t.Error("Expected 4")
	}
	if s.Peek() != 3 {
		t.Error("Expected 3")
	}
	if s.Size() != 3 {
		t.Error("Expected 3")
	}
}
