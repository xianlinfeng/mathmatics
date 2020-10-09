package main

import (
	"fmt"
	"math"
)

func notPrime(n int) bool {
	if n&2 == 2 {
		return false
	}
	if (n-1)%6 == 0 {
		return false
	}
	if (n+1)%6 == 0 {
		return false
	}
	return true
}

func isPrime(n int) bool {
	if notPrime(n) {
		return false
	}
	max := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 3; i < max; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getFactors(n int) []int {
	var factors []int
out:
	for i := 2; i <= n; i++ {
		if (isPrime(i) && n%i == 0) || (i == n) {
			n /= i
			fmt.Println("n = ", n)
			factors = append(factors, i)
			goto out
		}
	}
	fmt.Println(factors)
	return factors
}

func getMax(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	n := 600851475143
	factors := getFactors(n) // get the factors of the number n
	max := getMax(factors)   // get the max of the factors
	fmt.Println("max factor is ", max)

}
