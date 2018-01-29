package structures

import "sync"

// Queue is a collection designed for holding elements prior to processing.
type Queue struct {
	head  *snode
	tail  *snode
	count int
	lock  *sync.Mutex
}

// NewQueue creates a new Queue
func NewQueue() *Queue {
	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q
}

// Len returns the size of the queue
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.count
}

// Peek returns the head of the queue without removing it from the queue
func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.head == nil {
		return nil
	}
	return q.head.data
}

// Push adds an item to the end of the queue
func (q *Queue) Push(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := &snode{data: item}

	if q.tail == nil {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
	q.count++
}

// Poll removes an item from the head of the queue and returns it
func (q *Queue) Poll() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next
	q.count--

	if q.head == nil {
		q.tail = nil
	}

	return n.data
}
