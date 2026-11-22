package listing

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func cCount(n, m int) int {
	if n < 0 || m < 0 || m > n {
		return 0
	}
	if m == 0 || m == n {
		return 1
	}
	if m > n-m {
		m = n - m
	}
	res := 1
	for i := 0; i < m; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

func pCount(n, m int) int {
	if n < 0 || m < 0 || m > n {
		return 0
	}
	res := 1
	for i := 0; i < m; i++ {
		res *= n - i
	}
	return res
}

func hCount(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	}
	if m == 0 {
		return 1
	}
	return cCount(n+m-1, m)
}

func TestCountHelpers(t *testing.T) {
	cCases := []struct {
		n, m, want int
	}{
		{-1, 0, 0},
		{0, 0, 1},
		{3, 0, 1},
		{4, 3, 4},
		{5, 4, 5},
		{2, 3, 0},
		{5, -1, 0},
	}
	for _, tc := range cCases {
		if got := cCount(tc.n, tc.m); got != tc.want {
			t.Errorf("cCount(%d, %d) = %d, want %d", tc.n, tc.m, got, tc.want)
		}
	}

	if got := pCount(5, 4); got != 120 {
		t.Errorf("pCount(5, 4) = %d, want 120", got)
	}
	if got := hCount(5, 4); got != 70 {
		t.Errorf("hCount(5, 4) = %d, want 70", got)
	}
}

type CustomStruct struct {
	Name string
}

func (cs CustomStruct) String() string {
	return cs.Name
}

type CustomReplacer []CustomStruct

func (cr CustomReplacer) Len() int {
	return len(cr)
}

func (cr CustomReplacer) Replace(indices []int) Replacer {
	result := make(CustomReplacer, len(indices))
	for i, idx := range indices {
		result[i] = cr[idx]
	}
	return result
}

func (cr CustomReplacer) String() string {
	css := make([]string, cr.Len())
	for i, cs := range cr {
		css[i] = cs.String()
	}
	return "[" + strings.Join(css, ", ") + "]"
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

func collectInts(ch chan Replacer) [][]int {
	var got [][]int
	for item := range ch {
		ir := item.(IntReplacer)
		cp := make([]int, len(ir))
		copy(cp, ir)
		got = append(got, cp)
	}
	return got
}

func collectRunes(ch chan Replacer) [][]rune {
	var got [][]rune
	for item := range ch {
		rr := item.(RuneReplacer)
		cp := make([]rune, len(rr))
		copy(cp, rr)
		got = append(got, cp)
	}
	return got
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

func assertStringsEqual(t *testing.T, got, want [][]string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v (len=%d), want %v (len=%d)", got, len(got), want, len(want))
	}
}

func assertNoPanicOrHang(t *testing.T, fn func() [][]string, want [][]string) {
	t.Helper()
	done := make(chan [][]string, 1)
	errCh := make(chan interface{}, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errCh <- r
			}
		}()
		done <- fn()
	}()

	select {
	case r := <-errCh:
		t.Fatalf("unexpected panic: %v", r)
	case got := <-done:
		assertStringsEqual(t, got, want)
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for channel to close")
	}
}

type edgeCaseSpec struct {
	name      string
	list      Replacer
	selectNum int
	buf       int
	want      [][]string
}

func runEdgeCases(t *testing.T, gen func(Replacer, int, bool, int) chan Replacer, repeatable bool, cases []edgeCaseSpec) {
	t.Helper()
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assertNoPanicOrHang(t, func() [][]string {
				return collectStrings(gen(tc.list, tc.selectNum, repeatable, tc.buf))
			}, tc.want)
		})
	}
}
