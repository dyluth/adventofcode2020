package treemap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testData1 = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestCountTrees(t *testing.T) {
	m := ParseTreeMap(testData1)
	trees := m.CountTrees(0, 1)
	require.Equal(t, 3, trees)
	trees = m.CountTrees(1, 1)
	require.Equal(t, 2, trees)
	trees = m.CountTrees(0, 2)
	require.Equal(t, 1, trees)
}

func TestParseTreeMap(t *testing.T) {
	m := ParseTreeMap(testData1)
	require.Equal(t, OpenGround, MapElement(m.data[0][0]))
	require.Equal(t, Tree, MapElement(m.data[0][2]))
	require.Equal(t, 11, m.mapWidth)
}

func TestElementAt(t *testing.T) {
	m := ParseTreeMap(testData1)
	require.Equal(t, OpenGround, m.ElementAt(0, 0))
	require.Equal(t, Tree, m.ElementAt(2, 0))
	require.Equal(t, Tree, m.ElementAt(13, 0))
}

func TestMultipleSlopes(t *testing.T) {
	tm := ParseTreeMap(testData1)
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
	fmt.Printf("trees: %v\n", treeCounts)
	require.Equal(t, 336, total)
}
