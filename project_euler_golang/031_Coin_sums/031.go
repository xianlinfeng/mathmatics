package main

import "fmt"

func p31(n int) int {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	ways := make([]int, n+1)
	ways[0] = 1 // way[10] 表示达到10，有多种方法。
	for i := range coins {
		for j := coins[i]; j <= n; j++ { // if coins[i] > n, its impossible to find a solution
			ways[j] += ways[j-coins[i]]
			// 比如 coins[i] = 5, j = 13, ways[13] += ways[13 - 5 ] 即 ways[13] += ways[7]
			// 因为所有7 + coins[i]即可以为13，所以要加上所有为7的可能性。
		}
	}
	return ways[n]
}

func main() {
	fmt.Println(p31(200))
}
