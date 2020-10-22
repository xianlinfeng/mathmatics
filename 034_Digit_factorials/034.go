package main

import "fmt"

func main() {
	var factorial = make(map[int]int, 10)
	factorial[0] = 1
	factorial[1] = 1
	for n := 2; n <= 9; n++ {
		factorial[n] = n * factorial[n-1]
	}

	var result int
	for i := 10; i <= 3265920; i++ { // 9! * 9 = 3265920
		var sum, n = 0, i
		for n > 0 {
			sum += factorial[n%10]
			n /= 10
		}
		if sum == i {
			fmt.Println(sum)
			result += i
		}
	}
	fmt.Println(result)
}
