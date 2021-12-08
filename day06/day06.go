package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	fishes := make([]int, 9)

	for _, ns := range strings.Split(string(content), ",") {
		n, _ := strconv.Atoi(ns)
		fishes[n]++
	}

	for i := 1; i <= DAYS; i++ {
		fishes[0], fishes[1], fishes[2], fishes[3], fishes[4], fishes[5], fishes[6], fishes[7], fishes[8] = fishes[1], fishes[2], fishes[3], fishes[4], fishes[5], fishes[6], fishes[0]+fishes[7], fishes[8], fishes[0]
	}

	sum := 0

	for _, n := range fishes {
		sum += n
	}

	fmt.Println(sum)
}
