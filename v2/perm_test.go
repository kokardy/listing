package listing_test

import (
	"iter"
	"reflect"
	"slices"
	"sync"
	"testing"

	"github.com/kokardy/listing/v2"
)

func assertPermSliceSeqEqual[S ~[]E, E comparable](t *testing.T, got, want []S) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len mismatch: got %d (%v), want %d (%v)", len(got), got, len(want), want)
	}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("element %d mismatch: got %v, want %v", i, got[i], want[i])
		}
	}
}

func assertStrictLexOrder(t *testing.T, seq [][]int) {
	t.Helper()
	for i := 1; i < len(seq); i++ {
		if slices.Compare(seq[i-1], seq[i]) >= 0 {
			t.Fatalf("not strictly lexicographical at index %d: %v >= %v", i, seq[i-1], seq[i])
		}
	}
}

func TestPermutationsExact(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		items []int
		k     int
		want  [][]int
	}{
		{
			name:  "n4_k3",
			items: []int{1, 2, 3, 4},
			k:     3,
			want: [][]int{
				{1, 2, 3}, {1, 2, 4}, {1, 3, 2}, {1, 3, 4}, {1, 4, 2}, {1, 4, 3},
				{2, 1, 3}, {2, 1, 4}, {2, 3, 1}, {2, 3, 4}, {2, 4, 1}, {2, 4, 3},
				{3, 1, 2}, {3, 1, 4}, {3, 2, 1}, {3, 2, 4}, {3, 4, 1}, {3, 4, 2},
				{4, 1, 2}, {4, 1, 3}, {4, 2, 1}, {4, 2, 3}, {4, 3, 1}, {4, 3, 2},
			},
		},
		{
			name:  "n3_k2",
			items: []int{1, 2, 3},
			k:     2,
			want: [][]int{
				{1, 2}, {1, 3},
				{2, 1}, {2, 3},
				{3, 1}, {3, 2},
			},
		},
		{
			name:  "n3_k3",
			items: []int{1, 2, 3},
			k:     3,
			want: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name:  "n4_k1",
			items: []int{1, 2, 3, 4},
			k:     1,
			want: [][]int{
				{1}, {2}, {3}, {4},
			},
		},
		{
			name:  "n1_k1",
			items: []int{9},
			k:     1,
			want: [][]int{
				{9},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := slices.Collect(listing.Permutations(tc.items, tc.k))
			assertPermSliceSeqEqual(t, got, tc.want)
			assertStrictLexOrder(t, got)
		})
	}
}

func TestPermutationsWithReplacementExact(t *testing.T) {
	t.Parallel()

	n4k3Want := make([][]int, 0, 64)
	for a := 1; a <= 4; a++ {
		for b := 1; b <= 4; b++ {
			for c := 1; c <= 4; c++ {
				n4k3Want = append(n4k3Want, []int{a, b, c})
			}
		}
	}

	tests := []struct {
		name  string
		items []int
		k     int
		want  [][]int
	}{
		{
			name:  "n3_k2",
			items: []int{1, 2, 3},
			k:     2,
			want: [][]int{
				{1, 1}, {1, 2}, {1, 3},
				{2, 1}, {2, 2}, {2, 3},
				{3, 1}, {3, 2}, {3, 3},
			},
		},
		{
			name:  "n2_k3_k_gt_n",
			items: []int{1, 2},
			k:     3,
			want: [][]int{
				{1, 1, 1}, {1, 1, 2}, {1, 2, 1}, {1, 2, 2},
				{2, 1, 1}, {2, 1, 2}, {2, 2, 1}, {2, 2, 2},
			},
		},
		{
			name:  "n4_k3",
			items: []int{1, 2, 3, 4},
			k:     3,
			want:  n4k3Want,
		},
		{
			name:  "n1_k3",
			items: []int{5},
			k:     3,
			want: [][]int{
				{5, 5, 5},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := slices.Collect(listing.PermutationsWithReplacement(tc.items, tc.k))
			assertPermSliceSeqEqual(t, got, tc.want)
			assertStrictLexOrder(t, got)
		})
	}
}

