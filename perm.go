package listing

// Permutations is generator
func Permutations(list Replacer, selectNum int, repeatable bool, buf int) (c chan Replacer) {
	c = make(chan Replacer, buf)
	go func() {
		defer close(c)
		var permGenerator func([]int, int, int) chan []int
		if repeatable {
			permGenerator = repeatedPermutations
		} else {
			permGenerator = permutations
		}
		indices := make([]int, list.Len(), list.Len())
		for i := 0; i < list.Len(); i++ {
			indices[i] = i
		}
		for perm := range permGenerator(indices, selectNum, buf) {
			c <- list.Replace(perm)
		}
	}()
	return
}

func pop(l []int, i int) (v int, sl []int) {
	v = l[i]
	length := len(l)
	sl = make([]int, length-1, length-1)
	copy(sl, l[:i])
	copy(sl[i:], l[i+1:])
	return
}

//Permtation generator for int slice
func permutations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch selectNum {
		case 1:
			for _, v := range list {
				c <- []int{v}
			}
			return
		case 0:
			return
		case len(list):
			for i := 0; i < len(list); i++ {
				top, subList := pop(list, i)
				for perm := range permutations(subList, selectNum-1, buf) {
					c <- append([]int{top}, perm...)
				}
			}
		default:
			for comb := range combinations(list, selectNum, buf) {
				for perm := range permutations(comb, selectNum, buf) {
					c <- perm
				}
			}
		}
	}()
	return
}

//Repeated permtation generator for int slice
func repeatedPermutations(list []int, selectNum, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch selectNum {
		case 1:
			for _, v := range list {
				c <- []int{v}
			}
		default:
			for i := 0; i < len(list); i++ {
				for perm := range repeatedPermutations(list, selectNum-1, buf) {
					c <- append([]int{list[i]}, perm...)
				}
			}
		}
	}()
	return
}
