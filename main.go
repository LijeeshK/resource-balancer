package main

import (
	"fmt"
	"resource-balancer/balancer"
)

func main() {

	// var x balancer.BackendBalancer
	// x = &balancer.RoundRobinBalancer{}
	// elts1 := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(elts1)
	// elts := make([]interface{}, 5)
	// for i, v := range elts1 {
	// 	elts[i] = v
	// }
	// x.Load(elts)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(x.Next())
	// }

	//wes := []balancer.WeightedElement{balancer.WeightedElement{}, balancer.WeightedElement{}}

	x := balancer.RoundRobinBalancer{}

	var y interface{}

	y = x

	//var d balancer.RoundRobinBalancer

	_, ok := y.(balancer.WeightedRoundRobin)

	fmt.Println(ok)
}
