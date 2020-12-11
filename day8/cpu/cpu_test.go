package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	cpu := NewCPU()
	cpu.debug = true
	cpu.LoadInstruction("nop +0")
	cpu.LoadInstruction("acc +1")
	cpu.LoadInstruction("jmp +4")
	cpu.LoadInstruction("acc +3")
	cpu.LoadInstruction("jmp -3")
	cpu.LoadInstruction("acc -99")
	cpu.LoadInstruction("acc +1")
	cpu.LoadInstruction("jmp -4")
	cpu.LoadInstruction("acc +6	")

	require.Len(t, cpu.instructions, 9)

	cpu.RunTillLoopDetected()
	require.Equal(t, 5, Accumulator)
}
