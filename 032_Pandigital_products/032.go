package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPandigital(s string) bool {
	if len(s) != 9 || strings.Contains(s, strconv.Itoa(0)) { // should be exictly 9 numbers
		return false
	}

	var m = make(map[int]bool, len(s))

	for i := 1; i <= len(s); i++ {
		m[i] = true
	}

	n, _ := strconv.Atoi(s)

	for i := n % 10; n != 0; i = n % 10 {
		if m[i] {
			n /= 10
			m[i] = false
		} else {
			return false
		}

	}
	return true
}

func main() {
	var m = map[int]bool{}

	for i := 1; i < 99; i++ {
		for j := i; j < 9999; j++ {
			s := fmt.Sprintf("%d%d%d", i, j, i*j)
			if isPandigital(s) {
				fmt.Printf("i = %d, j = %d, i*j = %d \n", i, j, i*j)
				m[i*j] = true
			}
		}
	}

	var sum int
	for v := range m {
		sum += v
	}

	fmt.Println("m = ", sum)

}
