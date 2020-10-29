package main

import "fmt"

func move(i, j int) int {
	if i == 2 || j == 2 {
		return 1
	}

	step := move(i+1, j) + move(i, j+1)

	return step
}

func main() {
	fmt.Println(move(0, 0))
}
