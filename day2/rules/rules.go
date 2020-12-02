package rules

import (
	"fmt"
	"regexp"
	"strconv"
)

type Rule struct {
	minCount int
	maxCount int
	letter   string
	password string
}

// 17-19 p: pwpzpfbrcpppjppbmppp
// matches: \d+-\d+ \s: \s+
func ParseLine(line string) Rule {
	lineRE := regexp.MustCompile(`(\d+)-(\d+)\s+(\w):\s+(\w+)`)
	matches := lineRE.FindStringSubmatch(line)
	for i := range matches {
		fmt.Printf("%v: %v\n", i, matches[i])
	}
	min, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	return Rule{
		minCount: min,
		maxCount: max,
		letter:   matches[3],
		password: matches[4],
	}
}

// IsValid validate as per part 1
func (r *Rule) IsValid() bool {
	re := regexp.MustCompile(r.letter)
	matches := re.FindAllString(r.password, -1)
	count := len(matches)

	if count < r.minCount {
		fmt.Printf("%v not valid %v < %v\n", r.password, count, r.minCount)
		return false
	}
	if count > r.maxCount {
		fmt.Printf("%v not valid %v > %v\n", r.password, count, r.maxCount)
		return false
	}
	return true
}

// IsValidPt2 validate as per part 2
func (r *Rule) IsValidPt2() bool {
	// turn pw into byte array
	pwBytes := []byte(r.password)
	letterByte := []byte(r.letter)[0]
	matches := 0
	// if letter at min-1 matches password then ok!
	if pwBytes[r.minCount-1] == letterByte {
		matches++
	}
	// if letter at max-1 matches password then ok!
	if pwBytes[r.maxCount-1] == letterByte {
		matches++
	}
	if matches == 1 {
		return true
	}
	return false
}
