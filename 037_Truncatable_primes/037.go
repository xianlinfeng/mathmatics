package main

import (
	"fmt"
	"math"
	"strconv"
)

// splitInt will split a integer number into a slice
func splitInt(n uint64) []uint64 {
	var nums = []uint64{}
	for i := n % 10; n != 0; i = n % 10 {
		nums = append([]uint64{i}, nums...)
		n /= 10
	}
	return nums
}

// mergeSlice will merge a slice into a integer
func mergeSlice(nums []uint64) uint64 {
	var s string
	for i := range nums {
		s += strconv.Itoa(int(nums[i]))

	}
	n, _ := strconv.Atoi(s)
	return uint64(n)
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

func main() {
	var sum uint64
out:
	for i := uint64(10); i < 1000000; i++ {
		digits := splitInt(i)
		first := digits[0]
		last := digits[len(digits)-1]
		if isPrime(last) && isPrime(first) && isPrime(i) {
			for i := 1; i < len(digits); i++ {
				num1 := mergeSlice(digits[i:])
				num2 := mergeSlice(digits[:i])
				if !isPrime(num1) || !isPrime(num2) {
					continue out
				}
			}
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("sum is ", sum)
}
