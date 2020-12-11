package main

import (
	"fmt"
	"strings"

	"github.com/dyluth/adventofcode2020/day8/cpu"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	in := selectionbox.ReadInput()
	for i := range in {
		acc, loop := testSwitchIndex(i, in)
		if !loop {
			fmt.Printf("NORMAL TERMINATION! ACC: %v\n", acc)
			return
		}
	}
}

// returns acc when ended and true if infinite loop detected
func testSwitchIndex(index int, in []string) (int, bool) {
	c := cpu.NewCPU()
	for i, line := range in {
		if i == index {
			oldLine := line
			if strings.Contains(line, "jmp") {
				line = strings.ReplaceAll(line, "jmp", "nop")
			} else {
				line = strings.ReplaceAll(line, "nop", "jmp")
			}
			fmt.Printf("%v - Performing switch from [%v] to [%v]\n", i, oldLine, line)
		}
		c.LoadInstruction(line)
	}

	loop := c.RunTillLoopDetected()
	//fmt.Printf("ACC: %v\n", cpu.Accumulator)

	return cpu.Accumulator, loop
}
