package main

import (
	"fmt"
	"math"
)

func main() {
	// Print constants

	fmt.Printf("Euler number is %v\n", math.E)
	fmt.Printf("Pi number is %v\n", math.Pi)
	fmt.Printf("Phi number is %v\n", math.Phi)

	fmt.Printf("Sqrt2 is %v\n", math.Sqrt2)

	fmt.Printf("SqrtE is %v\n", math.SqrtE)

	fmt.Printf("SqrtPi is %v\n", math.SqrtPi)
	fmt.Printf("SqrtPhi is %v\n", math.SqrtPhi)

	fmt.Printf("Natural logarithm of 2 is %v\n", math.Ln2)
	fmt.Printf("Natural logarithm of 20 is %v\n", math.Ln10)

	fmt.Printf("Max Float 32 is %v\n", math.MaxFloat32)
	fmt.Printf("SmallestNonZeroFloat32 is %v\n", math.SmallestNonzeroFloat32)

	fmt.Printf("Max float 64 is %v\n", math.MaxFloat64)
	fmt.Printf("SmallestNonZeroFloat64 is %v\n", math.SmallestNonzeroFloat64)

	var a float64 = -5
	fmt.Printf(" Absolute value of -5 is %g\n", math.Abs(a))

	var b float64 = 0.5
	fmt.Printf(" Acosine of 0.5 is %g\n", math.Acos(b))
	fmt.Printf(" Asine of 0.5 os %g\n", math.Asin(b))
	fmt.Printf(" Atanh of 0.5 is %g\n", math.Atanh(b))
	fmt.Printf(" A least integer value greater than 0.5 is %v\n", math.Ceil(b))

	fmt.Printf(" e**0.5 is %g\n", math.Exp(0.5))
	fmt.Printf(" max value of -5 and 0.5 is %g\n", math.Max(a, b))
	fmt.Printf(" min value of -5 and 0.5 is %g\n", math.Min(a, b))

	var c float64 = 27
	fmt.Printf("A cube root of 27 is %g\n", math.Cbrt(c))
	fmt.Printf("Sqare root of 27 is %g\n", math.Sqrt(c))
}
