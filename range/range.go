package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "c", "d", "e"} // create a slice

	for key, value := range list {
		fmt.Printf("The character '%s' is at position %d\n", value, key)

	}

	for key, value := range "abcde" {
		fmt.Printf("The character '%c' is at position %d\n", value, key)

	}
}
