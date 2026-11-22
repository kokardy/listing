// Package listing provides generic, Go 1.23+ iterator-based generators for
// combinations and permutations over any slice type.
//
// All exported functions return an [iter.Seq] whose elements share the exact
// slice type of the input (including custom named slice types via the type
// constraint [~[]E]). Results are produced lazily on the caller's goroutine
// using non-recursive iterative index algorithms without spawning goroutines.
//
// # Lexicographic Order
//
// Every function yields elements in lexicographic order of input indices:
//   - [Combinations]: k-element combinations without replacement.
//   - [CombinationsWithReplacement]: k-element combinations with replacement
//     (non-decreasing input indices).
//   - [Permutations]: k-element permutations without replacement in strict
//     lexicographic index order.
//   - [PermutationsWithReplacement]: k-tuples with replacement (base-n order).
//
// # Safe Early Break
//
// Because iteration runs synchronously and checks the yield return value on
// every step, breaking out of a for-range loop or calling stop on an
// [iter.Pull] iterator immediately halts generation with zero leaked
// goroutines or resources.
//
// # Slice Independence and Reuse
//
// Each yielded slice is freshly allocated, so collecting results with
// [slices.Collect] or mutating a yielded slice during iteration never affects
// previous or subsequent elements. All mutable iteration state is scoped
// to each iteration pass, allowing a returned [iter.Seq] to be iterated
// multiple times or concurrently across goroutines (assuming the input slice
// itself is not mutated concurrently).
package listing
