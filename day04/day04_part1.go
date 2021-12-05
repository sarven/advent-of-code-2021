package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const COLUMNS int = 5

type number struct {
	n      int
	marked bool
}

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	data := make([]string, 0)

	for _, set := range strings.Split(string(content), "\n\n") {
		data = append(data, set)
	}

	numbers := make([]int, 0)

	for _, ns := range strings.Split(string(data[0]), ",") {
		n, err := strconv.Atoi(ns)

		if nil != err {
			panic(err)
		}

		numbers = append(numbers, n)
	}

	sets := make([][]map[int]number, 0)

	for i := 1; i < len(data); i++ {
		set := make([]map[int]number, COLUMNS)

		for j, s := range strings.Split(string(data[i]), "\n") {
			row := make(map[int]number, COLUMNS)

			x := 0
			for _, ns := range strings.Split(string(s), " ") {
				if "" == ns {
					continue
				}

				n, _ := strconv.Atoi(ns)
				row[x] = number{n: n, marked: false}
				x++
			}

			set[j] = row
		}
		sets = append(sets, set)
	}

	winnerIdx := 0
	lastNumber := 0

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(sets); j++ {
			for x := 0; x < len(sets[j]); x++ {
				for y, num := range sets[j][x] {
					if num.n == numbers[i] {
						sets[j][x][y] = number{n: num.n, marked: true}
					}
				}
			}
		}

		winnerIdx = findWinner(sets)

		if -1 != winnerIdx {
			lastNumber = numbers[i]
			break
		}
	}
	fmt.Printf("Winner Idx: %d \n", winnerIdx)
	fmt.Printf("lastNumber %d \n", lastNumber)
	fmt.Printf("Result: %d", calcResult(sets[winnerIdx], lastNumber))
}

func findWinner(sets [][]map[int]number) int {
	for i := 0; i < len(sets); i++ {
		for j := 0; j < len(sets[i]); j++ {
			sum := 0

			for _, num := range sets[i][j] {
				if num.marked {
					sum++
				}
			}

			if sum == COLUMNS {
				return i
			}
		}
	}

	for i := 0; i < len(sets); i++ {
		for j := 0; j < COLUMNS; j++ {
			sum := 0

			for c := 0; c < COLUMNS; c++ {
				if sets[i][c][j].marked {
					sum++
				}
			}

			if sum == COLUMNS {
				return i
			}
		}
	}

	return -1
}

func calcResult(set []map[int]number, lastNumber int) int {
	unmarked := 0

	for _, row := range set {
		for _, num := range row {
			if false == num.marked {
				unmarked += num.n
			}
		}
	}
	fmt.Printf("Unmarked: %d \n", unmarked)

	return lastNumber * unmarked
}
