package main

import "fmt"

func printGrid(grid [][]int) {
	for _, i := range grid {
		for _, j := range i {
			fmt.Printf("%3d", j)
		}
		fmt.Println()
	}
	fmt.Println()
}

func diagonalSum(grid [][]int) int {
	var size = len(grid)
	var sum int

	for i, j := 0, 0; i < size; {
		sum += grid[i][j]
		i++
		j++
	}

	for i, j := 0, size-1; i < size; {
		sum += grid[i][j]
		i++
		j--
	}
	sum--
	return sum

}

func main() {
	size := 1001

	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}

	// 0 right 1 down 2 left 3 up
	var direction = 0
	var step = 1
	var toggle = 0
	var v, i, j = 1, size / 2, size / 2

	for v < size*size {
		for n := 0; n < step; n++ {
			grid[i][j] = v
			switch direction {
			case 0:
				j++
			case 1:
				i++
			case 2:
				j--
			case 3:
				i--
			}
			v++
		}

		direction = (direction + 1) % 4 // change the direction
		// when change two directions, step++
		if toggle > 0 {
			step++
			toggle = 0
		} else {
			toggle++
		}
	}
	fmt.Println(diagonalSum(grid))
}
