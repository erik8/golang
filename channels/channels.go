package main

import (
	"fmt"
	"time"
)

//global channel variable, goroutines have access to it
var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}
func main() {
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("MAIN: I am waiting but not too long")

	//Wait for both channels
	<-c
	<-c

	fmt.Println("MAIN: The end of program")
}
