package queue

// Queue (FIFO)
type Queue struct {
	items []int
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes an item from the queue
func (q *Queue) Dequeue() int {
	first_item := q.items[0]
	q.items = q.items[1:]
	return first_item
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// IsEmpty returns whether the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Peek returns the first item in the queue
func (q *Queue) Peek() int {
	return q.items[0]
}
