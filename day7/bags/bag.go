package bags

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Name        string
	Contains    []BagCount
	ContainedBy []string
}

type BagCount struct {
	Name  string
	Count int
}

type BagRules struct {
	rules map[string]Bag
}

func NewBagRules() *BagRules {
	return &BagRules{
		rules: make(map[string]Bag),
	}
}

func (br *BagRules) updateBagRule(bag Bag) {
	// find bag - if not exist add it and return
	b, ok := br.rules[bag.Name]
	if !ok {
		br.rules[bag.Name] = bag
		for i := range bag.Contains {
			br.updateBagRule(Bag{Name: bag.Contains[i].Name, ContainedBy: []string{bag.Name}})
		}
		return
	}
	// if exist add all the contents of the 2 together then re-insert that
	// hopefully there are no circumstances where we need just the unique values
	bag.ContainedBy = append(bag.ContainedBy, b.ContainedBy...)
	bag.Contains = append(bag.Contains, b.Contains...)
	br.rules[bag.Name] = bag

}

func (br *BagRules) AddRule(s string) {
	// parse the rule
	name, contains := parseRule(s)
	bag := Bag{Name: name, Contains: contains}
	// check to see if its already therea and add this to it
	br.updateBagRule(bag)
	// for each of the Contains:
	for i := range contains {
		// check if persent, if not create empty (set Name)
		b := Bag{Name: contains[i].Name, ContainedBy: []string{name}}
		br.updateBagRule(b)
	}

}

func parseRule(s string) (name string, contains []BagCount) {

	//light red bags contain 1 bright white bag, 2 muted yellow bags.

	// split on ` bags contain `
	split := strings.Split(s, " bags contain ")
	name = split[0]
	// [0] is the name
	// rest use the following regex:
	re := regexp.MustCompile(`(\d+)\s+(\w+\s+\w+)\s+bags?[.,]`)
	matches := re.FindAllStringSubmatch(split[1], -1)

	for i := range matches {
		containsCount := matches[i][1]
		containsName := matches[i][2]
		c, err := strconv.Atoi(containsCount)
		if err != nil {
			panic(err)
		}
		contains = append(contains, BagCount{
			Name:  containsName,
			Count: c,
		})
	}
	return
}

func (br *BagRules) CountDown(bagName string) int {
	return br.countDown(bagName) - 1
}

func (br *BagRules) countDown(bagName string) int {

	bag, ok := br.rules[bagName]
	if !ok {
		panic(fmt.Sprintf("couldnt find %v\n", bagName))
	}
	count := 1 // count ourselves

	for _, b := range bag.Contains {
		count += b.Count * br.countDown(b.Name)
	}
	return count
}

func (br *BagRules) GetPossibleOuterBags(bagName string) map[string]bool {

	bag, ok := br.rules[bagName]
	if !ok {
		panic(fmt.Sprintf("couldnt find %v\n", bagName))
	}

	possible := make(map[string]bool)
	for i := range bag.ContainedBy {
		possible[bag.ContainedBy[i]] = true
		p2 := br.GetPossibleOuterBags(bag.ContainedBy[i])
		for k := range p2 {
			possible[k] = true
		}
	}
	return possible
}
