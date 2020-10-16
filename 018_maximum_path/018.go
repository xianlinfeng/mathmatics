package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("018input.txt")
	if err != nil {
		panic(err)
	}

	fileStr := string(buf)
	lines := strings.SplitN(fileStr, "\n", -1)
	arr := make([][]int, 15)
	for i := range arr {
		line := strings.SplitN(lines[i], " ", -1)
		arr[i] = make([]int, len(line))
		for j := range line {
			arr[i][j], _ = strconv.Atoi(line[j])
		}
	}

	for i := 14; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[i][j] > arr[i][j+1] {
				arr[i-1][j] += arr[i][j]
			} else {
				arr[i-1][j] += arr[i][j+1]
			}
		}
	}

	fmt.Println(arr[0][0])

}
