package main

import "fmt"

/* use long division to see how long the cycle is */
func getCycles(n, d int) int { // n and d are Numerator分子 and denominator分母
	set := map[int]int{}
	length := 0

	for n != 0 {
		if n < d { // 如果分子不为零，分子小于分母， 分子*=10
			n *= 10
			continue
		}
		n = n % d // 取余
		// if the remainder has been encountered before, we've cycled
		if set[n] > 0 {
			return length
		}
		set[n] = 1
		length++
	}

	// if we've gotten here, it doesn't cycle
	return 0
}

func main() {
	var maxLength int
	var maxIndex int
	for i := 2; i <= 1000; i++ {
		l := getCycles(1, i)
		if l > maxLength {
			maxLength = l
			maxIndex = i
		}
	}

	fmt.Println("max length = ", maxLength, "max Index = ", maxIndex)
}
