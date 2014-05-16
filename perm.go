package listing

//Permutation generator
func Permutations(list Replacer, select_num int, repeatable bool, buf int) (c chan Replacer) {
	c = make(chan Replacer, buf)
	go func() {
		defer close(c)
		var perm_generator func([]int, int, int) chan []int
		if repeatable {
			perm_generator = repeated_permutations
		} else {
			perm_generator = permutations
		}
		indices := make([]int, list.Len(), list.Len())
		for i := 0; i < list.Len(); i++ {
			indices[i] = i
		}
		for perm := range perm_generator(indices, select_num, buf) {
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
func permutations(list []int, select_num, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch select_num {
		case 1:
			for _, v := range list {
				c <- []int{v}
			}
			return
		case 0:
			return
		case len(list):
			for i := 0; i < len(list); i++ {
				top, sub_list := pop(list, i)
				for perm := range permutations(sub_list, select_num-1, buf) {
					c <- append([]int{top}, perm...)
				}
			}
		default:
			for comb := range combinations(list, select_num, buf) {
				for perm := range permutations(comb, select_num, buf) {
					c <- perm
				}
			}
		}
	}()
	return
}

//Repeated permtation generator for int slice
func repeated_permutations(list []int, select_num, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch select_num {
		case 1:
			for _, v := range list {
				c <- []int{v}
			}
		default:
			for i := 0; i < len(list); i++ {
				for perm := range repeated_permutations(list, select_num-1, buf) {
					c <- append([]int{list[i]}, perm...)
				}
			}
		}
	}()
	return
}
