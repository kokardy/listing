package listing

// Combinations generates combinations of list elements and sends them on c.
func Combinations(list Replacer, selectNum int, repeatable bool, buf int) (c chan Replacer) {
	if buf < 0 {
		buf = 0
	}
	c = make(chan Replacer, buf)
	if list == nil {
		close(c)
		return c
	}

	index := make([]int, list.Len())
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

	return c
}

// combinations generates non-repeating combinations of an int slice.
func combinations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch {
		case selectNum < 0 || len(list) < selectNum:
			return
		case selectNum == 0:
			c <- []int{}
		case selectNum == len(list):
			dup := make([]int, len(list))
			copy(dup, list)
			c <- dup
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range combinations(list[i+1:], selectNum-1, buf) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return c
}

// repeatedCombinations generates repeating combinations of an int slice.
func repeatedCombinations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch {
		case selectNum < 0:
			return
		case selectNum == 0:
			c <- []int{}
			return
		case len(list) == 0:
			return
		case selectNum == 1:
			for _, v := range list {
				c <- []int{v}
			}
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range repeatedCombinations(list[i:], selectNum-1, buf) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return c
}
