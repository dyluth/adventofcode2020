package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2020/day3/treemap"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	input := selectionbox.ReadInput()
	tm := treemap.ParseTreeMap(input)
	treeCounts := []int{}
	treeCounts = append(treeCounts, tm.CountTrees(1, 1))
	treeCounts = append(treeCounts, tm.CountTrees(3, 1))
	treeCounts = append(treeCounts, tm.CountTrees(5, 1))
	treeCounts = append(treeCounts, tm.CountTrees(7, 1))
	treeCounts = append(treeCounts, tm.CountTrees(1, 2))

	total := 1
	for _, t := range treeCounts {
		total = total * t
	}

	fmt.Printf("trees found at angles [%v], = %v\n", treeCounts, total)

}
