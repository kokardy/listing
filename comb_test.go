package listing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationsNonRepeatable(t *testing.T) {
	t.Run("README_example", func(t *testing.T) {
		list := StringReplacer{"A", "B", "C", "D"}
		want := [][]string{
			{"A", "B", "C"},
			{"A", "B", "D"},
			{"A", "C", "D"},
			{"B", "C", "D"},
		}
		assertStringsEqual(t, collectStrings(Combinations(list, 3, false, 5)), want)
	})

	t.Run("StringReplacer_count_and_values", func(t *testing.T) {
		list := StringReplacer{"ab", "cd", "ef", "gh"}
		m := 3
		want := [][]string{
			{"ab", "cd", "ef"},
			{"ab", "cd", "gh"},
			{"ab", "ef", "gh"},
			{"cd", "ef", "gh"},
		}
		got := collectStrings(Combinations(list, m, false, 5))
		if len(got) != cCount(len(list), m) {
			t.Fatalf("C(%d, %d): got %d items, want %d", len(list), m, len(got), cCount(len(list), m))
		}
		assertStringsEqual(t, got, want)
	})

	t.Run("IntReplacer", func(t *testing.T) {
		list := IntReplacer{10, 20, 30}
		want := [][]int{{10, 20}, {10, 30}, {20, 30}}
		got := collectInts(Combinations(list, 2, false, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("RuneReplacer", func(t *testing.T) {
		list := RuneReplacer{'x', 'y', 'z'}
		want := [][]rune{{'x', 'y'}, {'x', 'z'}, {'y', 'z'}}
		got := collectRunes(Combinations(list, 2, false, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("Float64Replacer", func(t *testing.T) {
		list := Float64Replacer{1.5, 2.5, 3.5}
		want := [][]float64{{1.5, 2.5}, {1.5, 3.5}, {2.5, 3.5}}
		got := collectFloat64s(Combinations(list, 2, false, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("CustomReplacer", func(t *testing.T) {
		custom := CustomReplacer{
			{Name: "custom1"},
			{Name: "custom2"},
			{Name: "custom3"},
			{Name: "custom4"},
			{Name: "custom5"},
			{Name: "custom6"},
		}
		want := []string{
			"[custom1, custom2, custom3, custom4, custom5]",
			"[custom1, custom2, custom3, custom4, custom6]",
			"[custom1, custom2, custom3, custom5, custom6]",
			"[custom1, custom2, custom4, custom5, custom6]",
			"[custom1, custom3, custom4, custom5, custom6]",
			"[custom2, custom3, custom4, custom5, custom6]",
		}
		var got []string
		for comb := range Combinations(custom, 5, false, 3) {
			got = append(got, fmt.Sprint(comb))
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("edge_cases", func(t *testing.T) {
		cases := []edgeCaseSpec{
			{name: "selectNum=0", list: StringReplacer{"A", "B", "C"}, selectNum: 0, buf: 2, want: [][]string{{}}},
			{name: "selectNum=1", list: StringReplacer{"A", "B", "C"}, selectNum: 1, buf: 2, want: [][]string{{"A"}, {"B"}, {"C"}}},
			{name: "selectNum=len", list: StringReplacer{"A", "B", "C"}, selectNum: 3, buf: 2, want: [][]string{{"A", "B", "C"}}},
			{name: "selectNum_gt_len", list: StringReplacer{"A", "B"}, selectNum: 3, buf: 2, want: nil},
			{name: "selectNum=-1", list: StringReplacer{"A", "B"}, selectNum: -1, buf: 1, want: nil},
			{name: "buf=-1", list: StringReplacer{"A", "B"}, selectNum: 1, buf: -1, want: [][]string{{"A"}, {"B"}}},
			{name: "list=nil", list: nil, selectNum: 2, buf: 1, want: nil},
			{name: "empty_list_selectNum=0", list: StringReplacer{}, selectNum: 0, buf: 1, want: [][]string{{}}},
			{name: "empty_list_selectNum=1", list: StringReplacer{}, selectNum: 1, buf: 1, want: nil},
			{name: "empty_list_selectNum=2", list: StringReplacer{}, selectNum: 2, buf: 1, want: nil},
		}
		runEdgeCases(t, Combinations, false, cases)
	})
}

func TestCombinationsRepeatable(t *testing.T) {
	t.Run("StringReplacer", func(t *testing.T) {
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
				assertStringsEqual(t, collectStrings(Combinations(tc.list, tc.selectNum, true, tc.buf)), tc.want)
			})
		}
	})

	t.Run("Float64Replacer", func(t *testing.T) {
		list := Float64Replacer{1.5, 2.5, 3.5}
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

	t.Run("edge_cases", func(t *testing.T) {
		cases := []edgeCaseSpec{
			{name: "selectNum=0", list: StringReplacer{"A", "B"}, selectNum: 0, buf: 1, want: [][]string{{}}},
			{name: "selectNum=-1", list: StringReplacer{"A", "B"}, selectNum: -1, buf: 1, want: nil},
			{name: "buf=-1", list: StringReplacer{"A", "B"}, selectNum: 1, buf: -1, want: [][]string{{"A"}, {"B"}}},
			{name: "list=nil", list: nil, selectNum: 2, buf: 1, want: nil},
			{name: "empty_list_selectNum=0", list: StringReplacer{}, selectNum: 0, buf: 1, want: [][]string{{}}},
			{name: "empty_list_selectNum=1", list: StringReplacer{}, selectNum: 1, buf: 1, want: nil},
			{name: "empty_list_selectNum=2", list: StringReplacer{}, selectNum: 2, buf: 1, want: nil},
		}
		runEdgeCases(t, Combinations, true, cases)
	})
}

func TestInternalCombinations(t *testing.T) {
	t.Run("combinations_count", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		m := 4
		want := cCount(len(list), m)
		count := 0
		for range combinations(list, m, 5) {
			count++
		}
		if count != want {
			t.Fatalf("C(%d, %d): got %d, want %d", len(list), m, count, want)
		}
	})

	t.Run("repeatedCombinations_count", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		m := 4
		want := hCount(len(list), m)
		count := 0
		for range repeatedCombinations(list, m, 5) {
			count++
		}
		if count != want {
			t.Fatalf("H(%d, %d): got %d, want %d", len(list), m, count, want)
		}
	})

	t.Run("full_len_copies_slice", func(t *testing.T) {
		input := []int{0, 1, 2}
		orig := []int{0, 1, 2}
		for got := range combinations(input, len(input), 1) {
			got[0] = 999
		}
		if !reflect.DeepEqual(input, orig) {
			t.Fatalf("input slice was mutated: got %v, want %v", input, orig)
		}
	})
}
