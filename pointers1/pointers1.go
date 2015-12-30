package main

import (
	"fmt"
)

func passbyvalue(i int) {
	i = i + 1
}
func passbyref(i *int) {
	*i = *i + 1
}

func main() {

	// In go there is no pointer arithmetics
	var p *int
	fmt.Printf("%v\n", p) // Points nil

	var i int // Declare integer variable i
	p = &i    // Take the address of i

	fmt.Printf("%v\n", p) // Address of variable

	*p = 8
	fmt.Printf("%v\n", *p) // Print value of pointer
	fmt.Printf("%v\n", i)  //Idem

	passbyvalue(i)
	fmt.Printf("%d\n", i)

	passbyref(&i)
	fmt.Printf("%d\n", i)
}
