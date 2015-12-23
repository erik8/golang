package main

import "fmt"

func main() {

	var s string = "hello"

	// s[0] = 'c' // This is an error, strings are immutable

	// s := "hello" this is an error, cannot be aasigned more than once

	c := []rune(s)

	c[0] = 'c'

	s2 := string(c)

	fmt.Printf("%s\n", s2)

}
