package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	in := bytes.NewBufferString("1 2 3 f4 5")
	nums := make([]int, 0)

	var d int

	for {
		_, err := fmt.Fscan(in, &d)
		if err == io.EOF {
			break
		}

		if err != nil {
			var s string
			_, err := fmt.Fscan(in, &s)
			if err != nil {
				break
			}
			fmt.Printf("Skipping bad number: %v\n", s)
		} else {
			nums = append(nums, d)
		}

	}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}

}
