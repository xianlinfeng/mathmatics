package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var digits []int
	var limit = 5 * math.Pow(9.0, 5.0)

	for i := 2; i < int(limit); i++ {
		s := strconv.Itoa(i)
		var sum float64
		for _, v := range s {
			i, _ := strconv.Atoi(string(v))
			sum += math.Pow(float64(i), float64(5))
		}
		if sum == float64(i) {
			digits = append(digits, i)
		}
	}
	fmt.Println(digits)

	var sum int
	for _, i := range digits {
		sum += i
	}
	fmt.Println("sum = ", sum)
}
