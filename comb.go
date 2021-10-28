package listing

// Combinations is generator of combinations
func Combinations(list Replacer, selectNum int, repeatable bool, buf int) (c chan Replacer) {
	c = make(chan Replacer, buf)
	index := make([]int, list.Len(), list.Len())
	for i := 0; i < list.Len(); i++ {
		index[i] = i
	}

	var combGenerator func([]int, int, int) chan []int
	if repeatable {
		combGenerator = repeatedCombinations
	} else {
		combGenerator = combinations
	}

	go func() {
		defer close(c)
		for comb := range combGenerator(index, selectNum, buf) {
			c <- list.Replace(comb)
		}
	}()

	return
}

// combinations generator is for int slice
func combinations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch {
		case selectNum == 0:
			c <- []int{}
		case selectNum == len(list):
			c <- list
		case len(list) < selectNum:
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range combinations(list[i+1:], selectNum-1, buf) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return
}

// repeatedCombination is generator for int slice
func repeatedCombinations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		if selectNum == 1 {
			for v := range list {
				c <- []int{v}
			}
			return
		}
		for i := 0; i < len(list); i++ {
			for subComb := range repeatedCombinations(list[i:], selectNum-1, buf) {
				c <- append([]int{list[i]}, subComb...)
			}
		}
	}()
	return
}
