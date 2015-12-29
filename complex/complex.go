package main

import "fmt"
import "math/cmplx"

/* To pick an argument from multi result function: */
func slice(args ...interface{}) []interface{} {
	return args
}

func main() {

	var c complex128 = 5 + 5i

	fmt.Printf("Value is: %v\n", c)

	fmt.Printf("The log of 5+5i is %v\n", cmplx.Log(c))

	fmt.Printf("Polar value r and theta such that c = r * e**thetai %v, %v\n", slice(cmplx.Polar(c))[0], slice(cmplx.Polar(c))[1])
	fmt.Printf("Polar coordinates of 5 + 5i = c are %v\n", cmplx.Rect(2, 1))

}
