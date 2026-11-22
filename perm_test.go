package listing

import (
	"reflect"
	"testing"
)

func TestPermutationsNonRepeatable(t *testing.T) {
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
		assertStringsEqual(t, collectStrings(Permutations(list, 3, false, 5)), want)
	})

	t.Run("StringReplacer_count", func(t *testing.T) {
		list := StringReplacer{"a", "b", "c", "d", "e"}
		m := 4
		want := pCount(len(list), m)
		got := collectStrings(Permutations(list, m, false, 5))
		if len(got) != want {
			t.Fatalf("P(%d, %d): got %d items, want %d", len(list), m, len(got), want)
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
		got := collectInts(Permutations(list, 2, false, 2))
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
		got := collectRunes(Permutations(list, 2, false, 2))
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("Float64Replacer", func(t *testing.T) {
		list := Float64Replacer{1.5, 2.5, 3.5}
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

	t.Run("edge_cases", func(t *testing.T) {
		cases := []edgeCaseSpec{
			{name: "selectNum=0", list: StringReplacer{"A", "B", "C"}, selectNum: 0, buf: 2, want: nil},
			{name: "selectNum=1", list: StringReplacer{"A", "B", "C"}, selectNum: 1, buf: 2, want: [][]string{{"A"}, {"B"}, {"C"}}},
			{
				name:      "selectNum=len",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 3,
				buf:       2,
				want: [][]string{
					{"A", "B", "C"},
					{"A", "C", "B"},
					{"B", "A", "C"},
					{"B", "C", "A"},
					{"C", "A", "B"},
					{"C", "B", "A"},
				},
			},
			{name: "selectNum_gt_len", list: StringReplacer{"A", "B"}, selectNum: 3, buf: 2, want: nil},
			{name: "selectNum=-1", list: StringReplacer{"A", "B"}, selectNum: -1, buf: 1, want: nil},
			{name: "buf=-1", list: StringReplacer{"A", "B"}, selectNum: 1, buf: -1, want: [][]string{{"A"}, {"B"}}},
			{name: "list=nil", list: nil, selectNum: 2, buf: 1, want: nil},
			{name: "empty_list_selectNum=0", list: StringReplacer{}, selectNum: 0, buf: 1, want: nil},
			{name: "empty_list_selectNum=1", list: StringReplacer{}, selectNum: 1, buf: 1, want: nil},
			{name: "empty_list_selectNum=2", list: StringReplacer{}, selectNum: 2, buf: 1, want: nil},
		}
		runEdgeCases(t, Permutations, false, cases)
	})
}

func TestPermutationsRepeatable(t *testing.T) {
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
				buf:       3,
				want: [][]string{
					{"A", "A"},
					{"A", "B"},
					{"A", "C"},
					{"B", "A"},
					{"B", "B"},
					{"B", "C"},
					{"C", "A"},
					{"C", "B"},
					{"C", "C"},
				},
			},
			{
				name:      "selectNum=1",
				list:      StringReplacer{"A", "B", "C"},
				selectNum: 1,
				buf:       1,
				want:      [][]string{{"A"}, {"B"}, {"C"}},
			},
			{
				name:      "selectNum_gt_len",
				list:      StringReplacer{"A", "B"},
				selectNum: 3,
				buf:       2,
				want: [][]string{
					{"A", "A", "A"},
					{"A", "A", "B"},
					{"A", "B", "A"},
					{"A", "B", "B"},
					{"B", "A", "A"},
					{"B", "A", "B"},
					{"B", "B", "A"},
					{"B", "B", "B"},
				},
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				assertStringsEqual(t, collectStrings(Permutations(tc.list, tc.selectNum, true, tc.buf)), tc.want)
			})
		}
	})

	t.Run("Float64Replacer", func(t *testing.T) {
		list := Float64Replacer{1.5, 2.5, 3.5}
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

	t.Run("edge_cases", func(t *testing.T) {
		cases := []edgeCaseSpec{
			{name: "selectNum=0", list: StringReplacer{"A", "B"}, selectNum: 0, buf: 1, want: nil},
			{name: "selectNum=-1", list: StringReplacer{"A", "B"}, selectNum: -1, buf: 1, want: nil},
			{name: "buf=-1", list: StringReplacer{"A", "B"}, selectNum: 1, buf: -1, want: [][]string{{"A"}, {"B"}}},
			{name: "list=nil", list: nil, selectNum: 2, buf: 1, want: nil},
			{name: "empty_list_selectNum=0", list: StringReplacer{}, selectNum: 0, buf: 1, want: nil},
			{name: "empty_list_selectNum=1", list: StringReplacer{}, selectNum: 1, buf: 1, want: nil},
			{name: "empty_list_selectNum=2", list: StringReplacer{}, selectNum: 2, buf: 1, want: nil},
		}
		runEdgeCases(t, Permutations, true, cases)
	})
}

func TestInternalPermutations(t *testing.T) {
	t.Run("permutations_count", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		m := 4
		want := pCount(len(list), m)
		count := 0
		for range permutations(list, m, 5) {
			count++
		}
		if count != want {
			t.Fatalf("P(%d, %d): got %d, want %d", len(list), m, count, want)
		}
	})

	t.Run("repeatedPermutations_count", func(t *testing.T) {
		list := []int{1, 2, 3, 4}
		m := 2
		want := 16
		count := 0
		for range repeatedPermutations(list, m, 5) {
			count++
		}
		if count != want {
			t.Fatalf("repeated P(%d, %d): got %d, want %d", len(list), m, count, want)
		}
	})
}
