package main

import "fmt"

func main() {
	var triangle = map[uint64]bool{}
	var pentagonal = map[uint64]bool{}
	var hexagonal = map[uint64]bool{}

	for i := uint64(1); i < 100000; i++ {
		t := i * (i + 1) / 2
		p := i * (3*i - 1) / 2
		h := i * (2*i - 1)
		triangle[t] = true
		pentagonal[p] = true
		hexagonal[h] = true
	}

	for i := range triangle {
		if pentagonal[i] && hexagonal[i] {
			fmt.Println(i)
		}
	}
}
