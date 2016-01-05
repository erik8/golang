package main

import (
	"fmt"
	"time"
)

var c chan string

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	//the code is blocked on the following line until there is a reader of channel
	c <- w
	fmt.Println("The end of goroutine", w)
}

func main() {

	// buffered channel, where the size of the buffer is 2
	c = make(chan string, 2)

	go ready("Tea", 2)
	go ready("Coffee", 1)
	go ready("Chocolate", 3)
	go ready("Expresso", 4)

	fmt.Println("MAIN: Just waiting but not for too long.. ")
	var i int = 0
	// time.Sleep(5 * time.Second)

L:
	for {
		time.Sleep(4 * time.Second)

		select {
		// the code is blocked on the following line until something is read
		case output := <-c:
			fmt.Printf("The output is %s\n", output)
			i++
			if i > 3 {
				break L
			}

		}
	}
	fmt.Println("MAIN: end of program")
}
