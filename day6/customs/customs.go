package customs

import (
	"regexp"
	"strings"
)

type Answers struct {
	input string
}

func NewAnswers(a string) Answers {
	return Answers{input: a}
}

// GetUnique
func (a *Answers) GetUnique() int {
	unique := make(map[rune]bool)
	// remove all whitepspace from input
	re := regexp.MustCompile(`\n|\s+`)
	reduced := re.ReplaceAllString(a.input, "")
	// turn to rune array
	runes := []rune(reduced)

	//add each rune to map (removes dups)
	for _, r := range runes {
		unique[r] = true
	}
	// return length of map
	return len(unique)
}

func (a *Answers) GetConsistent() int {
	// split on lines
	old := make(map[rune]bool) // map from previous iteration of loop
	in := strings.TrimSpace(a.input)
	for i, line := range strings.Split(in, "\n") {
		line = strings.TrimSpace(line)
		unique := make(map[rune]bool) // new map
		runes := []rune(line)
		for _, r := range runes {
			//  create new map, only add in to new map if in old map
			_, inOld := old[r]
			if i == 0 || inOld { // if i==0 load all of first into map
				unique[r] = true
			}
		}
		// replace old map with new map
		old = unique

	}
	return len(old)
}
