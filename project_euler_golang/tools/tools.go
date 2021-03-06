package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
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

func permute(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	perms := []string{}
	head := string(s[0])
	tail := s[1:]

	for _, perm := range permute(tail) {
		for i := 0; i < len(s); i++ {
			newperm := perm[:i] + head + perm[i:]
			perms = append(perms, newperm)
		}
	}
	return perms
}

func readFile() []string {
	buf, err := ioutil.ReadFile("fileName.txt")
	if err != nil {
		panic(err)
	}
	fileStr := string(buf)
	names := strings.SplitN(fileStr, ",", -1)

	var strs []string
	for _, n := range names {
		fmt.Println(n)
		s := strings.Trim(n, "\"")
		strs = append(strs, s)
	}
	return strs
}

// isPandigital check if the number in string is pandigital or not.
func isPandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	for i := 1; i < 10; i++ {
		if !strings.Contains(s, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}

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
