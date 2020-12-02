package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Printf("Uh oh: %v\n", err)
		return
	}

	// for each element i
	// for each element j
	// if i+j<=2020
	//  store total=i*j
	//  calc opposite
	//  lookup in map

	data := make(map[int][]int)

	for i := range input {
		opp := 2020 - input[i]
		d, ok := data[opp]
		if ok {
			fmt.Printf("results: %v, %v, %v: %v\n", input[i], d[0], d[1], input[i]*d[0]*d[1])
			return
		}
		for j := range input {
			sum := input[i] + input[j]
			if sum <= 2020 {
				data[sum] = []int{input[i], input[j]}
			}

		}

	}

}

func part1() {
	input, err := readInput()
	if err != nil {
		fmt.Printf("Uh oh: %v\n", err)
		return
	}
	// load to hash
	//calc opposite number
	// lookup in hash
	// return if found

	data := make(map[int]bool)

	for i := range input {
		data[input[i]] = true
		opp := 2020 - input[i]

		_, ok := data[opp]
		if ok {
			fmt.Printf("results: %v and %v: %v\n", input[i], opp, input[i]*opp)
			return
		}
	}

}

func readInput() ([]int, error) {
	input := []int{}
	dat, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		return nil, err
	}
	data := string(dat)

	// now split and parse
	fields := strings.Fields(data)
	for i := range fields {
		val, err := strconv.Atoi(fields[i])
		if err != nil {

			return nil, err
		}
		input = append(input, val)
	}
	return input, nil
}
