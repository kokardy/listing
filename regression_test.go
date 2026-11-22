package listing

import (
	"reflect"
	"testing"
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
