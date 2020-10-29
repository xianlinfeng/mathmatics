package main

import (
	"fmt"
	"math"
)

func main() {
	var penta = map[uint64]bool{}

	for i := 1; i < 10000; i++ {
		j := i * (3*i - 1) / 2
		penta[uint64(j)] = true
	}

	var min uint64 = math.MaxUint64
	for i := range penta {
		for j := range penta {
			if j > i {
				k := j - i
				l := j + i
				if penta[k] && penta[l] {
					fmt.Println(i, j, k, l)
					if k < min {
						min = k
					}
				}
			}
		}
	}

	fmt.Println(min)
}
