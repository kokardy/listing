package listing_test

import (
	"iter"
	"testing"

	"github.com/kokardy/listing/v2"
)

var benchSink int

func runSeqBench(b *testing.B, seq iter.Seq[[]int], limit int) {
	b.Helper()
	b.ReportAllocs()
	for range b.N {
		count := 0
		for v := range seq {
			benchSink += v[0]
			count++
			if limit > 0 && count >= limit {
				break
			}
		}
	}
}

func BenchmarkCombinations(b *testing.B) {
	items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.Run("Full_n10_k4", func(b *testing.B) {
		runSeqBench(b, listing.Combinations(items, 4), 0)
	})
	b.Run("EarlyBreak10_n10_k4", func(b *testing.B) {
		runSeqBench(b, listing.Combinations(items, 4), 10)
	})
}

func BenchmarkCombinationsWithReplacement(b *testing.B) {
	items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.Run("Full_n10_k4", func(b *testing.B) {
		runSeqBench(b, listing.CombinationsWithReplacement(items, 4), 0)
	})
	b.Run("EarlyBreak10_n10_k4", func(b *testing.B) {
		runSeqBench(b, listing.CombinationsWithReplacement(items, 4), 10)
	})
}

func BenchmarkPermutations(b *testing.B) {
	items := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b.Run("Full_n8_k4", func(b *testing.B) {
		runSeqBench(b, listing.Permutations(items, 4), 0)
	})
	b.Run("EarlyBreak10_n8_k4", func(b *testing.B) {
		runSeqBench(b, listing.Permutations(items, 4), 10)
	})
}

func BenchmarkPermutationsWithReplacement(b *testing.B) {
	items := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b.Run("Full_n8_k4", func(b *testing.B) {
		runSeqBench(b, listing.PermutationsWithReplacement(items, 4), 0)
	})
	b.Run("EarlyBreak10_n8_k4", func(b *testing.B) {
		runSeqBench(b, listing.PermutationsWithReplacement(items, 4), 10)
	})
}
