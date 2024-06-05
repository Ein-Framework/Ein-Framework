package queue

import "sync"

type Queue[T any] struct {
	mu sync.Mutex
	q  []*T
}

// FifoQueue
type FifoQueue[T any] interface {
	Length() int
	Empty() []*T
	Insert(item *T) error
	Remove() (*T, error)
	RemoveIf(condition func(*T) bool) (*T, error)
}
