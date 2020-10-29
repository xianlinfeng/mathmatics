package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "0"
	for i := 1; i < 1000000; i++ {
		s += strconv.Itoa(i)
		if len(s) >= 1000005 {
			break
		}
	}

	fmt.Println(len(s))

	var result int = 1
	for i := 1; i <= 1000000; i *= 10 {
		r := int(s[i] - '0')
		result *= r
		fmt.Println(r, result)
	}

}
