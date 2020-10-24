package main

import (
	"fmt"
	"testing"
)

func TestPrime(t *testing.T) {
	var nums = []uint64{}
	for i := uint64(1); i < 1000000000; i++ {
		if isPrime(i) {
			nums = append(nums, i)
		}
	}
	// fmt.Println(nums)
	fmt.Println(len(nums))
	// 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,   53, 59, 61, 67, 71, 73,   79, 83, 89,  97
	// 2  3  5  7  11  13  17  19  23  29  31  37  41  43  47 49 53  59  61  67  71  73 77 79  83  89 91 97]
}

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(uint64(i))
	}
}
