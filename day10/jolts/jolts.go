package jolts

func CountChainSteps(adapters []int) []int {
	steps := []int{0, 0, 0, 0}
	for i, v := range adapters {
		if i == 0 {
			continue
		}
		diff := v - adapters[i-1]
		steps[diff] = steps[diff] + 1
	}
	return steps
}

func CountOptions(adapters []int) int {

	l := len(adapters)
	options := make(map[int]int)
	adapterMap, _ := toMap(adapters)

	// we want to satart from the second last element and work backwards
	for i := l - 2; i >= 0; i-- {
		children := findAdapterOpts(adapterMap, adapters[i])
		sum := 0
		// we have the same number of options as all our children added together
		for _, v := range children {
			childOpt, _ := options[v]
			sum = sum + childOpt
		}
		// regardless if the children have any options, we have 1 less options than the number of children we have
		sum += len(children) - 1
		options[adapters[i]] = sum
		//fmt.Printf("(%v) Adapter %v, children: %+v options %v\n", i, adapters[i], children, sum)
	}
	ZeroVal := options[0]
	return ZeroVal + 1 // we have the count of choices, but even with 0 choices, there is still 1 option, so need to increase by 1
}

func toMap(adapters []int) (map[int]bool, int) {
	largestRating := 0
	adapterMap := make(map[int]bool)
	for _, a := range adapters {
		adapterMap[a] = true
		if a > largestRating {
			largestRating = a
		}
	}
	return adapterMap, largestRating
}

func BuildMinimumJoltChain(adapters []int) (chain []int) {
	chain = []int{0} // include the initial joltage
	// chuck the adapters - in a map
	// find the largest one - this is the end point
	adapterMap, largestRating := toMap(adapters)

	joltage := 0
	for {
		// if we have got to the largest, then we are done
		if joltage == largestRating {
			// add in 3 for the last jump
			chain = append(chain, largestRating+3)
			return chain
		}
		joltage = findNextAdapter(adapterMap, joltage)
		chain = append(chain, joltage)
	}
}

func findAdapterOpts(adapters map[int]bool, joltage int) []int {
	options := []int{}
	for i := 1; i < 4; i++ {
		_, ok := adapters[i+joltage]
		if ok {
			options = append(options, i+joltage)
		}
	}
	return options
}

func findNextAdapter(adapters map[int]bool, joltage int) int {
	return findAdapterOpts(adapters, joltage)[0]
	//panic("Uh-oh! didnt find anything!")
}
