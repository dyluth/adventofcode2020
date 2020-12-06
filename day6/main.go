package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2020/day6/customs"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	input := selectionbox.ReadGroupedInput()
	total := 0
	consistent := 0
	for _, in := range input {
		a := customs.NewAnswers(in)
		total += a.GetUnique()
		consistent += a.GetConsistent()

	}
	fmt.Printf("total count: %v\nconsistent count: %v\n", total, consistent)
}
