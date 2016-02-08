package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func count(i int) {
	for {
		fmt.Printf("%d\n", i)
		time.Sleep(time.Second)
		i += 1
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	var i int = 0

	go count(i)

	// sleep for 20 seconds
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":4321", nil))
	/*
		<-quitChannel
		log.Println("Done!")
	*/
	log.Println("done")
}
