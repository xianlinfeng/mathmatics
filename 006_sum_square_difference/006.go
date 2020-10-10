package main

import (
	"fmt"
)

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func main() {
	s1 := sumOfSquares(100)
	s2 := squareOfSum(100)
	ret := s2 - s1
	fmt.Println("result is ", ret)
}
