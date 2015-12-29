package main

import (
	"fmt"
	"os"
)

func error_func() error {
	var err error
	if err := os.Chmod("file", 0644); err != nil { // nil is like c's null
		fmt.Println("Error in opening file")
		return err
	}
	return err

}

func quasi_loop() {
	// Looop in disquise, no while and do while loops just goto statement
	i := 0
Here: // first word on the line ending with colon is label

	if i > 5 && i < 10 {
		i++
		goto Here
	}
	if i > 20 {
		return
	}

	fmt.Println(i)

	i++

	goto Here // jump
}

func for_loop() {
	// clasic for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("The sum is %d\n", sum)
	// loop in reverse
	var s string = "hello"
	s_ := []rune(s)
	for i, j := 0, len(s_)-1; i < j; i, j = i+1, j-1 {
		s_[i], s_[j] = s_[j], s_[i] // parallel assignment
	}
	fmt.Println(string(s_))

	for i := 0; i < 10; i++ {
		if i > 5 && i < 8 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Printf("the i is %d\n", i)

	}
}

func main() {
	// structure of if
	var x int = -5
	var y int = 5
	var result int
	if x > 0 {
		result = y
	} else {
		result = x
	}

	fmt.Printf("Result is %d\n", result)

	error_func()

	if true && true { // paranthesis always on this line
		fmt.Println("true")
	}
	if !false {
		fmt.Println("true")
	}

	quasi_loop()

	for_loop()

}
