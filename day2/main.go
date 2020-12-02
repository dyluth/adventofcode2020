package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/dyluth/adventofcode2020/day2/rules"
)

func main() {
	lines := readInput()
	validCount := 0
	for i := range lines {
		rule := rules.ParseLine(lines[i])
		if rule.IsValidPt2() {
			validCount++
		}
	}
	fmt.Printf("\nValid count: %v\n", validCount)
}

func readInput() []string {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	data := string(dat)
	return strings.Split(data, "\n")
}
