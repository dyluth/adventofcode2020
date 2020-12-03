package treemap

// TreeMap represents a map of trees over ground
type TreeMap struct {
	//for this map, first index is the depth(y), second index is the width (x)
	data     [][]rune
	mapWidth int
}

// ElementAt returns the element at a specific point on the map
func (tm *TreeMap) ElementAt(x, y int) MapElement {
	// Modulus the x as the pattern repeats to the right many times
	x = x % tm.mapWidth
	e := tm.data[y][x]
	return MapElement(e)
}

func (tm *TreeMap) CountTrees(xStep, yStep int) int {
	treecount := 0

	for y := 0; y < len(tm.data); y = y + yStep {
		element := tm.ElementAt((y/yStep)*xStep, y)
		if element == Tree {
			treecount++
		}
	}
	return treecount
}

type MapElement rune

var (
	OpenGround = MapElement('.')
	Tree       = MapElement('#')
)

func ParseTreeMap(input []string) TreeMap {
	tm := TreeMap{
		data: [][]rune{},
	}
	//for each line
	//convert to byte array
	// add to map
	for _, line := range input {
		tm.data = append(tm.data, []rune(line))
	}
	tm.mapWidth = len(tm.data[0])
	return tm
}
