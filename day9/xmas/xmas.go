package xmas

import "fmt"

func IsValid(previous []int, current int) bool {
	// put all the values in a map so we can find them later
	vals := make(map[int]bool)
	for _, v := range previous {
		vals[v] = true
	}

	for _, v := range previous {
		// if the diff between v and the current is in the map, then this is valid
		diff := current - v
		if diff == v { // 2 numbers must be different
			continue
		}
		_, ok := vals[diff]
		if ok {
			fmt.Printf(" %v+%v\n", v, diff)
			return true
		}
	}
	fmt.Printf(" NO MATCH\n")
	return false
}

func FindFirstInvalid(input []int, preambleSize int) (int, error) {

	for i := preambleSize; i < len(input); i++ {
		slice := input[i-(preambleSize) : i]
		fmt.Printf("%v (%v): %+v: ", i, input[i], slice)
		if !IsValid(slice, input[i]) {
			return input[i], nil
		}
	}
	return 0, fmt.Errorf("not found any")
}

func FindWeakness(input []int, sum int) []int {

	for i := range input {
		total := input[i]
		j := i
		notDone := true
		for notDone {
			j++
			total += input[j]
			if total == sum {
				// found the right answer
				return input[i : j+1]
			}
			if total > sum {
				notDone = false
			}
		}
	}
	return []int{}

}
