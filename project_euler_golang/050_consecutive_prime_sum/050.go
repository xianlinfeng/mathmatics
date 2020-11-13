package main

import (
	"fmt"

	"github.com/frrad/euler"
)

func main() {
	// generate a list of prime numbers
	primes := []int64{}
	for i := int64(2); i < 99999; i++ {
		if euler.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	ans := make([]int, 1000) // create a list to store results

	// algorithm
	for j := 0; j < 100; j++ { // different start point of the prime list, such as primes[0], primes[1]
		sum := int64(0)
	inner:
		for i := 0; i < len(primes); i++ {
			sum += primes[i+j] // start form primes[i+j] to Consecutive sum the primes
			if sum < 1000000 {
				if euler.IsPrime(sum) {
					ans[i] = int(sum) // i is the number of prime to sum, so with bigger i, there is more prime numbers to sum to get the result of sum
				}
			} else {
				break inner
			}
		}
	}

	// get the biggest index of ans if the value is not zero
	for i := len(ans) - 1; i != 0; i-- {
		if ans[i] != 0 {
			fmt.Println(ans[i])
			break
		}
	}

}
