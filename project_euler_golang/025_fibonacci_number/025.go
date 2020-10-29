package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)

	for i := 2; ; i++ {
		a, b = b, a.Add(a, b)
		s := a.String()
		fmt.Println(s)
		if len(s) == 1000 {
			fmt.Println(i)
			break
		}

	}
}
