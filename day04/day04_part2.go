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

type won struct {
	idx        int
	lastNumber int
	set        []map[int]number
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

	winners := make([]won, 0)
	winnersIdx := make([]int, 0)
	winnersCount := 0

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(sets); j++ {
			for x := 0; x < len(sets[j]); x++ {
				for y, num := range sets[j][x] {
					if num.n == numbers[i] {
						sets[j][x][y] = number{n: num.n, marked: true}
					}
				}
			}

			winnersIdx = findWinners(sets)

			if winnersCount < len(winnersIdx) {
				// check if exists
				winners = append(winners, won{idx: j, lastNumber: numbers[i], set: makeCopy(sets[j])})
			}

			winnersCount = len(winners)
		}
	}
	fmt.Printf("lastNumber %d \n", winners[len(winners)-1].lastNumber)
	fmt.Printf("Result: %d", calcResult(winners[len(winners)-1].set, winners[len(winners)-1].lastNumber))
}

func findWinners(sets [][]map[int]number) []int {
	idx := make([]int, 0)
	for i := 0; i < len(sets); i++ {
		if isWinner(sets[i]) {
			idx = append(idx, i)
		}
	}

	return idx
}

func isWinner(set []map[int]number) bool {
	for j := 0; j < len(set); j++ {
		sum := 0

		for _, num := range set[j] {
			if num.marked {
				sum++
			}
		}

		if sum == COLUMNS {
			return true
		}
	}

	for j := 0; j < COLUMNS; j++ {
		sum := 0

		for c := 0; c < COLUMNS; c++ {
			if set[c][j].marked {
				sum++
			}
		}

		if sum == COLUMNS {
			return true
		}
	}

	return false
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

func makeCopy(set []map[int]number) []map[int]number {
	copy := make([]map[int]number, 0)

	for _, row := range set {
		newRow := make(map[int]number, COLUMNS)
		for i, num := range row {
			newRow[i] = number{n: num.n, marked: num.marked}
		}
		copy = append(copy, newRow)
	}

	return copy
}
