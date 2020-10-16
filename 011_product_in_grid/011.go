package main

import "fmt"

var grid = [][]int{
	[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	[]int{1, 1, 1, 8, 02, 22, 97, 38, 15, 0, 40, 0, 75, 04, 05, 07, 78, 52, 12, 50, 77, 91, 8, 1, 1, 1},
	[]int{1, 1, 1, 49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 0, 1, 1, 1},
	[]int{1, 1, 1, 81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65, 1, 1, 1},
	[]int{1, 1, 1, 52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91, 1, 1, 1},
	[]int{1, 1, 1, 22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80, 1, 1, 1},
	[]int{1, 1, 1, 24, 47, 32, 60, 99, 3, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50, 1, 1, 1},
	[]int{1, 1, 1, 32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70, 1, 1, 1},
	[]int{1, 1, 1, 67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21, 1, 1, 1},
	[]int{1, 1, 1, 24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72, 1, 1, 1},
	[]int{1, 1, 1, 21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95, 1, 1, 1},
	[]int{1, 1, 1, 78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 04, 62, 16, 14, 9, 53, 56, 92, 1, 1, 1},
	[]int{1, 1, 1, 16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57, 1, 1, 1},
	[]int{1, 1, 1, 86, 56, 0, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58, 1, 1, 1},
	[]int{1, 1, 1, 19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40, 1, 1, 1},
	[]int{1, 1, 1, 04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66, 1, 1, 1},
	[]int{1, 1, 1, 88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69, 1, 1, 1},
	[]int{1, 1, 1, 04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36, 1, 1, 1},
	[]int{1, 1, 1, 20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16, 1, 1, 1},
	[]int{1, 1, 1, 20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54, 1, 1, 1},
	[]int{1, 1, 1, 01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48, 1, 1, 1},
	[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

func printGrid(g [][]int) {
	for i := range g {
		for j := range g[i] {
			fmt.Printf("%2d ", g[i][j])
		}
		fmt.Println()
	}
}

func horizontal(g [][]int) int {
	max := 0
	for i := 3; i <= 22; i++ {
		for j := 3; j <= 22; j++ {
			product := g[i][j] * g[i][j+1] * g[i][j+2] * g[i][j+3]
			fmt.Println("product = ", product)
			if product > max {
				max = product
			}
		}
	}
	return max
}

func vertical(g [][]int) int {
	max := 0
	for j := 3; j <= 22; j++ {
		for i := 3; i <= 22; i++ {
			product := g[i][j] * g[i+1][j] * g[i+2][j] * g[i+3][j]
			fmt.Println("product = ", product)
			if product > max {
				max = product
			}
		}
	}
	return max
}

func diagonal1(g [][]int) int {
	max := 0
	for i := 3; i <= 22; i++ {
		for j := 3; j <= 22; j++ {
			product := g[i][j] * g[i+1][j+1] * g[i+2][j+2] * g[i+3][j+3]
			fmt.Println("product = ", product)
			if product > max {
				max = product
			}
		}
	}
	return max
}
func diagonal2(g [][]int) int {
	max := 0
	for i := 3; i <= 22; i++ {
		for j := 3; j <= 22; j++ {
			product := g[i][j] * g[i-1][j+1] * g[i-2][j+2] * g[i-3][j+3]
			fmt.Println("product = ", product)
			if product > max {
				max = product
			}
		}
	}
	return max
}

func findMax(nums []int) int {
	max := 0
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	return max
}

func main() {
	a := horizontal(grid)
	b := vertical(grid)
	c := diagonal1(grid)
	d := diagonal2(grid)
	var nums = []int{a, b, c, d}
	fmt.Println(a, b, c, d)
	fmt.Println("the largest product is ", findMax(nums))
}