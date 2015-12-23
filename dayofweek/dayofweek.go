package main

import "fmt"

func main() {
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	var i int
	i = Thursday
	fmt.Printf(" This is the constant string %d\n", i)

}
