package prime

import "math"

// IsPrime check if the int is a prime number
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
