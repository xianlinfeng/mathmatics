package main

import "fmt"

func isDividable(n int) bool {
	primes := []int{3, 7, 11, 13, 17, 19}
	rest := []int{4, 8, 12, 14, 16, 18}
	for _, i := range primes {
		if n%i != 0 {
			return false
		}
	}
	for _, i := range rest {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 2520; ; i += 10 {
		if isDividable(i) {
			fmt.Println(i)
			break
		}
	}
}
