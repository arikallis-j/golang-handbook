package main

import (
	"fmt"
)

func test_slice() {
	// var mySlice []int
	// mySlice := make([]int, 0)     // слайс [], базовый массив []
	// mySlice := make([]int, 5)     // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	// mySlice := make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]
	// fmt.Println(cap(mySlice))
	// s := []int{1, 2, 3} // [1 2 3]
	// fmt.Println(cap(s))
	// s = s[:len(s)-1] // [1 2]
	// fmt.Println(cap(s))

	// a := []int{1, 2, 3, 4}
	// b := a[2:3] // b = [3]
	// b = append(b, 7)
	// fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	// fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
	// b = append(b, 8, 9, 10)
	// b[0] = 11
	// fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	// fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6
	N, n := 100, 10
	slice := make([]int, N)

	for k := 1; k <= N; k++ {
		slice[k-1] = k
	}
	fmt.Println(slice)
	slice = append(slice[:n], slice[N-n:]...)
	fmt.Println(slice)
	for k := 0; k <= int((len(slice)-1)/2); k++ {
		slice[k], slice[len(slice)-k-1] = slice[len(slice)-k-1], slice[k]
	}
	fmt.Println(slice)

}
