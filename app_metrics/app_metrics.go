package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start_time := time.Now()
	resp, err := http.Get("http://www.example.com/")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	elapsed_time := time.Since(start_time).Seconds()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	fmt.Println()
	fmt.Println(elapsed_time)
}
