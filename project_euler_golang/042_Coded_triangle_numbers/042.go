package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile() []string {
	buf, err := ioutil.ReadFile("p042_words.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(buf)
	names := strings.SplitN(fileStr, ",", -1)

	var strs []string
	for _, n := range names {
		// fmt.Println(n)
		s := strings.Trim(n, "\"")
		strs = append(strs, s)
	}
	return strs
}

func main() {
	var chars = map[byte]int{'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9, 'J': 10, 'K': 11, 'L': 12, 'M': 13, 'N': 14, 'O': 15, 'P': 16, 'Q': 17, 'R': 18, 'S': 19, 'T': 20, 'U': 21, 'V': 22, 'W': 23, 'X': 24, 'Y': 25, 'Z': 26}
	var strs = readFile()
	var triangle = []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55, 66, 78, 91, 105, 120, 136, 153, 171, 190, 210, 231, 253, 276, 300, 325}

	var names []string
	for _, s := range strs {
		var sum int
		for i := range s {
			sum += chars[s[i]]
		}
		for _, v := range triangle {
			if sum == v {
				names = append(names, s)
			}
		}
	}

	fmt.Println(len(names))

}
