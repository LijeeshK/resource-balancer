package balancer_test

import (
	"fmt"
	"resource-balancer/balancer"
	"testing"
)

func TestWeightedRoundRobin(t *testing.T) {

	x := &balancer.WeightedRRBalancer{}
	elts := make([]interface{}, 5)
	for i, v := range []int{1, 2, 3, 4, 5} {
		wt := int8(1)
		if i > 3 {
			wt = int8(2)
		}
		elts[i] = balancer.WeightedElement{Element: v, Weight: wt}
	}
	x.Load(elts)
	total := 102
	ch := make(chan int, total)
	for i := 0; i < total; i++ {
		go func() {
			val, _ := x.Next()
			ch <- val.(int)
		}()
	}
	m := make(map[int]int)
	for i := 0; i < total; i++ {
		v := <-ch
		m[v]++
	}

	for k, v := range m {
		if (k == 5 && v != 34) || (k != 5 && v != 17) {
			fmt.Printf("Key: %v, Value: %v\n", k, v)
			t.Fail()
		}
	}
}