func TestPermutationsBoundary(t *testing.T) {
	t.Parallel()

	type boundaryCase struct {
		name      string
		items     []int
		k         int
		wantCount int
		wantFirst []int
		wantLast  []int
	}

	permCases := []boundaryCase{
		{name: "k_neg1", items: []int{1, 2, 3}, k: -1, wantCount: 0},
		{name: "k_0", items: []int{1, 2, 3}, k: 0, wantCount: 1, wantFirst: []int{}, wantLast: []int{}},
		{name: "k_1", items: []int{1, 2, 3}, k: 1, wantCount: 3, wantFirst: []int{1}, wantLast: []int{3}},
		{name: "k_n", items: []int{1, 2, 3}, k: 3, wantCount: 6, wantFirst: []int{1, 2, 3}, wantLast: []int{3, 2, 1}},
		{name: "k_n_plus_1", items: []int{1, 2, 3}, k: 4, wantCount: 0},
		{name: "nil_k_neg1", items: nil, k: -1, wantCount: 0},
		{name: "nil_k_0", items: nil, k: 0, wantCount: 1, wantFirst: []int{}, wantLast: []int{}},
		{name: "nil_k_1", items: nil, k: 1, wantCount: 0},
		{name: "empty_k_neg1", items: []int{}, k: -1, wantCount: 0},
		{name: "empty_k_0", items: []int{}, k: 0, wantCount: 1, wantFirst: []int{}, wantLast: []int{}},
		{name: "empty_k_1", items: []int{}, k: 1, wantCount: 0},
	}

	for _, tc := range permCases {
		t.Run("Permutations/"+tc.name, func(t *testing.T) {
			t.Parallel()
			got := slices.Collect(listing.Permutations(tc.items, tc.k))
			if len(got) != tc.wantCount {
				t.Fatalf("len = %d, want %d (%v)", len(got), tc.wantCount, got)
			}
			if tc.wantCount > 0 {
				if got[0] == nil {
					t.Fatal("expected non-nil slice element")
				}
				if !slices.Equal(got[0], tc.wantFirst) {
					t.Fatalf("first = %v, want %v", got[0], tc.wantFirst)
				}
				if !slices.Equal(got[len(got)-1], tc.wantLast) {
					t.Fatalf("last = %v, want %v", got[len(got)-1], tc.wantLast)
				}
			}
		})
	}

	permRepCases := []boundaryCase{
		{name: "k_neg1", items: []int{1, 2, 3}, k: -1, wantCount: 0},
		{name: "k_0", items: []int{1, 2, 3}, k: 0, wantCount: 1, wantFirst: []int{}, wantLast: []int{}},
		{name: "k_1", items: []int{1, 2, 3}, k: 1, wantCount: 3, wantFirst: []int{1}, wantLast: []int{3}},
		{name: "k_n", items: []int{1, 2, 3}, k: 3, wantCount: 27, wantFirst: []int{1, 1, 1}, wantLast: []int{3, 3, 3}},
		{name: "k_n_plus_1", items: []int{1, 2, 3}, k: 4, wantCount: 81, wantFirst: []int{1, 1, 1, 1}, wantLast: []int{3, 3, 3, 3}},
		{name: "nil_k_neg1", items: nil, k: -1, wantCount: 0},
		{name: "nil_k_0", items: nil, k: 0, wantCount: 1, wantFirst: []int{}, wantLast: []int{}},
		{name: "nil_k_1", items: nil, k: 1, wantCount: 0},
		{name: "empty_k_neg1", items: []int{}, k: -1, wantCount: 0},
		{name: "empty_k_0", items: []int{}, k: 0, wantCount: 1, wantFirst: []int{}, wantLast: []int{}},
		{name: "empty_k_2", items: []int{}, k: 2, wantCount: 0},
	}

	for _, tc := range permRepCases {
		t.Run("PermutationsWithReplacement/"+tc.name, func(t *testing.T) {
			t.Parallel()
			got := slices.Collect(listing.PermutationsWithReplacement(tc.items, tc.k))
			if len(got) != tc.wantCount {
				t.Fatalf("len = %d, want %d (%v)", len(got), tc.wantCount, got)
			}
			if tc.wantCount > 0 {
				if got[0] == nil {
					t.Fatal("expected non-nil slice element")
				}
				if !slices.Equal(got[0], tc.wantFirst) {
					t.Fatalf("first = %v, want %v", got[0], tc.wantFirst)
				}
				if !slices.Equal(got[len(got)-1], tc.wantLast) {
					t.Fatalf("last = %v, want %v", got[len(got)-1], tc.wantLast)
				}
			}
		})
	}
}

