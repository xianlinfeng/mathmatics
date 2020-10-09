package main

import (
	"fmt"
	"strconv"
)

func isPalindromic(n int) bool {
	str := strconv.Itoa(n)
	rev := ""
	for _, s := range str {
		rev = string(s) + rev
	}
	return str == rev
}

func main() {
	largest := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if isPalindromic(product) && product > largest {
				largest = product
				fmt.Printf("%d = %d * %d \n", product, i, j)
			}
		}
	}
	fmt.Println(largest)
}
