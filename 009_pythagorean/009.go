package main

import "fmt"

//毕达哥拉斯三胞胎
//勾股数

func isTriplet(a, b, c int) bool {
	return a*a+b*b == c*c
}

func main() {
	for i := 1; i < 1000; i++ {
		for j := i; j < 1000; j++ { // j := i, because i< j
			for k := j; k < 1000; k++ { // k := j because j < k
				if isTriplet(i, j, k) {
					if i+j+k == 1000 {
						fmt.Printf("%d^2 + %d^2 = %d^2\n", i, j, k)
						fmt.Printf("%d + %d + %d = 1000\n", i, j, k)
						fmt.Printf("%d * %d * %d = %d\n", i*i, j*j, k*k, i*j*k)
						return
					}
				}
			}
		}
	}
}
