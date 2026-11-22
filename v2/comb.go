package listing

import "iter"

// Combinations returns an iterator over all k-element combinations of items
// without replacement, in lexicographic order of input indices.
//
// If k < 0 or k > len(items), the returned sequence yields no elements.
// If k == 0, the returned sequence yields a single empty slice make(S, 0).
// Each yielded slice is freshly allocated, and breaking early stops iteration
// immediately without leaking goroutines.
func Combinations[S ~[]E, E any](items S, k int) iter.Seq[S] {
	n := len(items)
	if seq, done := edgeCaseSeq[S](n, k, false); done {
		return seq
	}

	return func(yield func(S) bool) {
		idx := make([]int, k)
		for i := range idx {
			idx[i] = i
		}

		for {
			if !yield(materialize(items, idx)) {
				return
			}
			i := k - 1
			for i >= 0 && idx[i] == i+n-k {
				i--
			}
			if i < 0 {
				return
			}
			idx[i]++
			for j := i + 1; j < k; j++ {
				idx[j] = idx[j-1] + 1
			}
		}
	}
}

// CombinationsWithReplacement returns an iterator over all k-element
// combinations of items with replacement (non-decreasing input indices), in
// lexicographic order.
//
// If k < 0, or if len(items) == 0 and k > 0, the returned sequence yields no
// elements. If k == 0, the returned sequence yields a single empty slice
// make(S, 0). Each yielded slice is freshly allocated, and breaking early stops
// iteration immediately without leaking goroutines.
func CombinationsWithReplacement[S ~[]E, E any](items S, k int) iter.Seq[S] {
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
				i--
			}
			if i < 0 {
				return
			}
			next := idx[i] + 1
			for j := i; j < k; j++ {
				idx[j] = next
			}
		}
	}
}

func edgeCaseSeq[S ~[]E, E any](n, k int, withReplacement bool) (iter.Seq[S], bool) {
	if k < 0 {
		return emptySeq[S](), true
	}
	if k == 0 {
		return singleEmptySeq[S](), true
	}
	if n == 0 || (!withReplacement && k > n) {
		return emptySeq[S](), true
	}
	return nil, false
}

func emptySeq[S ~[]E, E any]() iter.Seq[S] {
	return func(func(S) bool) {}
}

func singleEmptySeq[S ~[]E, E any]() iter.Seq[S] {
	return func(yield func(S) bool) {
		if !yield(make(S, 0)) {
			return
		}
	}
}

func materialize[S ~[]E, E any](items S, idx []int) S {
	out := make(S, len(idx))
	for i, p := range idx {
		out[i] = items[p]
	}
	return out
}
