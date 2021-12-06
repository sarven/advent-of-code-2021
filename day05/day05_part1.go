package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

const MAX int = 1000
const MIN int = 0

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	lines := make([]Line, 0)

	for _, lineStr := range strings.Split(string(content), "\n") {
		dataLine := strings.Split(string(lineStr), " -> ")
		dataPoint1 := strings.Split(string(dataLine[0]), ",")
		dataPoint2 := strings.Split(string(dataLine[1]), ",")

		x1, _ := strconv.Atoi(dataPoint1[0])
		y1, _ := strconv.Atoi(dataPoint1[1])

		p1 := Point{x: x1, y: y1}

		x2, _ := strconv.Atoi(dataPoint2[0])
		y2, _ := strconv.Atoi(dataPoint2[1])
		p2 := Point{x: x2, y: y2}

		lines = append(lines, Line{p1: p1, p2: p2})
	}

	var iPoints [MAX][MAX]int

	for _, line := range lines {
		iPoints = markIntersection(iPoints, line)
	}
	fmt.Println("")

	for _, row := range iPoints {
		fmt.Println(row)
	}
	fmt.Printf("Result: %d", calcResult(iPoints))
}

func markIntersection(iPoints [MAX][MAX]int, line Line) [MAX][MAX]int {
	if line.p1.x == line.p2.x {
		for i := min(line.p1.y, line.p2.y); i <= max(line.p1.y, line.p2.y); i++ {
			iPoints[i][line.p1.x]++
		}
	}

	if line.p1.y == line.p2.y {
		for i := min(line.p1.x, line.p2.x); i <= max(line.p1.x, line.p2.x); i++ {
			iPoints[line.p1.y][i]++
		}
	}

	return iPoints
}

func calcResult(iPoints [MAX][MAX]int) int {
	sum := 0
	for _, row := range iPoints {
		for _, n := range row {
			if n >= 2 {
				sum++
			}
		}
	}

	return sum
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n2
	}

	return n1
}
