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

	h, d := 0, 0

	for _, line := range strings.Split(string(content), "\n") {
		data := strings.Split(line, " ")
		n, err := strconv.Atoi(data[1])

		if nil != err {
			panic(err)
		}

		if "forward" == data[0] {
			h += n
			continue
		}

		if "down" == data[0] {
			d += n
			continue
		}

		if "up" == data[0] {
			d -= n
			continue
		}
	}

	fmt.Printf("%d * %d = %d", h, d, h*d)
}
