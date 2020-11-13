package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(0)
	for i := int64(1); i <= 1000; i++ {
		n.Add(n, big.NewInt(i).Exp(big.NewInt(i), big.NewInt(i), nil))
	}
	s := n.String()
	fmt.Println(s[len(s)-10:])
}
