package queue

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	mu sync.Mutex
	q  []*T
}

// FifoQueue
type FifoQueue[T any] interface {
	Insert(item *T) error
	Remove() (*T, error)
}

// Insert inserts the item into the queue
func (q *Queue[T]) Insert(item *T) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.q = append(q.q, item)
	return nil
}

// Remove removes the oldest element from the queue
func (q *Queue[T]) Remove() (*T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]
		return item, nil
	}
	return nil, errors.New("Queue is empty")
}

// CreateQueue creates an empty queue
func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]*T, 0),
	}
}
