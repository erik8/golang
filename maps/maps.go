package main

import (
	"fmt"
)

func main() {
	// initialize map
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}

	// define map
	// monthdays := make(map[string] int)
	fmt.Printf("%d\n", monthdays["Dec"])

	// looping over map
	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Printf("The number of days in a year is: %d\n", year)

	//adding elements to the map
	monthdays["Imaginary"] = 30
	monthdays["Feb"] = 29 // Override entry for leap years

	/* to test for existence, you could use the following */
	var value int
	var present bool

	value, present = monthdays["Jan"]

	if present {
		fmt.Printf("January has %d days \n", value)
	}

	v, ok := monthdays["Jan"]

	if ok {
		fmt.Printf("January has %d days \n", v)
	}

	// delete imaginary month
	delete(monthdays, "Imaginary")

}
