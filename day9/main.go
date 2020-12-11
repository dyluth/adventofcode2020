package main

import (
	"fmt"
	"strconv"

	"github.com/dyluth/adventofcode2020/day9/xmas"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	in := selectionbox.ReadInput()
	// convert to ints
	inInts := make([]int, len(in))

	for i, line := range in {
		asInt, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		inInts[i] = asInt
	}

	invalid, err := xmas.FindFirstInvalid(inInts, 25)
	if err != nil {
		fmt.Printf("oh failed.. sadface\n")
		return
	}
	fmt.Printf("First Invalid: %v\n", invalid)

	r := xmas.FindWeakness(inInts, invalid)

	sum := 0
	smallest := r[0]
	largest := r[0]
	for _, val := range r {
		if smallest > val {
			smallest = val
		}
		if largest < val {
			largest = val
		}
	}
	sum = smallest + largest
	fmt.Printf("Weakness range: %+v, sum: %v\n", r, sum)
}
