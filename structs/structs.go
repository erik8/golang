package main

import (
	"fmt"
)

type NameAge struct {
	name string // not exported
	age  int    //not exported
}

func NewNameAge(name string, age int) *NameAge {
	// to long
	n := new(NameAge)
	n.name = name
	n.age = age
	return n
}

func NewNameAge2(name string, age int) *NameAge {
	return &NameAge{name: name, age: age}
}
func main() {
	a := new(NameAge)
	b := NewNameAge("Erik", 23)
	c := NewNameAge2("Jano", 15)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
