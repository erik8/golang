package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)



func main() {
	r := bufio.NewReader(os.Stdin)

	str, ok := r.ReadString('\n')

	if ok != nil {
		log.Fatalf("The error is: %s\n", ok.Error())
	}

	str_fields := strings.Fields(str)

	for i := 0; i < len(str_fields); i++ {
		fmt.Println(str_fields[i])
	}

	fmt.Printf("The input string is: %s\n", str)

	fmt.Println("String split *********************")
	str_splits := strings.SplitAfter(str, ", ")

	for i := 0; i < len(str_splits); i++ {
		fmt.Println(str_splits[i])
	}

}
