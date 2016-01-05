package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	panic(http.ListenAndServe(":8000", nil))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>XXXXX</title></head><body><a href='http://eeerik9.org'>My Home Page </a> </body> </html>")
}
