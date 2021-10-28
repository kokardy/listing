package listing

// Replacer is interface for combinations and permutations
type Replacer interface {
	Len() int
	Replace([]int) Replacer
}

// IntReplacer is Replacer for int
type IntReplacer []int

// Len implements Replacer.Len
func (il IntReplacer) Len() int {
	return len(il)
}

// Replace implements Replacer.Replace
func (il IntReplacer) Replace(indices []int) Replacer {
	result := make(IntReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = il[idx]
	}
	return result
}

// RuneReplacer is Replacer for rune
type RuneReplacer []rune

// Len implements Replacer.Len
func (rl RuneReplacer) Len() int {
	return len(rl)
}

// Replace implements Replacer.Replace
func (rl RuneReplacer) Replace(indices []int) Replacer {
	result := make(RuneReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = rl[idx]
	}
	return result
}

// StringReplacer is Replacer for String
type StringReplacer []string

// Len implements Replacer.Len
func (sl StringReplacer) Len() int {
	return len(sl)
}

// Replace implements Replacer.Replace
func (sl StringReplacer) Replace(indices []int) Replacer {
	result := make(StringReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = sl[idx]
	}
	return result
}

// Float64Replacer is Replacer for Float64
type Float64Replacer []string

// Len implements Replacer.Len
func (fr Float64Replacer) Len() int {
	return len(fr)
}

// Replace implements Replacer.Replace
func (fr Float64Replacer) Replace(indices []int) Replacer {
	result := make(Float64Replacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = fr[idx]
	}
	return result
}
