package heap

import (
	"testing"
)

func TestHeapOperations(t *testing.T) {
	// Create a new heap
	h := New()

	// Test inserting elements into the heap
	h.Insert(10)
	h.Insert(5)
	h.Insert(15)
	h.Insert(7)
	h.Insert(12)

	// Test if the heap property is maintained after insertion
	if h.arr[0] != 15 && ((h.arr[1] != 12 && h.arr[2] != 10) || (h.arr[1] != 10 && h.arr[2] != 12)) {
		t.Error("Heap property is not maintained after insertion")
	}

	// Test removing elements from the heap
	h.Remove()
	h.Remove()
	// Test if the heap property is maintained after removal
	if h.arr[0] != 10 && (h.arr[1] != 7 || h.arr[2] != 5) {
		t.Error("Heap property is not maintained after removal")
	}

}
