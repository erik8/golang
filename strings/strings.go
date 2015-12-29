package main

import "fmt"

func main() {

	var s string = "hello"

	// s[0] = 'c' // This is an error, strings are immutable

	// s := "hello" this is an error, cannot be aasigned more than once

	c := []rune(s) // convert s to array of runes
	/* Rune is an alias to int32, it is utf-8 encoded code point, type is useful when iterating characters in string */

	c[0] = 'c'

	s2 := string(c)

	fmt.Printf("%s\n", s2)

	// Multiline strings
	//put plus always on the first line
	// interpreted string literal
	var s3 string = "Starting string" +
		"Ending string"

	//this string s4 also contains new line character
	// raw string literal
	var s4 string = ` Starting string
            Ending string `

	fmt.Printf("This is a string s3 %s\n", s3)
	fmt.Printf("Thjis is a tring s4 %s\n", s4)
}
