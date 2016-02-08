package main

import (
	"fmt"
	"strings"
)

var s = "hello <format> world"

func main() {
	w := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case '<', '>', ' ':
			return true
		}
		return false
	})
	fmt.Printf("%q\n", w)
}
