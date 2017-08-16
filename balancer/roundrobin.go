package balancer

import (
	"sync"
	"sync/atomic"
)

type circularArray struct {
	elements []interface{}
	index    int32
	length   int32
}

func newCircularArray(elements []interface{}) *circularArray {
	return &circularArray{elements: elements, index: 0, length: int32(len(elements))}
}

func (c *circularArray) next() interface{} {
	currIndex := atomic.AddInt32(&c.index, 1)
	if currIndex > c.length {
		currIndex = currIndex % c.length
		atomic.StoreInt32(&c.index, currIndex)
	}
	return c.elements[currIndex-1]
}

// RRBalancer implementation of BackendBalancer based on roundrobin algorithm
// This implementation is thread safe and can be accessed concurrently.
type RRBalancer struct {
	currentBackends atomic.Value
	mu              sync.Mutex
}

// Load prepares a circular list of backend elements
// Load has to be called befor calling Next()
// This implementation protects concurrent access, but will not lock access to Next
func (rrb *RRBalancer) Load(elements []interface{}) error {
	rrb.mu.Lock()
	defer rrb.mu.Unlock()
	rrb.currentBackends.Store(newCircularArray(elements))
	return nil
}

// Next returns next backend element from the circular list and moves the pointer to next
func (rrb *RRBalancer) Next() (interface{}, error) {
	carr1 := rrb.currentBackends.Load().(*circularArray)
	return carr1.next(), nil
}

// Reload reloads the backend list from new elements. It's safe to call concurrently while reading Next()
// This implementation protects concurrent access, but will not lock access to Next
func (rrb *RRBalancer) Reload(elements []interface{}) error {
	rrb.Load(elements)
	return nil
}
