package selectionbox

import (
	"io/ioutil"
	"strings"
)

func ReadInput() []string {
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	data := string(dat)
	return strings.Split(data, "\n")
}
