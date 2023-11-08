package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	if q.Dequeue() != 1 {
		t.Error("Expected 1")
	}
	if q.Dequeue() != 2 {
		t.Error("Expected 2")
	}
	if q.Peek() != 3 {
		t.Error("Expected 3")
	}
	if q.Size() != 3 {
		t.Error("Expected 3")
	}
}
