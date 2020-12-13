package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2020/day10/jolts"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {

	in := selectionbox.ConvertStringsToInts(selectionbox.ReadInput())

	chain := jolts.BuildMinimumJoltChain(in)
	steps := jolts.CountChainSteps(chain)
	options := jolts.CountOptions(chain)
	fmt.Printf("steps of 1: %v steps of 3: %v\n multiplied: %v, Options: %v\n", steps[1], steps[3], steps[1]*steps[3], options)
}
