package main

import "fmt"

func main() {
	const ( // iota is reset to 0
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)
	fmt.Printf("This is c0, c1 and c2 %d, %d, %d\n", c0, c1, c2)

	const (
		a = 1 << iota // a == 1 (iota has been reset)
		b = 1 << iota // b == 2
		c = 1 << iota // c == 4
	)
	fmt.Printf("This is a, b, c %d, %d, %d\n", a, b, c)

	const (
		u         = iota * 42 // u == 0     (untyped integer constant)
		v float64 = iota * 42 // v == 42.0  (float64 constant)
		w         = iota * 42 // w == 84    (untyped integer constant)
	)

	fmt.Printf("This is %d,%9.2f, %d\n", u, v, w)

}
