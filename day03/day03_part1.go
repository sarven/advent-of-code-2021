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

	gData := make([]int, BITS)
	eData := make([]int, BITS)

	g, e := 0.0, 0.0

	for _, line := range strings.Split(string(content), "\n") {
		for j, ns := range strings.Split(string(line), "") {
			n, err := strconv.Atoi(ns)

			if nil != err {
				panic(err)
			}

			if 1 == n {
				gData[j]++
				continue
			}

			if 0 == n {
				eData[j]++
				continue
			}
		}
	}

	for i := 0; i < BITS; i++ {
		if gData[i] > eData[i] {
			g += math.Pow(2, float64(BITS-i-1))
			continue
		}

		e += math.Pow(2, float64(BITS-i-1))
	}

	fmt.Printf("%d * %d = %d", int64(g), int64(e), int64(g*e))

}
