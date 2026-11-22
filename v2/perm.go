package listing

import "iter"

// Permutations returns an iterator over all k-element permutations of items
// without replacement, in strict lexicographic order of input indices.
//
// If k < 0 or k > len(items), the returned sequence yields no elements.
// If k == 0, the returned sequence yields a single empty slice make(S, 0).
// Each yielded slice is freshly allocated, and breaking early stops iteration
// immediately without leaking goroutines.
func Permutations[S ~[]E, E any](items S, k int) iter.Seq[S] {
	n := len(items)
	if seq, done := edgeCaseSeq[S](n, k, false); done {
		return seq
	}

	return func(yield func(S) bool) {
		indices := make([]int, n)
		for i := range indices {
			indices[i] = i
		}
		cycles := make([]int, k)
		for i := range cycles {
			cycles[i] = n - i
		}

		for {
			if !yield(materialize(items, indices[:k])) {
				return
			}
			advanced := false
			for i := k - 1; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					// Rotate indices[i:] left by one so the remaining suffix
					// returns to ascending index order for the next prefix.
					first := indices[i]
					copy(indices[i:], indices[i+1:])
					indices[n-1] = first
					cycles[i] = n - i
					continue
				}
				swapIdx := n - cycles[i]
				indices[i], indices[swapIdx] = indices[swapIdx], indices[i]
				advanced = true
				break
			}
			if !advanced {
				return
			}
		}
	}
}

// PermutationsWithReplacement returns an iterator over all k-tuples of items
// with replacement, in lexicographic order of input indices.
//
// If k < 0, or if len(items) == 0 and k > 0, the returned sequence yields no
// elements. If k == 0, the returned sequence yields a single empty slice
// make(S, 0). Each yielded slice is freshly allocated, and breaking early stops
// iteration immediately without leaking goroutines.
func PermutationsWithReplacement[S ~[]E, E any](items S, k int) iter.Seq[S] {
	n := len(items)
	if seq, done := edgeCaseSeq[S](n, k, true); done {
		return seq
	}

	return func(yield func(S) bool) {
		idx := make([]int, k)

		for {
			if !yield(materialize(items, idx)) {
				return
			}
			i := k - 1
			for i >= 0 && idx[i] == n-1 {
				idx[i] = 0
				i--
			}
			if i < 0 {
				return
			}
			idx[i]++
		}
	}
}
