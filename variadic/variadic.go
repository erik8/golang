package main

import (
	"fmt"
)

func var1(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}
func var2(arg ...int) {
	var1(arg...)
	var1(arg[1:]...)
}
func main() {
	var2(1, 2, 3, 4, 5)
}