func TestPermutationsEarlyBreakAndPull(t *testing.T) {
	t.Parallel()

	targets := []struct {
		name string
		seq  iter.Seq[[]int]
		want [][]int
	}{
		{
			name: "Permutations",
			seq:  listing.Permutations([]int{1, 2, 3, 4}, 3),
			want: [][]int{{1, 2, 3}, {1, 2, 4}},
		},
		{
			name: "PermutationsWithReplacement",
			seq:  listing.PermutationsWithReplacement([]int{1, 2, 3}, 2),
			want: [][]int{{1, 1}, {1, 2}},
		},
	}

	for _, tc := range targets {
		t.Run(tc.name+"/ForRangeBreak", func(t *testing.T) {
			t.Parallel()
			var got [][]int
			for p := range tc.seq {
				got = append(got, p)
				if len(got) == 2 {
					break
				}
			}
			assertPermSliceSeqEqual(t, got, tc.want)
		})

		t.Run(tc.name+"/PullEarlyStop", func(t *testing.T) {
			t.Parallel()
			next, stop := iter.Pull(tc.seq)
			defer stop()

			var got [][]int
			for range 2 {
				v, ok := next()
				if !ok {
					t.Fatal("expected value before stop")
				}
				got = append(got, v)
			}
			stop()
			if v, ok := next(); ok {
				t.Fatalf("expected stopped iterator to yield ok=false, got %v", v)
			}
			assertPermSliceSeqEqual(t, got, tc.want)
		})

		t.Run(tc.name+"/PullImmediateStop", func(t *testing.T) {
			t.Parallel()
			next, stop := iter.Pull(tc.seq)
			stop()
			if v, ok := next(); ok {
				t.Fatalf("expected immediately stopped iterator to yield ok=false, got %v", v)
			}
		})
	}
}

func TestPermutationsSliceIndependence(t *testing.T) {
	t.Parallel()

	targets := []struct {
		name string
		seq  iter.Seq[[]int]
		want [][]int
	}{
		{
			name: "Permutations",
			seq:  listing.Permutations([]int{1, 2, 3}, 2),
			want: [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}},
		},
		{
			name: "PermutationsWithReplacement",
			seq:  listing.PermutationsWithReplacement([]int{1, 2, 3}, 2),
			want: [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}},
		},
	}

	for _, tc := range targets {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			collected := slices.Collect(tc.seq)
			assertPermSliceSeqEqual(t, collected, tc.want)

			collected[0][0] = -999
			if slices.Equal(collected[1], collected[0]) || collected[1][0] == -999 {
				t.Fatalf("mutating collected[0] corrupted collected[1]: %v", collected[1])
			}

			var seen [][]int
			for p := range tc.seq {
				seen = append(seen, slices.Clone(p))
				for i := range p {
					p[i] = -777
				}
				_ = append(p, 888)
			}
			assertPermSliceSeqEqual(t, seen, tc.want)
			assertPermSliceSeqEqual(t, slices.Collect(tc.seq), tc.want)
		})
	}
}

