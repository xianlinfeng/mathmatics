package main

import (
	"fmt"
	"sort"
)

// findMIn : give a lice and a int n, find the minimum iteger's index which is greater than n
// used to swap tail[index] and permu[i-1]
func findMin(nums []int, n int) int {
	var min = 100
	var index int

	for i, v := range nums {
		if v > n {
			if nums[i] < min {
				min = nums[i]
				index = i
			}
		}
	}
	return index
}

func next(permu []int) []int {
	for i := len(permu) - 1; i >= 1; i-- {
		if permu[i-1] < permu[i] {
			tail := permu[i:]
			minIndex := findMin(tail, permu[i-1])
			permu[i-1], tail[minIndex] = tail[minIndex], permu[i-1]
			sort.Ints(tail)
			return permu
		}
	}
	return nil
}

func main() {
	permu := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// permu := []int{0, 1, 2}

	for i := 1; permu != nil; i++ {
		if i == 1000000 {
			fmt.Println(permu)
			return
		}
		permu = next(permu)
	}

}
