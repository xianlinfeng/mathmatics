package main

import (
	"fmt"
	"math"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	// 不在6的倍数两侧的一定不是质数
	if n%6 != 1 && n%6 != 5 {
		return false
	}
	sqrt := math.Sqrt(float64(n))
	for i := 5; i <= int(sqrt); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func splitInt(n int) []int {
	var nums = []int{}
	for i := n % 10; n != 0; i = n % 10 {
		nums = append([]int{i}, nums...)
		n /= 10
	}
	return nums

}

func mergeSlice(nums []int) int {
	var s string
	for i := range nums {
		s += strconv.Itoa(nums[i])
	}

	n, _ := strconv.Atoi(s)
	return n
}

func rotate(n int) []int {
	var nums = splitInt(n)
	var result []int
	// result = append(result, n)
	for i := 0; i < len(nums); i++ {
		result = append(result, mergeSlice(nums))
		nums = append(nums[1:], nums[0])
	}
	return result
}

func main() {
	var result []int
out:
	for i := 1; i < 1000000; i++ {
		if isPrime(i) {
			nums := rotate(i)
			for _, n := range nums {
				if !isPrime(n) {
					continue out
				}
			}
			result = append(result, i)
			fmt.Println(i)
		}
	}

	fmt.Println("there are", len(result), "prime rotate numbers below 1 million.")
}
