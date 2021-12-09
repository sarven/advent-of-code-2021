package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	count := 0

	for _, line := range lines {
		chunks := strings.Split(string(line), " | ")

		signals := make([]string, 0)
		outputs := make([]string, 0)

		for _, chunk := range strings.Split(string(chunks[0]), " ") {
			signals = append(signals, chunk)
		}

		for _, chunk := range strings.Split(string(chunks[1]), " ") {
			outputs = append(outputs, chunk)
		}

		dictionary := createDictionary(signals)

		for _, output := range outputs {
			if _, ok := dictionary[sortString(output)]; ok {
				count++
			}
		}

		fmt.Println(dictionary)
	}

	fmt.Printf("Count: %d \n", count)
}

func createDictionary(signals []string) map[string]int {
	dictionary := make(map[string]int, 9)

	for _, signal := range signals {
		sortedSignal := sortString(signal)
		if 4 == len(sortedSignal) {
			dictionary[sortedSignal] = 4
		}

		if 3 == len(sortedSignal) {
			dictionary[sortedSignal] = 7
		}

		if 2 == len(sortedSignal) {
			dictionary[sortedSignal] = 1
		}

		if 7 == len(sortedSignal) {
			dictionary[sortedSignal] = 8
		}
	}

	return dictionary
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
