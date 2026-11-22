package listing

import (
	"reflect"
	"testing"
	"time"
)

type regCustomItem struct {
	ID   int
	Name string
}

type regCustomReplacer []regCustomItem

func (r regCustomReplacer) Len() int {
	return len(r)
}

func (r regCustomReplacer) Replace(indices []int) Replacer {
	out := make(regCustomReplacer, len(indices))
	for i, idx := range indices {
		out[i] = r[idx]
	}
	return out
}

func collectStrings(ch chan Replacer) [][]string {
	var got [][]string
	for item := range ch {
		sr := item.(StringReplacer)
		cp := make([]string, len(sr))
		copy(cp, sr)
		got = append(got, cp)
	}
	return got
}

func TestRegressionCombinationsNonRepeatable(t *testing.T) {
	t.Run("README_example", func(t *testing.T) {
		list := StringReplacer{"A", "B", "C", "D"}
		want := [][]string{
			{"A", "B", "C"},
			{"A", "B", "D"},
			{"A", "C", "D"},
			{"B", "C", "D"},
		}
		got := collectStrings(Combinations(list, 3, false, 5))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("edge_cases", func(t *testing.T) {
		cases := []struct {
			name      string
			list      StringReplacer
			selectNum int
			want      [][]string
		}{
			{
				name:      "selectNum=0",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 0,
				want:      [][]string{{}},
			},
			{
				name:      "selectNum=1",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 1,
				want:      [][]string{{"A"}, {"B"}, {"C"}},
			},
			{
				name:      "selectNum=len",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 3,
				want:      [][]string{{"A", "B", "C"}},
			},
			{
				name:      "selectNum_gt_len",
				list:      StringReplacer{"A", "B"},
				selectNum: 3,
				want:      nil,
			},
			{
				name:      "empty_list_selectNum=1",
				list:      StringReplacer{},
				selectNum: 1,
				want:      nil,
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				got := collectStrings(Combinations(tc.list, tc.selectNum, false, 2))
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("got %v, want %v", got, tc.want)
				}
			})
		}
	})

	t.Run("IntReplacer", func(t *testing.T) {
		list := IntReplacer{10, 20, 30}
		want := [][]int{{10, 20}, {10, 30}, {20, 30}}
		var got [][]int
		for item := range Combinations(list, 2, false, 2) {
			ir := item.(IntReplacer)
			cp := make([]int, len(ir))
			copy(cp, ir)
			got = append(got, cp)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("RuneReplacer", func(t *testing.T) {
		list := RuneReplacer{'x', 'y', 'z'}
		want := [][]rune{{'x', 'y'}, {'x', 'z'}, {'y', 'z'}}
		var got [][]rune
		for item := range Combinations(list, 2, false, 2) {
			rr := item.(RuneReplacer)
			cp := make([]rune, len(rr))
			copy(cp, rr)
			got = append(got, cp)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("CustomReplacer", func(t *testing.T) {
		list := regCustomReplacer{
			{ID: 1, Name: "one"},
			{ID: 2, Name: "two"},
			{ID: 3, Name: "three"},
		}
		want := []regCustomReplacer{
			{{ID: 1, Name: "one"}, {ID: 2, Name: "two"}},
			{{ID: 1, Name: "one"}, {ID: 3, Name: "three"}},
			{{ID: 2, Name: "two"}, {ID: 3, Name: "three"}},
		}
		var got []regCustomReplacer
		for item := range Combinations(list, 2, false, 2) {
			got = append(got, item.(regCustomReplacer))
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestRegressionPermutationsNonRepeatable(t *testing.T) {
	t.Run("README_example", func(t *testing.T) {
		list := StringReplacer{"A", "B", "C", "D"}
		want := [][]string{
			{"A", "B", "C"},
			{"A", "C", "B"},
			{"B", "A", "C"},
			{"B", "C", "A"},
			{"C", "A", "B"},
			{"C", "B", "A"},
			{"A", "B", "D"},
			{"A", "D", "B"},
			{"B", "A", "D"},
			{"B", "D", "A"},
			{"D", "A", "B"},
			{"D", "B", "A"},
			{"A", "C", "D"},
			{"A", "D", "C"},
			{"C", "A", "D"},
			{"C", "D", "A"},
			{"D", "A", "C"},
			{"D", "C", "A"},
			{"B", "C", "D"},
			{"B", "D", "C"},
			{"C", "B", "D"},
			{"C", "D", "B"},
			{"D", "B", "C"},
			{"D", "C", "B"},
		}
		got := collectStrings(Permutations(list, 3, false, 5))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("edge_cases", func(t *testing.T) {
		cases := []struct {
			name      string
			list      StringReplacer
			selectNum int
			want      [][]string
		}{
			{
				name:      "selectNum=0",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 0,
				want:      nil,
			},
			{
				name:      "selectNum=1",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 1,
				want:      [][]string{{"A"}, {"B"}, {"C"}},
			},
			{
				name:      "selectNum=len",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 3,
				want: [][]string{
					{"A", "B", "C"},
					{"A", "C", "B"},
					{"B", "A", "C"},
					{"B", "C", "A"},
					{"C", "A", "B"},
					{"C", "B", "A"},
				},
			},
			{
				name:      "selectNum_gt_len",
				list:      StringReplacer{"A", "B"},
				selectNum: 3,
				want:      nil,
			},
			{
				name:      "empty_list_selectNum=1",
				list:      StringReplacer{},
				selectNum: 1,
				want:      nil,
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				got := collectStrings(Permutations(tc.list, tc.selectNum, false, 2))
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("got %v, want %v", got, tc.want)
				}
			})
		}
	})

	t.Run("IntReplacer", func(t *testing.T) {
		list := IntReplacer{10, 20, 30}
		want := [][]int{
			{10, 20},
			{20, 10},
			{10, 30},
			{30, 10},
			{20, 30},
			{30, 20},
		}
		var got [][]int
		for item := range Permutations(list, 2, false, 2) {
			ir := item.(IntReplacer)
			cp := make([]int, len(ir))
			copy(cp, ir)
			got = append(got, cp)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("RuneReplacer", func(t *testing.T) {
		list := RuneReplacer{'x', 'y'}
		want := [][]rune{
			{'x', 'y'},
			{'y', 'x'},
		}
		var got [][]rune
		for item := range Permutations(list, 2, false, 2) {
			rr := item.(RuneReplacer)
			cp := make([]rune, len(rr))
			copy(cp, rr)
			got = append(got, cp)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestRegressionPermutationsRepeatable(t *testing.T) {
	t.Run("ABC_select2", func(t *testing.T) {
		list := StringReplacer{"A", "B", "C"}
		want := [][]string{
			{"A", "A"},
			{"A", "B"},
			{"A", "C"},
			{"B", "A"},
			{"B", "B"},
			{"B", "C"},
			{"C", "A"},
			{"C", "B"},
			{"C", "C"},
		}
		got := collectStrings(Permutations(list, 2, true, 3))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("selectNum=1", func(t *testing.T) {
		list := StringReplacer{"A", "B", "C"}
		want := [][]string{{"A"}, {"B"}, {"C"}}
		got := collectStrings(Permutations(list, 1, true, 1))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("selectNum_gt_len", func(t *testing.T) {
		list := StringReplacer{"A", "B"}
		want := [][]string{
			{"A", "A", "A"},
			{"A", "A", "B"},
			{"A", "B", "A"},
			{"A", "B", "B"},
			{"B", "A", "A"},
			{"B", "A", "B"},
			{"B", "B", "A"},
			{"B", "B", "B"},
		}
		got := collectStrings(Permutations(list, 3, true, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func collectFloat64s(ch chan Replacer) [][]float64 {
	var got [][]float64
	for item := range ch {
		fr := item.(Float64Replacer)
		cp := make([]float64, len(fr))
		copy(cp, fr)
		got = append(got, cp)
	}
	return got
}

func TestBugFixCombinationsRepeatable(t *testing.T) {
	cases := []struct {
		name      string
		list      StringReplacer
		selectNum int
		buf       int
		want      [][]string
	}{
		{
			name:      "ABC_select2",
			list:      StringReplacer{"A", "B", "C"},
			selectNum: 2,
			buf:       2,
			want: [][]string{
				{"A", "A"},
				{"A", "B"},
				{"A", "C"},
				{"B", "B"},
				{"B", "C"},
				{"C", "C"},
			},
		},
		{
			name:      "AB_select3",
			list:      StringReplacer{"A", "B"},
			selectNum: 3,
			buf:       2,
			want: [][]string{
				{"A", "A", "A"},
				{"A", "A", "B"},
				{"A", "B", "B"},
				{"B", "B", "B"},
			},
		},
		{
			name:      "ABC_select1",
			list:      StringReplacer{"A", "B", "C"},
			selectNum: 1,
			buf:       1,
			want: [][]string{
				{"A"},
				{"B"},
				{"C"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collectStrings(Combinations(tc.list, tc.selectNum, true, tc.buf))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestBugFixFloat64Replacer(t *testing.T) {
	list := Float64Replacer{1.5, 2.5, 3.5}

	t.Run("Combinations_non_repeatable", func(t *testing.T) {
		want := [][]float64{{1.5, 2.5}, {1.5, 3.5}, {2.5, 3.5}}
		got := collectFloat64s(Combinations(list, 2, false, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("Combinations_repeatable", func(t *testing.T) {
		want := [][]float64{
			{1.5, 1.5},
			{1.5, 2.5},
			{1.5, 3.5},
			{2.5, 2.5},
			{2.5, 3.5},
			{3.5, 3.5},
		}
		got := collectFloat64s(Combinations(list, 2, true, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("Permutations_non_repeatable", func(t *testing.T) {
		want := [][]float64{
			{1.5, 2.5},
			{2.5, 1.5},
			{1.5, 3.5},
			{3.5, 1.5},
			{2.5, 3.5},
			{3.5, 2.5},
		}
		got := collectFloat64s(Permutations(list, 2, false, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("Permutations_repeatable", func(t *testing.T) {
		want := [][]float64{
			{1.5, 1.5},
			{1.5, 2.5},
			{1.5, 3.5},
			{2.5, 1.5},
			{2.5, 2.5},
			{2.5, 3.5},
			{3.5, 1.5},
			{3.5, 2.5},
			{3.5, 3.5},
		}
		got := collectFloat64s(Permutations(list, 2, true, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestBugFixEdgeCasesNoPanicOrHang(t *testing.T) {
	type mode struct {
		name       string
		fn         func(Replacer, int, bool, int) chan Replacer
		repeatable bool
	}
	modes := []mode{
		{name: "Combinations_non_repeatable", fn: Combinations, repeatable: false},
		{name: "Combinations_repeatable", fn: Combinations, repeatable: true},
		{name: "Permutations_non_repeatable", fn: Permutations, repeatable: false},
		{name: "Permutations_repeatable", fn: Permutations, repeatable: true},
	}

	type edgeCase struct {
		name      string
		list      Replacer
		selectNum int
		buf       int
		wantFor   func(m mode) [][]string
	}

	cases := []edgeCase{
		{
			name:      "selectNum=0",
			list:      StringReplacer{"A", "B"},
			selectNum: 0,
			buf:       1,
			wantFor: func(m mode) [][]string {
				if m.name == "Combinations_non_repeatable" || m.name == "Combinations_repeatable" {
					return [][]string{{}}
				}
				return nil
			},
		},
		{
			name:      "selectNum=-1",
			list:      StringReplacer{"A", "B"},
			selectNum: -1,
			buf:       1,
			wantFor:   func(m mode) [][]string { return nil },
		},
		{
			name:      "buf=-1",
			list:      StringReplacer{"A", "B"},
			selectNum: 1,
			buf:       -1,
			wantFor:   func(m mode) [][]string { return [][]string{{"A"}, {"B"}} },
		},
		{
			name:      "list=nil",
			list:      nil,
			selectNum: 2,
			buf:       1,
			wantFor:   func(m mode) [][]string { return nil },
		},
		{
			name:      "empty_slice_select1",
			list:      StringReplacer{},
			selectNum: 1,
			buf:       1,
			wantFor:   func(m mode) [][]string { return nil },
		},
		{
			name:      "empty_slice_select2",
			list:      StringReplacer{},
			selectNum: 2,
			buf:       1,
			wantFor:   func(m mode) [][]string { return nil },
		},
		{
			name:      "empty_slice_select0",
			list:      StringReplacer{},
			selectNum: 0,
			buf:       1,
			wantFor: func(m mode) [][]string {
				if m.name == "Combinations_non_repeatable" || m.name == "Combinations_repeatable" {
					return [][]string{{}}
				}
				return nil
			},
		},
	}

	for _, m := range modes {
		for _, tc := range cases {
			t.Run(m.name+"/"+tc.name, func(t *testing.T) {
				done := make(chan [][]string, 1)
				errCh := make(chan interface{}, 1)
				go func() {
					defer func() {
						if r := recover(); r != nil {
							errCh <- r
						}
					}()
					done <- collectStrings(m.fn(tc.list, tc.selectNum, m.repeatable, tc.buf))
				}()

				select {
				case r := <-errCh:
					t.Fatalf("unexpected panic: %v", r)
				case got := <-done:
					want := tc.wantFor(m)
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("got %v, want %v", got, want)
					}
				case <-time.After(2 * time.Second):
					t.Fatal("timed out waiting for channel to close")
				}
			})
		}
	}
}

func TestBugFixCombinationsFullLenCopiesSlice(t *testing.T) {
	input := []int{0, 1, 2}
	orig := []int{0, 1, 2}
	for got := range combinations(input, len(input), 1) {
		got[0] = 999
	}
	if !reflect.DeepEqual(input, orig) {
		t.Fatalf("input slice was mutated: got %v, want %v", input, orig)
	}
}
