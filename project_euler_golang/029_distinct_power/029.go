package main

import (
	"fmt"
	"math/big"
)

func main() {
	var limit int64 = 100
	var sequeue = make(map[string]bool)

	for a := int64(2); a <= limit; a++ {
		for b := int64(2); b <= limit; b++ {
			m := big.NewInt(a)
			m.Exp(m, big.NewInt(b), nil)
			if !sequeue[m.String()] {
				sequeue[m.String()] = true
			} else {
				continue
			}
		}
	}

	// fmt.Println(sequeue)
	fmt.Println(len(sequeue))

}
