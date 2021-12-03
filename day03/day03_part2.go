package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const BITS int = 12

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	data := make([][]int, 0)

	for _, line := range strings.Split(string(content), "\n") {
		nArr := make([]int, 0)
		for _, ns := range strings.Split(string(line), "") {
			n, err := strconv.Atoi(ns)

			if nil != err {
				panic(err)
			}

			nArr = append(nArr, n)
		}

		data = append(data, nArr)
	}

	oxData := data

	for i := 0; ; i++ {
		if i == BITS || 1 == len(oxData) {
			break
		}

		p, n := 0, 0

		for j := 0; j < len(oxData); j++ {
			if 1 == oxData[j][i] {
				p++
			} else {
				n++
			}
		}

		filtered := make([][]int, 0)

		for j := 0; j < len(oxData); j++ {
			if p >= n && 1 == oxData[j][i] {
				filtered = append(filtered, oxData[j])
				continue
			}

			if p < n && 0 == oxData[j][i] {
				filtered = append(filtered, oxData[j])
				continue
			}
		}

		oxData = filtered
	}

	coData := data

	for i := 0; ; i++ {
		if i == BITS || 1 == len(coData) {
			break
		}

		p, n := 0, 0

		for j := 0; j < len(coData); j++ {
			if 1 == coData[j][i] {
				p++
			} else {
				n++
			}
		}

		filtered := make([][]int, 0)

		for j := 0; j < len(coData); j++ {
			if p < n && 1 == coData[j][i] {
				filtered = append(filtered, coData[j])
				continue
			}

			if p >= n && 0 == coData[j][i] {
				filtered = append(filtered, coData[j])
				continue
			}
		}

		coData = filtered
	}

	o := convertToDec(oxData[0])
	c := convertToDec(coData[0])

	fmt.Println(oxData)
	fmt.Println(coData)
	fmt.Printf("%d * %d = %d", o, c, o*c)
}

func convertToDec(data []int) int {
	d := 0

	for i, n := range data {
		if 0 == n {
			continue
		}

		d += int(math.Pow(2, float64(BITS-i-1)))
	}

	return d
}
