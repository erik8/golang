package main

import (
	"fmt"
)

func increment(x int) (y int) {
	y = x + 2
	return
}
func callback(y []int, f func(int) int) {
	for i := 0; i < len(y); i++ {
		y[i] = f(y[i])
	}

	for i := 0; i < len(y); i++ {
		fmt.Printf("%d, ", y[i])
	}
}

func main() {
	a := make([]int, 5)
	callback(a, increment)

}
