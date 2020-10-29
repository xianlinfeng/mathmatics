package main

import "fmt"

func findDivisors(n uint64) []uint64 {
	var divisors []uint64
	for i := uint64(1); i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func findAbundant() []uint64 {
	var abundant []uint64
	for i := uint64(1); i < 28123; i++ {
		divisors := findDivisors(i)
		var sum uint64
		for _, v := range divisors {
			sum += v
		}
		if sum > i {
			abundant = append(abundant, i)
		}
	}
	return abundant
}

func findNonAbundant() uint64 {
	abundants := findAbundant()
	var sum uint64
out:
	for i := uint64(1); i < 28123; i++ {
		for j := 0; j < len(abundants); j++ {
			for k := j; k < len(abundants); k++ {
				tmp := abundants[j] + abundants[k]
				if i == tmp {
					continue out
				}
			}
		}
		sum += i
	}
	return sum
}
func main() {
	fmt.Println(findNonAbundant())
}
