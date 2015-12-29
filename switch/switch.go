package main

import (
	"fmt"
)

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func f() {
	fmt.Printf("you are in option 2\n")
}

func main() {
	fmt.Printf("The character is %d\n", unhex('F'))

	var i int = 0

	fmt.Printf("Program starts\n")
START:
	fmt.Printf("Read number 0-3 number\n")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Printf("The error in the input occured")
	}
	switch i {
	case 0:
	case 1:
		fmt.Printf("Option 0 and 1\n")
	case 2:
		f()

	case 3:
		goto END

	default:
		fmt.Printf("Only 0-3 option, illegal\n")
	}
	goto START
END:
	fmt.Printf("End of program\n")
}
