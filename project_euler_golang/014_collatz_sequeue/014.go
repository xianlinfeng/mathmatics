package main

import "fmt"

func collatz(n uint64) uint64 {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func sequeue(start uint64) uint64 {
	var length uint64 = 1
	for {
		start = collatz(start)
		length++
		if start == 1 {
			break
		}
	}
	return length
}

func main() {
	var longgest uint64
	var num uint64
	for i := uint64(1); i < 1000000; i++ {
		length := sequeue(i)
		if length > longgest {
			longgest = length
			num = i
		}
	}

	fmt.Println(num, longgest)
}
