package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if nil != err {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	sum := 0

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

		mapping := createMapping(signals)

		number := ""

		for _, output := range outputs {
			mappedOutput := mapOutput(output, mapping)
			outputNumber := convertToNumber(mappedOutput)

			number += outputNumber
		}

		n, _ := strconv.Atoi(number)
		sum += n
	}

	fmt.Printf("Sum: %d \n", sum)
}

func createMapping(signals []string) map[string]string {
	dictionary := make(map[int]string, 9)
	mapping := make(map[string]string, 9)
	signalsToCheck := make([]string, 0)

	for _, signal := range signals {
		sortedSignal := sortString(signal)

		if 2 == len(sortedSignal) {
			dictionary[1] = sortedSignal
			continue
		}

		if 4 == len(sortedSignal) {
			dictionary[4] = sortedSignal
			continue
		}

		if 3 == len(sortedSignal) {
			dictionary[7] = sortedSignal
			continue
		}

		if 7 == len(sortedSignal) {
			dictionary[8] = sortedSignal
			continue
		}

		signalsToCheck = append(signalsToCheck, signal)
	}

	mapping["a"] = diffString(dictionary[1], dictionary[7])

	for _, signal := range signalsToCheck {
		diff := diffString(dictionary[4]+mapping["a"], signal)
		if 1 == len(diff) {
			mapping["g"] = diff
		}
	}

	for _, signal := range signalsToCheck {
		diff := diffString(dictionary[1]+mapping["a"]+mapping["g"], signal)

		if 1 == len(diff) {
			mapping["d"] = diff
		}
	}

	mapping["b"] = diffString(diffString(dictionary[1], dictionary[4]), mapping["d"])

	for _, signal := range signalsToCheck {
		diff := diffString(mapping["a"]+mapping["b"]+mapping["d"]+mapping["g"], signal)

		if 1 == len(diff) {
			mapping["f"] = diff
		}
	}

	mapping["c"] = diffString(dictionary[1], mapping["f"])

	for _, signal := range signalsToCheck {
		diff := diffString(mapping["a"]+mapping["c"]+mapping["d"]+mapping["g"], signal)

		if 1 == len(diff) && mapping["f"] != diff {
			mapping["e"] = diff
		}
	}

	return mapping
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func diffString(str1 string, str2 string) string {
	strArr1 := strings.Split(str1, "")
	strArr2 := strings.Split(str2, "")

	diff := ""

	for _, s1 := range strArr1 {
		contains := false
		for _, s2 := range strArr2 {
			if s1 == s2 {
				contains = true
				break
			}
		}

		if !contains {
			diff += s1
		}
	}

	for _, s2 := range strArr2 {
		contains := false
		for _, s1 := range strArr1 {
			if s2 == s1 {
				contains = true
				break
			}
		}

		if !contains {
			diff += s2
		}
	}

	return diff
}

func mapOutput(output string, mapping map[string]string) string {
	mapped := ""

	for _, str := range strings.Split(output, "") {
		for k, m := range mapping {
			if str == m {
				mapped += k
				break
			}
		}
	}

	return sortString(mapped)
}

func convertToNumber(strNum string) string {
	numbers := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}

	return numbers[strNum]
}
