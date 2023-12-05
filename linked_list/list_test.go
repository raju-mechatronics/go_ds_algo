package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := LinkedList[int]{}
	l.AddFirst(1)
	//test if AddFirst is valid or not
	if l.head.data != 1 || l.tail.data != 1 {
		t.Error("AddFirst() is not valid")
	}
	l.AddFirst(2)
	l.AddLast(3)
	if l.tail.data != 3 {
		t.Error("AddLast() is not valid")
	}
	//test if ToArray is valid or not
	arr := l.ToArray()
	if arr[0] != 2 || arr[1] != 1 || arr[2] != 3 {
		t.Error("ToArray() is not valid")
	}
	//test get
	if v, e := l.Get(0); v != 2 || e != nil {
		t.Error("Get() is not valid")
	}
	//test remove
	l.Remove(1)
	arr = l.ToArray()
	if arr[0] != 2 || arr[1] != 3 {
		t.Error("Remove() is not valid")
	}
	//test size
	if l.size != 2 {
		t.Error("Size() is not valid")
	}
	// test reverse
	l.Reverse()
	arr = l.ToArray()
	if arr[0] != 3 || arr[1] != 2 {
		t.Error("Reverse() is not valid")
	}

	t_next := l.Iterator()
	for t_cur, t_ok := t_next(); t_ok == true; t_cur, t_ok = t_next() {
		*t_cur = *t_cur + 1
	}

	//test iterator
	next := l.Iterator()
	for cur, ok := next(); ok == true; cur, ok = next() {
		fmt.Println(ok, *cur)
	}
}
