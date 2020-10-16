package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var factorial = new(big.Int).SetInt64(1)
	for i := 100; i >= 1; i-- {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}

	var sum int
	str := factorial.String()
	for _, i := range str {
		n, _ := strconv.Atoi(string(i))
		println(n)
		sum += n
	}

	fmt.Println(sum)
}
