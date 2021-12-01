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

	for i, n := range numbers {
		if 0 == i {
			continue
		}
		if n > numbers[i-1] {
			increasedSum++
		}
	}

	fmt.Println(increasedSum)
}
