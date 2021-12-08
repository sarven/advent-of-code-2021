package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	crabs := make([]int, 0)

	for _, ns := range strings.Split(string(content), ",") {
		n, _ := strconv.Atoi(ns)
		crabs = append(crabs, n)
	}

	distances := make(map[int]int, 0)

	for i := min(crabs); i <= max(crabs); i++ {
		for _, n := range crabs {
			distances[i] += int(math.Abs(float64(i - n)))
		}
	}

	fmt.Println(distances)

	fuel := -1

	for _, d := range distances {
		if -1 == fuel || d < fuel {
			fuel = d
		}
	}

	fmt.Println(fuel)
}

func min(numbers []int) int {
	min := -1

	for _, n := range numbers {
		if -1 == min || n < min {
			min = n
		}
	}

	return min
}

func max(numbers []int) int {
	max := -1

	for _, n := range numbers {
		if -1 == max || n > max {
			max = n
		}
	}

	return max
}
