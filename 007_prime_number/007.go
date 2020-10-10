package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}

	// 不在6的倍数两侧的一定不是质数
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	sqrt := math.Sqrt(float64(n))
	for i := 5; i <= int(sqrt); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	for i := 0; ; i++ {
		if isPrime(i) {
			count++
			fmt.Printf("No. %d prime: %d\n", count, i)
			if count == 10001 {
				break
			}
		}
	}
}
