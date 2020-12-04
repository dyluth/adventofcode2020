package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2020/day4/passport"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	input := selectionbox.ReadInput()

	passportStrings := passport.SplitInput(input)
	validCount := 0
	for _, p := range passportStrings {
		p, err := passport.ParsePassport(p)
		if err == nil {
			err = p.Validate()
			if err == nil {
				validCount++
			}
		}
	}
	fmt.Printf("Valid passports: %v\n", validCount)

}
