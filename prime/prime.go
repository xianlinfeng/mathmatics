package main

import (
	"math"
	"strconv"
)

// splitInt will split a integer number into a slice
func splitInt(n int) []int {
	var nums = []int{}
	for i := n % 10; n != 0; i = n % 10 {
		nums = append([]int{i}, nums...)
		n /= 10
	}
	return nums
}

// mergeSlice will merge a slice into a integer
func mergeSlice(nums []int) int {
	var s string
	for i := range nums {
		s += strconv.Itoa(nums[i])
	}

	n, _ := strconv.Atoi(s)
	return n
}

// IsPrime check if the int is a prime number
func isPrime(n uint64) bool {
	if n <= 3 {
		return n > 1
	}

	// 不在6的倍数两侧的一定不是质数
	if n%6 != 1 && n%6 != 5 {
		return false
	}

	sqrt := math.Sqrt(float64(n))
	for i := uint64(5); i <= uint64(sqrt); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

//筛选法
func sieve(ch chan int, max int) {
	nums := make([]int, max+1)
	ch <- 2
	for i := 3; i < max; i += 2 { //去除偶数
		if nums[i] == 0 {
			ch <- i
			for n := i; n <= max; n += i { // 去除该prime number的倍数
				nums[n] = 1
			}
		}
	}
	close(ch)
}

func main() {
	// isPrime(7)
}
