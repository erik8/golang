package main

import (
	"fmt"
)

func main() {

	var arr [10]int
	arr[0] = 42
	arr[1] = 43
	fmt.Printf("The first element is %d\n", arr[0])

	//another declaration
	a := [...]int{1, 2, 3}
	fmt.Printf("a at 2 is %d\n", a[2])

	// multidimensional arrays
	a2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Printf("a2 at 1,1 is %d\n", a2[1][1])
}
