package listing

//Combination generator
func Combinations(list Replacer, select_num int, repeatable bool, buf int) (c chan Replacer) {
	c = make(chan Replacer, buf)
	index := make([]int, list.Len(), list.Len())
	for i := 0; i < list.Len(); i++ {
		index[i] = i
	}

	var comb_generator func([]int, int, int) chan []int
	if repeatable {
		comb_generator = repeated_combinations
	} else {
		comb_generator = combinations
	}

	go func() {
		defer close(c)
		for comb := range comb_generator(index, select_num, buf) {
			c <- list.Replace(comb)
		}
	}()

	return
}

//Combination generator for int slice
func combinations(list []int, select_num, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch {
		case select_num == 0:
			c <- []int{}
		case select_num == len(list):
			c <- list
		case len(list) < select_num:
			return
		default:
			for i := 0; i < len(list); i++ {
				for sub_comb := range combinations(list[i+1:], select_num-1, buf) {
					c <- append([]int{list[i]}, sub_comb...)
				}
			}
		}
	}()
	return
}

//Repeated combination generator for int slice
func repeated_combinations(list []int, select_num, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		if select_num == 1 {
			for v := range list {
				c <- []int{v}
			}
			return
		}
		for i := 0; i < len(list); i++ {
			for sub_comb := range repeated_combinations(list[i:], select_num-1, buf) {
				c <- append([]int{list[i]}, sub_comb...)
			}
		}
	}()
	return
}
