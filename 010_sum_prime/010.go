package main

import (
	"fmt"
)

func sieve(ch chan int, max int) {
	var sequeue = make([]int, max+1)
	ch <- 2

	for i := 3; i < max; i += 2 {
		if sequeue[i] == 0 {
			ch <- i
			for j := i; j <= max; j += i {
				sequeue[j] = 1
			}
		}
	}
	close(ch)
}

func main() {
	sum := 0
	max := 2000000
	ch := make(chan int, 100)
	go sieve(ch, max)

	for i := range ch {
		sum += i
	}
	fmt.Println(sum)
}
