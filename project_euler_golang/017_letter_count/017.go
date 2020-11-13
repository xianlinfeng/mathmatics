package main

var digit = map[int]int{
	1: 3, // one
	2: 3, // two
	3: 5, // three
	4: 4, // four
	5: 4, // five
	6: 3, // six
	7: 5, // seven
	8: 5, // eight
	9: 4, // nine
}

var ten = map[int]int{
	11: 6, // eleven
	12: 6, // twelve
	13: 8, // thirteen
	14: 8, // fourteen
	15: 7, // fifteen
	16: 7, // sixteen
	17: 9, // seventeen
	18: 8, // eighteen
	19: 8, // nineteen
	1:  3, // ten
	2:  6, // twenty
	3:  6, // thirty
	4:  5, // forty
	5:  5, // fifty
	6:  5, // sixty
	7:  7, // seventy
	8:  6, // eighty
	9:  6, //ninety
}

func getLetters(n int) int {
	if digit[n] > 0 {
		return digit[n]
	}
	if ten[n] > 0 {
		return ten[n]
	}

	if n == 1000 {
		return 3 + 8 // one thousand
	}

	// frome 100, 200, 300
	if n%100 == 0 {
		return digit[n/100] + 7 // hundred
	}

	if n > 100 {
		return digit[n/100] + 7 + 3 + getLetters(n%100) // 7 + 3 = handred and
	}

	if n%10 == 0 {
		return ten[n/10] // like 10, 20, 30, 40 ...
	}

	return ten[n/10] + digit[n%10] // like 23

}

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += getLetters(i)
	}

	println(sum)
}
