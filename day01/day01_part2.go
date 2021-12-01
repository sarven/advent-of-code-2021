package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	numbers := make(map[int]int)

	for i, line := range strings.Split(string(content), "\n") {
		n, err := strconv.Atoi(line)

		if nil != err {
			panic(err)
		}

		numbers[i] = n
	}

	increasedSum := 0

	for i, _ := range numbers {
		if 3 > i {
			continue
		}
		if numbers[i-2]+numbers[i-1]+numbers[i] > numbers[i-3]+numbers[i-2]+numbers[i-1] {
			increasedSum++
		}
	}

	fmt.Println(increasedSum)
}
