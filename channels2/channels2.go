package main

import (
	"fmt"
	"runtime"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	// the program is not run in parallel until u set the maximum number of processes to be run on
	runtime.GOMAXPROCS(4)

	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)

	fmt.Println("MAIN: I am waiting, but not too long")

	var i int = 0
	// In case we do not know how many goroutines were run, we receive out
	// only when we receive more than one reply on channel c we will exit the loop
L:
	for {
		select {
		case <-c:
			i++
			if i > 1 {
				break L
			}
		}
	}
	fmt.Println("MAIN: The end of program")
}
