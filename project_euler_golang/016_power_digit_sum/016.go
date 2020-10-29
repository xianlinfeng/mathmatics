package main

import (
	"math/big"
	"strconv"
)

func sumDigit(s string) uint64 {
	sum := uint64(0)
	for _, i := range s {
		n, _ := strconv.Atoi(string(i))
		sum += uint64(n)
	}
	return sum
}

func main() {
	n := new(big.Int)
	n.Exp(big.NewInt(2), big.NewInt(1000), nil)
	println(sumDigit(n.String()))
}
