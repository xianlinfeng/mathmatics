package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func alphabeticalValue(s string) int {
	var sum int
	for _, i := range s {
		for j, v := range alpha {
			if i == v {
				sum += (j + 1)
			}
		}
	}
	return sum
}

func main() {
	buf, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(buf)

	var sum int
	names := strings.SplitN(fileStr, ",", -1)
	sort.Strings(names)
	for i, n := range names {
		// n = strings.ReplaceAll(n, "\"", "")
		n = strings.Trim(n, "\"")
		al := alphabeticalValue(n)
		total := al * (i + 1)
		fmt.Printf("%s : %d * %d = %d \n", n, al, (i + 1), total)

		sum += total
	}

	fmt.Println(sum)
}
