package main

import (
	"fmt"
	"strconv"
)

func isPalindromic(str string) bool {
	rev := ""
	for _, s := range str {
		rev = string(s) + rev
	}
	return str == rev
}

func main() {
	var sum int
	for i := 0; i < 1000000; i++ {
		if isPalindromic(strconv.Itoa(i)) {
			s := fmt.Sprintf("%b", i) // have to use this way to convert integer to binary
			// strconv.FormatInt(i, 2)  this dont work in this scenario
			if isPalindromic(s) {
				sum += i
				fmt.Println(i)
			}
		}
	}

	fmt.Println("the sum is ", sum)
}
