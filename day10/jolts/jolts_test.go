package jolts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildMinimumJoltChain(t *testing.T) {
	answer := []int{0, 1, 2, 4, 5, 8}
	chain := BuildMinimumJoltChain([]int{1, 2, 4, 5})
	require.ElementsMatch(t, answer, chain)
	chain = BuildMinimumJoltChain([]int{5, 2, 4, 1})
	require.ElementsMatch(t, answer, chain)
}

func TestCountChainSteps(t *testing.T) {
	chain := []int{0, 1, 2, 4, 5, 8}
	res := CountChainSteps(chain)
	require.ElementsMatch(t, []int{0, 3, 1, 1}, res)
}

func TestCountOptions(t *testing.T) {
	o := CountOptions([]int{0, 1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22})
	require.Equal(t, 8, o)

	o = CountOptions([]int{0, 3, 6, 9, 12})
	require.Equal(t, 1, o)

	o = CountOptions([]int{0, 1, 2, 3, 6})
	require.Equal(t, 4, o)

	chain := BuildMinimumJoltChain([]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3})
	o = CountOptions(chain)
	require.Equal(t, 19208, o)

}
