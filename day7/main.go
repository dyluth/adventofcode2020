package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2020/day7/bags"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	rules := bags.NewBagRules()

	in := selectionbox.ReadInput()
	for i := range in {
		rules.AddRule(in[i])
	}
	possible := rules.GetPossibleOuterBags("shiny gold")
	fmt.Printf("POSSIBLE: %v\n%+v\n", len(possible), possible)

	contains := rules.CountDown("shiny gold")
	fmt.Printf("conatins: %v bags\n", contains)

}
