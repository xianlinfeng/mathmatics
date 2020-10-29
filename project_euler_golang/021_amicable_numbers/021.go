package main

import "fmt"

func d(n int) int {
	var divisors []int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	var sum int
	for _, i := range divisors {
		sum += i
	}
	return sum
}

func main() {
	var sum int
	for i := 2; i <= 10000; i++ {
		j := d(i)
		if j > i && d(j) == i {
			fmt.Println("pair ", i, j)
			sum += i
			sum += j
		}
	}

	fmt.Println(sum)
}
