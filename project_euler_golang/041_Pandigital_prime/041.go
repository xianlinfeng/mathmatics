package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// IsPrime check if the int is a prime number
func isPrime(n uint64) bool {
	if n <= 3 {
		return n > 1
	}

	// 不在6的倍数两侧的一定不是质数
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	sqrt := math.Sqrt(float64(n))
	for i := uint64(5); i <= uint64(sqrt); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// isPandigital check if the number in string is pandigital or not.
func isPandigital(s string) bool {
	// if len(s) != 9 {
	// 	return false
	// }
	for i := 1; i < len(s); i++ {
		if !strings.Contains(s, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

func main() {
	var max uint64

	for i := uint64(7654321); i >= 2; i-- {
		s := strconv.Itoa(int(i))
		// fmt.Println("s is ", s)
		if isPandigital(s) && isPrime(i) {
			if i > max {
				max = i
				fmt.Println("find a bigger one: ", i)
			}
		}
	}
}
