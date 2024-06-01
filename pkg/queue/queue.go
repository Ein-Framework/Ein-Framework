package queue

import (
	"errors"
)

// Length
func (q *Queue[T]) Length() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.q)
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

// Empty
func (q *Queue[T]) Empty() []*T {
	q.mu.Lock()
	defer q.mu.Unlock()
	old := q.q
	q.q = make([]*T, 0)
	return old
}

func (q *Queue[T]) RemoveIf(condition func(*T) bool) (*T, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for idx, item := range q.q {
		if condition(item) {
			q.q = append(q.q[:idx], q.q[idx+1:]...)
			return item, nil
		}
	}
	return nil, errors.New("item not found")
}

// CreateQueue creates an empty queue
func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{
		q: make([]*T, 0),
	}
}