func TestPermutationsReuseAndConcurrent(t *testing.T) {
	t.Parallel()

	targets := []struct {
		name string
		seq  iter.Seq[[]int]
		want [][]int
	}{
		{
			name: "Permutations",
			seq:  listing.Permutations([]int{1, 2, 3}, 2),
			want: [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}},
		},
		{
			name: "PermutationsWithReplacement",
			seq:  listing.PermutationsWithReplacement([]int{1, 2, 3}, 2),
			want: [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {3, 2}, {3, 3}},
		},
	}

	for _, tc := range targets {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			first := slices.Collect(tc.seq)
			second := slices.Collect(tc.seq)
			assertPermSliceSeqEqual(t, first, tc.want)
			assertPermSliceSeqEqual(t, second, tc.want)

			const workers = 8
			results := make([][][]int, workers)
			var wg sync.WaitGroup
			wg.Add(workers)
			for w := range workers {
				go func(idx int) {
					defer wg.Done()
					results[idx] = slices.Collect(tc.seq)
				}(w)
			}
			wg.Wait()

			for w := range workers {
				assertPermSliceSeqEqual(t, results[w], tc.want)
			}
		})
	}
}

func TestPermutationsTypes(t *testing.T) {
	t.Parallel()

	type customItem struct {
		ID   int
		Name string
	}
	type customSlice []customItem

	t.Run("string", func(t *testing.T) {
		t.Parallel()
		got := slices.Collect(listing.Permutations([]string{"a", "b", "c"}, 2))
		want := [][]string{{"a", "b"}, {"a", "c"}, {"b", "a"}, {"b", "c"}, {"c", "a"}, {"c", "b"}}
		assertPermSliceSeqEqual(t, got, want)

		gotRep := slices.Collect(listing.PermutationsWithReplacement([]string{"x", "y"}, 2))
		wantRep := [][]string{{"x", "x"}, {"x", "y"}, {"y", "x"}, {"y", "y"}}
		assertPermSliceSeqEqual(t, gotRep, wantRep)
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()
		got := slices.Collect(listing.Permutations([]float64{1.5, 2.5}, 2))
		want := [][]float64{{1.5, 2.5}, {2.5, 1.5}}
		assertPermSliceSeqEqual(t, got, want)

		gotRep := slices.Collect(listing.PermutationsWithReplacement([]float64{1.5, 2.5}, 2))
		wantRep := [][]float64{{1.5, 1.5}, {1.5, 2.5}, {2.5, 1.5}, {2.5, 2.5}}
		assertPermSliceSeqEqual(t, gotRep, wantRep)
	})

	t.Run("customStructAndNamedSlice", func(t *testing.T) {
		t.Parallel()
		items := customSlice{
			{ID: 1, Name: "one"},
			{ID: 2, Name: "two"},
		}
		wantType := reflect.TypeOf(customSlice(nil))

		got := slices.Collect(listing.Permutations(items, 2))
		want := []customSlice{
			{{ID: 1, Name: "one"}, {ID: 2, Name: "two"}},
			{{ID: 2, Name: "two"}, {ID: 1, Name: "one"}},
		}
		assertPermSliceSeqEqual(t, got, want)
		for _, elem := range got {
			if reflect.TypeOf(elem) != wantType {
				t.Fatalf("runtime type = %v, want %v", reflect.TypeOf(elem), wantType)
			}
			if _, ok := any(elem).(customSlice); !ok {
				t.Fatalf("type assertion to customSlice failed for %T", elem)
			}
		}

		gotZero := slices.Collect(listing.Permutations(items, 0))
		if len(gotZero) != 1 || reflect.TypeOf(gotZero[0]) != wantType || gotZero[0] == nil {
			t.Fatalf("k=0 yielded %#v (%T), want non-nil empty customSlice", gotZero, gotZero)
		}

		gotRep := slices.Collect(listing.PermutationsWithReplacement(items, 2))
		wantRep := []customSlice{
			{{ID: 1, Name: "one"}, {ID: 1, Name: "one"}},
			{{ID: 1, Name: "one"}, {ID: 2, Name: "two"}},
			{{ID: 2, Name: "two"}, {ID: 1, Name: "one"}},
			{{ID: 2, Name: "two"}, {ID: 2, Name: "two"}},
		}
		assertPermSliceSeqEqual(t, gotRep, wantRep)
		for _, elem := range gotRep {
			if reflect.TypeOf(elem) != wantType {
				t.Fatalf("runtime type = %v, want %v", reflect.TypeOf(elem), wantType)
			}
		}
	})
}
