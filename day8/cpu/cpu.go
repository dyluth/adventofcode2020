package cpu

import (
	"fmt"
	"regexp"
	"strconv"
)

type Instruction interface {
	Execute() int // returns the relative index to execute next
	Print()
}

func NewCPU() CPU {
	Accumulator = 0
	return CPU{instructions: []Instruction{}}
}

type CPU struct {
	instructions     []Instruction
	programCounter   int
	instructionCount []int
	debug            bool
}

func (c *CPU) LoadInstruction(s string) {
	c.instructionCount = append(c.instructionCount, 0)

	re := regexp.MustCompile(`(\w+)\s+([+-])(\d+)`)
	matches := re.FindStringSubmatch(s)
	value, err := strconv.Atoi(matches[3])
	if err != nil {
		panic(err)
	}
	if matches[2] == "-" {
		value = -value
	}
	switch matches[1] {
	case "nop":
		c.instructions = append(c.instructions, &InstructionNOOP{})
	case "jmp":
		c.instructions = append(c.instructions, &InstructionJMP{Jmp: value})
	case "acc":
		c.instructions = append(c.instructions, &InstructionACC{Acc: value})
	default:
		panic(matches[1])
	}
}

// RunTillLoopDetected returns true if a loop was detected, false if it terminated on its own
func (c *CPU) RunTillLoopDetected() bool {
	for {
		c.RunNext()
		if c.debug {
			fmt.Printf("ACC: %+v Count %v next: ", Accumulator, c.instructionCount[c.programCounter])
			c.instructions[c.programCounter].Print()
		}
		if c.programCounter >= len(c.instructionCount) {
			return false // terminated on its own
		}
		// next command to run is already at 1, so would be 2 when executed
		if c.instructionCount[c.programCounter] == 1 {
			return true // loop detected
		}
	}
}

// RunNext - runs the command
func (c *CPU) RunNext() {
	c.instructionCount[c.programCounter] = c.instructionCount[c.programCounter] + 1
	newPC := c.programCounter + c.instructions[c.programCounter].Execute()
	c.programCounter = newPC
}

var Accumulator = 0

type InstructionNOOP struct{}

func (in *InstructionNOOP) Execute() (relativeIndex int) {
	return 1
}
func (in *InstructionNOOP) Print() {
	fmt.Printf("nop +0\n")
}

type InstructionJMP struct {
	Jmp int
}

func (ij *InstructionJMP) Execute() (relativeIndex int) {
	return ij.Jmp
}
func (ij *InstructionJMP) Print() {
	fmt.Printf("jmp %v\n", ij.Jmp)
}

type InstructionACC struct {
	Acc int
}

func (ia *InstructionACC) Execute() (relativeIndex int) {
	Accumulator += ia.Acc
	return 1
}
func (ia *InstructionACC) Print() {
	fmt.Printf("acc %v\n", ia.Acc)
}
