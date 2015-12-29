package main

import (
	"fmt"
)

func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d ", i)
}

// function with multiple return values

func main() {
	rec(0)
	fmt.Println()

	//functions as an value
	a := rec
	a(2)
}
