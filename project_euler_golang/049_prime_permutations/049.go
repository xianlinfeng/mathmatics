package main

import (
	"fmt"
	"strconv"

	"github.com/frrad/euler"
)

func main() {
out:
	for n := int64(1000); n <= 9999; n++ {
		for step := int64(1); step <= 8999; step++ {
			if euler.SortInt(n) == euler.SortInt(n+step) && euler.SortInt(n) == euler.SortInt(n+step*2) {
				if euler.IsPrime(n) && euler.IsPrime(n+step) && euler.IsPrime(n+step*2) {
					if n != 1487 {
						fmt.Println(strconv.Itoa(int(n)) + strconv.Itoa(int(n+step)) + strconv.Itoa(int(n+step*2)))
						break out
					}

				}
			}

		}
	}
}
