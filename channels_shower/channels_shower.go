package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true
}

func shower(ch chan int, quit chan bool) {
L:
	for {
		select {
		case i := <-ch:
			fmt.Printf("The number is %d\n", i)

		case <-quit:

			break L
		}
	}
}

// race condition when main terminates, sometimes 9 is not printed
// solution is adding another channel to finish
/*
func main() {
	ch := make(chan int)
	go shower(ch)
	for i := 0; i < 10; i++ {
		ch <- i

	}
}

func shower(ch chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Printf("The number is %d\n", i)
		}
	}
}
*/
