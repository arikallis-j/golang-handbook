package main

import "fmt"

func test_fsum() {
	N := 1000000
	var A = make([]int, N)
	for i := range A {
		A[i] = i + 1
	}
	var k int = 100
	var m = map[int]int{}

	for i, a := range A {
		if j, ok := m[k-a]; ok {
			fmt.Println(i, j)
		} else {
			m[a] = i
		}
	}
}
