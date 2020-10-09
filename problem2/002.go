package main

import "fmt"

func main() {
	a, b := 1, 2
	sum := 0
	for {
		if a <= 4000000 {
			if a%2 == 0 {
				sum += a
				fmt.Println(a)
			}
		} else {
			break
		}
		a, b = b, a+b
	}
	fmt.Println("the result is ", sum)
}
