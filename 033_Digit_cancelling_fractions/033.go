package main

import "fmt"

func main() {
	var result = 1.0

	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			for k := 1; k < i; k++ {
				ki := k*10 + i
				ij := i*10 + j
				if ki*j == ij*k { // if ki/ij = k/j
					fmt.Println(ki, ij, k, j)
					result *= float64(ij) / float64(ki) // get the reciprocal of ki/ij
				}
			}
		}
	}

	println(int(result))
}
