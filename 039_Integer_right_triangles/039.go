package main

import "fmt"

func main() {

	var m = map[int]int{}
	for p := 1; p <= 1000; p++ {
		for i := 1; i < p; i++ {
			for j := i; j < p; j++ {
				for k := j; k < p; k++ {
					if i+j+k == p && i*i+j*j == k*k {
						m[p]++
						fmt.Println(p, m[p], i, j, k)
					}
				}
			}
		}
	}
	fmt.Println(m)
	var max int
	for p, v := range m {
		if v > max {
			max = v
			fmt.Println("when p = ", p, " got the maximun number of solution, which is  ", max)
		}
	}
}
