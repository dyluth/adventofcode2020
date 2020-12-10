package bags

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRule(t *testing.T) {

	name, contains := parseRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	require.Equal(t, "light red", name)
	require.Len(t, contains, 2)
	require.Equal(t, 1, contains[0].Count)
	require.Equal(t, "bright white", contains[0].Name)
	require.Equal(t, 2, contains[1].Count)
	require.Equal(t, "muted yellow", contains[1].Name)

	name, contains = parseRule("vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.")
	require.Equal(t, "vibrant plum", name)
	require.Len(t, contains, 2)
	require.Equal(t, 5, contains[0].Count)
	require.Equal(t, "faded blue", contains[0].Name)
	require.Equal(t, 6, contains[1].Count)
	require.Equal(t, "dotted black", contains[1].Name)

	name, contains = parseRule("dotted black bags contain no other bags.")
	require.Equal(t, "dotted black", name)
	require.Len(t, contains, 0)

}

func TestUpdateBagRule(t *testing.T) {
	rules := NewBagRules()

	rules.updateBagRule(Bag{Name: "one", Contains: []BagCount{{Name: "two", Count: 1}}})
	require.Len(t, rules.rules, 2)
	require.Len(t, rules.rules["two"].ContainedBy, 1)
	require.Equal(t, "one", rules.rules["two"].ContainedBy[0])

}
func TestAddRule(t *testing.T) {
	rules := NewBagRules()

	rules.AddRule("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	rules.AddRule("dark orange bags contain 3 bright white bags, 4 muted yellow bags.")
	rules.AddRule("bright white bags contain 1 shiny gold bag.")
	rules.AddRule("muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.")
	rules.AddRule("shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.")
	rules.AddRule("dark olive bags contain 3 faded blue bags, 4 dotted black bags.")
	rules.AddRule("vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.")
	rules.AddRule("faded blue bags contain no other bags.")
	rules.AddRule("dotted black bags contain no other bags.")

	require.Equal(t, 9, len(rules.rules))

	require.Len(t, rules.rules["dotted black"].Contains, 0)
	require.Len(t, rules.rules["dotted black"].ContainedBy, 2)
	require.Len(t, rules.rules["faded blue"].ContainedBy, 3)

	possible := rules.GetPossibleOuterBags("shiny gold")
	require.Len(t, possible, 4)
	fmt.Printf("POSSIBLE: %+v\n", possible)
	require.Contains(t, possible, "bright white")
	require.Contains(t, possible, "dark orange")
	require.Contains(t, possible, "light red")
	require.Contains(t, possible, "muted yellow")
}

func TestCountDown(t *testing.T) {
	rules := NewBagRules()

	rules.AddRule("shiny gold bags contain 2 dark red bags.")
	rules.AddRule("dark red bags contain 2 dark orange bags.")
	rules.AddRule("dark orange bags contain 2 dark yellow bags.")
	rules.AddRule("dark yellow bags contain 2 dark green bags.")
	rules.AddRule("dark green bags contain 2 dark blue bags.")
	rules.AddRule("dark blue bags contain 2 dark violet bags.")
	rules.AddRule("dark violet bags contain no other bags.")

	count := rules.CountDown("shiny gold")
	require.Equal(t, count, 126)
}

func TestCountDown2(t *testing.T) {
	rules := NewBagRules()

	rules.AddRule("shiny gold bags contain 2 dark red bags.")
	rules.AddRule("dark red bags contain no other bags.")

	count := rules.CountDown("shiny gold")

	require.Equal(t, 2, count)
}
