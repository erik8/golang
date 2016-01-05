package main

import (
	"fmt"
	"time"
)

// right now we wait 5 seconds but in fact we have no idea how long we should wait until all goroutines have exited

// if we did not wait for routines, the program would be terminated immediately and any running routines would die with it.
func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tea", 2)
	go ready("Coffee", 1)

	fmt.Println("MAIN: I am wainting..")
	time.Sleep(5 * time.Second)
	fmt.Println("MAIN: End of program..")
}
