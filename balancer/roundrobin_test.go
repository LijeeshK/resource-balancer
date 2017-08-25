package balancer_test

import (
	"fmt"
	"resource-balancer/balancer"
	"testing"
	"time"
)

func TestRoundRobinSequencial(t *testing.T) {

	x := balancer.RRBalancer{}
	elts1 := [5]int{1, 2, 3, 4, 5}
	elts := make([]interface{}, 5)
	for i, v := range elts1 {
		elts[i] = v
	}
	x.Load(elts)
	total := 100
	ch := make(chan int, total)
	for i := 0; i < total; i++ {
		go func() {
			val, _ := x.Next()
			ch <- val.(int)
		}()
	}
	m := make(map[int]int)
	for i := 0; i < total; i++ {
		m[<-ch]++
	}

	for _, v := range m {
		if v != total/len(elts1) {
			t.Fail()
		}
	}
}

type TestStruct struct {
	val   int
	group int
}

func TestRoundRobinSequenceWithReload(t *testing.T) {

	x := &balancer.RRBalancer{}

	elts := make([]interface{}, 5)
	for i := 0; i < 5; i++ {
		elts[i] = &TestStruct{i, 1}
	}

	elts2 := make([]interface{}, 10)
	for i := 0; i < 10; i++ {
		elts2[i] = &TestStruct{i, 2}
	}

	x.Load(elts)
	time.AfterFunc(100*time.Microsecond, func() {
		x.Reload(elts2)
	})

	total := 10000
	noOfGoRoutine := 100
	ch := make(chan *TestStruct, total)
	for i := 0; i < noOfGoRoutine; i++ {
		go func() {
			for j := 0; j < (total / noOfGoRoutine); j++ {
				val, _ := x.Next()
				ch <- val.(*TestStruct)
				time.Sleep(10 * time.Millisecond)
			}
		}()

	}

	m := make(map[int]int)
	group1Count := 0
	group2Count := 0
	for i := 0; i < total; i++ {
		v := <-ch
		if v.group == 1 {
			group1Count++
		} else {
			group2Count++
		}
		m[(v.group*100+v.val)]++
	}
	fmt.Printf("Total request in group 1: %d, per key possible count : %d\n", group1Count, group1Count/5)
	fmt.Printf("Total request in group 2: %d, per key possible count : %d\n", group2Count, group2Count/10)
	for k, v := range m {
		fmt.Printf("Key : %d, Value: %d\n", k, v)
	}
}
