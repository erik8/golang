package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://walter-russell.com/")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	b, err := ioutil.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		fmt.Println("Error has occured", err.Error())
		os.Exit(1)
	}
	// Print html response
	fmt.Printf("%s", string(b))
}
