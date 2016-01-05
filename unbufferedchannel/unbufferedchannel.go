package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	//the code is blocked on the following line until there is a reader of channel
	c <- 1
}

func main() {

	c = make(chan int)

	go ready("Tea", 2)
	go ready("Coffee", 1)

	fmt.Println("MAIN: Just waiting but not for too long.. ")
	var i int = 0
	time.Sleep(5 * time.Second)

L:
	for {
		select {
		// the code is blocked on the following line until something is read
		case output := <-c:
			fmt.Printf("The output is %d\n", output)
			i++
			if i > 1 {
				break L
			}

		}
	}
	fmt.Println("MAIN: end of program")
}
