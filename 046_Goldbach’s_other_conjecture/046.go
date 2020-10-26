package main

import (
	"fmt"
	"math"
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

func main() {
	var oddComp []uint64
	var primes []uint64

	for i := uint64(2); i <= 200000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		} else if !isPrime(i) && i%2 != 0 {
			oddComp = append(oddComp, i)
		}
	}

out:
	for _, i := range oddComp {
		for _, j := range primes {
			for k := uint64(1); k < j; k++ {
				if j < i && i == j+2*k*k {
					fmt.Println("i = ", i, "j = ", j, "k = ", k)
					continue out
				}
			}
		}
		fmt.Println("cannot find a solution for ", i)
		break
	}

}
