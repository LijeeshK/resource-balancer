package balancer

import (
	"fmt"
)

// WeightedElement represent an element with it's weight
type WeightedElement struct {
	Element interface{}
	Weight  int8
}

// WeightedRRBalancer implements BackendBalancer as weighted round robin.
// In this implementation every element is associated with a weight (represented using WeightedElement).
type WeightedRRBalancer struct {
	rrb *RRBalancer
}

// IncompatibleTypeError represents error while using incompatible types in WeightedRoundRobin
type IncompatibleTypeError struct {
	expectedType string
	actualType   interface{}
}

func (ite *IncompatibleTypeError) Error() string {
	return fmt.Sprintf("Expected [%v], acutally got [%t] ", ite.expectedType, ite.actualType)
}

// Load loads weighted elements to balance.
// Incase of incompatible type to load, returns IncompatibleTypeError is returned
func (wrr *WeightedRRBalancer) Load(elements []interface{}) error {

	iElts := make([]interface{}, 0)
	for _, v := range elements {
		we, ok := v.(WeightedElement)
		if !ok {
			return &IncompatibleTypeError{"WeightedElement", v}
		}
		for i := 0; i < int(we.Weight); i++ {
			iElts = append(iElts, we.Element)
		}
	}
	wrr.rrb = &RRBalancer{}
	wrr.rrb.Load(iElts)
	return nil
}

// Next returns backend to use for next request
func (wrr *WeightedRRBalancer) Next() (interface{}, error) {
	return wrr.rrb.Next()
}

// Reload backends from new backend elements
func (wrr *WeightedRRBalancer) Reload(elements []interface{}) error {
	return wrr.Load(elements)
}
