package linked_list

import (
	"fmt"
)

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

type LinkedList struct {
	head *node
	tail *node
	size int
}

func (l *LinkedList) AddLast(data int) bool {
	if l.head == nil && l.tail == nil {
		n := create_node(data)
		l.head = n
		l.tail = n
		l.size += 1
		return true
	}

	for cur := l.head; true; {
		if cur.next == nil {
			cur.next = create_node(data)
			l.size++
			l.tail = cur.next
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList) isEmpty() bool {
	if l.head == nil && l.tail == nil {
		return true
	}
	return false
}

func (l *LinkedList) AddFirst(value int) bool {
	if l.isEmpty() {
		l.head = create_node(value)
		l.tail = l.head
		l.size++
		return true
	}

	new_node := create_node(value)
	new_node.next = l.head
	l.head = new_node
	l.size++
	return true
}

func (l *LinkedList) Insert(index int, value int) error {
	if index > l.size {
		return fmt.Errorf("out of index. max is %d", l.size)
	} else if index < 0 {
		return fmt.Errorf("index start from 0")
	}

	if index == 0 {
		l.AddFirst(value)
		return nil
	}

	for cur, i := l.head, index; cur != nil; cur, i = cur.next, i-1 {
		if i == 1 {
			new_node := create_node(value)
			new_node.next = cur.next
			cur.next = new_node
			l.size++
			return nil
		}
	}
	return fmt.Errorf("out of index. max is %d", l.size)
}

func (l *LinkedList) ToArray() []int {
	fmt.Println("size:", l.size)
	sl := make([]int, 0)
	for cur := l.head; cur != nil; cur = cur.next {
		sl = append(sl, cur.data)
	}
	return sl
}

func (l *LinkedList) Print() {
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Println(cur.data)
	}
}

func (l *LinkedList) Remove(index int) error {
	if index > l.size {
		return fmt.Errorf("out of index. max is %d", l.size)
	} else if index < 0 {
		return fmt.Errorf("index start from 0")
	}

	if index == 0 {
		l.head = l.head.next
		l.size--
		return nil
	}

	for cur, i := l.head, index; cur != nil; cur, i = cur.next, i-1 {
		if i == 1 {
			cur.next = cur.next.next
			l.size--
			return nil
		}
	}
	return fmt.Errorf("out of index. max is %d", l.size)
}

func (l *LinkedList) Get(index int) (int, error) {
	if index > l.size {
		return 0, fmt.Errorf("out of index. max is %d", l.size)
	} else if index < 0 {
		return 0, fmt.Errorf("index start from 0")
	}

	for cur, i := l.head, index; cur != nil; cur, i = cur.next, i-1 {
		if i == 0 {
			return cur.data, nil
		}
	}
	return 0, fmt.Errorf("out of index. max is %d", l.size)
}

/*
	A -> B -> C -> D -> E prev = A, cur = B, next = C
	A <- B	  C -> D -> E prev.next = nil, cur.next = prev, next = C
	A <- B <- C    D -> E prev = cur, cur = next, next = next.next
	A <- B <- C <- D    E
	A <- B <- C <- D <- E

	sulphate 1380 per bosta
	brack29 (mann ghosito) 10 KG Bosta -> 880
	unique29 (vitti) 10 kg bosta -> 850
	BADC29 (prottaito) 10KG bosta -> 570
	unique29 2 KG -> 190
	supreme28 2 KG -> 180

*/

func (l *LinkedList) Reverse() {
	if l.head == nil {
		return
	}
	prev, cur := l.head, l.head.next
	prev.next = nil
	l.tail = prev
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}
