package main

import (
	"fmt"
	"time"
)

func main() {
	timeChan := time.NewTicker(10 * time.Second).C

	tickChan := time.NewTicker(time.Millisecond * 500).C

	doneChan := make(chan bool)

	go func() {
		time.Sleep(time.Second * 20)
		doneChan <- true
	}()

	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
		case <-tickChan:
			fmt.Println("Ticker ticked")
		case <-doneChan:
			fmt.Println("Done")
			return

		}
	}
}
