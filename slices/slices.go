package main

import (
	"fmt"
)

func print_slice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d, ", slice[i])
	}
	fmt.Println()
}

func main() {

	sl := make([]int, 10)

	for i := 0; i < 5; i++ {
		sl[i] = 1
	}

	fmt.Printf("The length of slice is %d and capacity is %d\n", len(sl), cap(sl))

	fmt.Println("**************")

	// new slices from older arrays
	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[2:4]
	s2 := a[:]
	s3 := a[:4]
	s4 := a[2:4:5]

	print_slice(s1)
	print_slice(s2)
	print_slice(s3)
	print_slice(s4)
	fmt.Println("****************")
	// Slice  is a reference on array with two values length and capacity
	// to copy all use copy function
	/*
	       a2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	   	s5 := make([]int, 6)
	   	n1 := copy(s5, a2[0:])
	   	print_slice(s5)
	   	fmt.Println(n1)
	   	n2 := copy(s5, s5[2:])
	   	print_slice(s5)
	   	fmt.Println(n2)
	*/

	head := 2
	tail := 4
	a2 := make([]int, 5)
	copy(a2, a[head:tail])
	print_slice(a2)

}
