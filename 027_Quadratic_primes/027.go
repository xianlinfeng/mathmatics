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

func quadraticFunc(a, b int) (ab int) {
	var primes []int

	for n := 0; ; n++ {
		p := n*n + a*n + b
		if isPrime(p) {
			primes = append(primes, p)
		} else {
			return len(primes)
		}
	}

}

func main() {
	var maxLenght int
	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			l := quadraticFunc(a, b)
			if l > maxLenght {
				maxLenght = l
				fmt.Println("a = ", a, "b = ", b, "l = ", l, "a * b = ", a*b)
			}
		}
	}

}
