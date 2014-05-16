package listing

type Replacer interface {
	Len() int
	Replace([]int) Replacer
}

type IntReplacer []int

func (il IntReplacer) Len() int {
	return len(il)
}

func (il IntReplacer) Replace(indices []int) Replacer {
	result := make(IntReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = il[idx]
	}
	return result
}

type RuneReplacer []rune

func (rl RuneReplacer) Len() int {
	return len(rl)
}

func (rl RuneReplacer) Replace(indices []int) Replacer {
	result := make(RuneReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = rl[idx]
	}
	return result
}

type StringReplacer []string

func (sl StringReplacer) Len() int {
	return len(sl)
}
func (sl StringReplacer) Replace(indices []int) Replacer {
	result := make(StringReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = sl[idx]
	}
	return result
}

type Float64Replacer []string

func (fr Float64Replacer) Len() int {
	return len(fr)
}
func (fr Float64Replacer) Replace(indices []int) Replacer {
	result := make(Float64Replacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = fr[idx]
	}
	return result
}
