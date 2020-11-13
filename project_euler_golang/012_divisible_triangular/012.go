package main

import "fmt"

func findFactors(n int) int {
	if n%2 != 0 || n%3 != 0 || n%4 != 0 || n%5 != 0 || n%7 != 0 || n%11 != 0 || n%13 != 0 { // filter some non-possible num
		return 0
	}
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return len(factors)
}

func main() {
	triangle := 0
	for i := 1; ; i++ {
		triangle += i
		len := findFactors(triangle)
		// fmt.Println(triangle, len)
		if len > 500 {
			fmt.Println(triangle)
			break
		}
	}
}
