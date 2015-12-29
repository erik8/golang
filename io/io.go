/*
the file consists of
$ echo hello > /erik/home/Temp/dat
$ echo go >> /erik/home/Temp/dat
the file looks like
hello\ngoEOF
9 bytes in all
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check_panic(e error) (b bool) {
	// catch panic
	/*
	       defer func() {
	   		if x := recover(); x != nil {
	   			b = true
	   		}
	   	}()
	*/
	if e != nil {
		panic(e)
	}
	return
}

func check_print(e error) bool {
	if e != nil {
		fmt.Println(e)
		return false
	} else {
		return true
	}
}

func main() {
	read_allfile()
	read_buffered()
}

func read_allfile() {
	// read whole file
	dat, err := ioutil.ReadFile("/home/erik/Temp/dat1")
	if !check_panic(err) {
		fmt.Println("Did not panicked")
	} else {
		fmt.Println("Panicked")
	}
	fmt.Print(string(dat))
}

func read_buffered() bool {
	file, err := os.Open("/home/erik/Temp/dat1")
	defer file.Close()

	if !check_print(err) {
		return false
	}

	b1 := make([]byte, 5)
	n1, err := file.Read(b1)
	if !check_print(err) {
		return false
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	n1, err = file.Read(b1)
	if !check_print(err) {
		return false
	}
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// seek the position
	b1 = make([]byte, 5)
	n2, err := file.Seek(6, 0)
	if !check_print(err) {
		return false
	}
	n3, err := io.ReadAtLeast(file, b1, 3)
	if !check_print(err) {
		return false
	}
	fmt.Printf("%d seeked %d bytes: %s\n", n2, n3, string(b1))

	return true
}
