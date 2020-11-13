package main

import (
	"fmt"

	"github.com/frrad/euler"
)

func main() {
	for i := int64(2); i < 999999; i++ {
		if len(euler.Factors(i)) == 4 && len(euler.Factors(i+1)) == 4 && len(euler.Factors(i+2)) == 4 && len(euler.Factors(i+3)) == 4 {
			fmt.Println(i)
			break
		}
	}
}
