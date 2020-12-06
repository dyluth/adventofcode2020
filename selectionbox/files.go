package selectionbox

import (
	"io/ioutil"
	"regexp"
	"strings"
)

//ReadInput returns lines of input separated by new lines
func ReadInput() []string {
	data := read()
	return strings.Split(data, "\n")
}

//ReadGroupedInput returns groups of input separated by empty lines
func ReadGroupedInput() []string {
	data := read()
	re := regexp.MustCompile(`\n\s*\n`)
	passList := re.Split(data, -1)
	return passList
}

func read() string {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)
}
