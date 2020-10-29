package main

import (
	"strconv"
	"strings"
)

func isPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	for i := 1; i < 10; i++ {
		if !strings.Contains(s, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

func main() {

	max := 0
	var tmp int
	for i := 1; i < 10000; i++ {
		str := ""
		for j := 1; j <= 9 && len(str) < 9; j++ {
			str += strconv.Itoa(i * j)
		}
		if len(str) == 9 && isPandigital(str) {
			tmp, _ = strconv.Atoi(str)
			if tmp > max {
				max = tmp
			}
		}
	}
	println(max)
}
