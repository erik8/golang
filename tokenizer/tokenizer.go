package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var page []byte = get_page("http://walter-russell.com/")

	// display downloaded page
	//fmt.Printf("%s", string(page))

	tokenizer := html.NewTokenizer(page)

	for {
		tit := tokenizer.Next()

		switch tit {
		case html.ErrorToken:
			// skip error token
			continue
		case html.StartTagToken:
			t := tokenizer.Token()
			isTitle := t.Data == "title"
			if isTitle {
				fmt.Println("We found a title")
			}
		}
	}

}

func get_page(url string) (page []byte) {
	response, err := http.Get(url)
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
	page = b
}
