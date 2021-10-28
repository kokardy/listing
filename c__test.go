package listing

import (
	"strings"
	"testing"
)

func TestComtStringSlice(t *testing.T) {
	buf := 5
	list := StringReplacer([]string{
		"ab",
		"cd",
		"ef",
		"gh",
	})
	n := len(list)
	m := 3
	expected := cCount(n, m)
	count := 0
	for comb := range Combinations(list, m, false, buf) {
		count++
		_ = comb
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}

func TestCombInt(t *testing.T) {
	buf := 5
	list := []int{
		1, 2, 3, 4, 5,
	}
	n := len(list)
	m := 4
	expected := cCount(n, m)
	count := 0
	for comb := range combinations(list, m, buf) {
		count++
		_ = comb
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}

func TestRepeatedlyCombInt(t *testing.T) {
	buf := 5
	list := []int{
		1, 2, 3, 4, 5,
	}
	n := len(list)
	m := 4
	expected := hCount(n, m)
	count := 0
	for comb := range repeatedCombinations(list, m, buf) {
		count++
		_ = comb
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}

type CustomStruct struct {
	Name string
	// something
}

func (cs CustomStruct) Stringer() string {
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

func (cr CustomReplacer) Stringer() string {
	css := make([]string, cr.Len())
	for i, cs := range cr {
		css[i] = cs.Stringer()
	}
	return "[" + strings.Join(css, ", ") + "]"
}

func TestCustomStruct(t *testing.T) {
	selectNum := 5
	repeatable := false
	buf := 3
	var custom = CustomReplacer{
		{Name: "custom1"},
		{Name: "custom2"},
		{Name: "custom3"},
		{Name: "custom4"},
		{Name: "custom5"},
		{Name: "custom6"},
	}
	count := 0
	expected := 6
	for comb := range Combinations(custom, selectNum, repeatable, buf) {
		_ = comb
		count++
	}
	if count != expected {
		t.Fatalf("count: %d but expected: %d", count, expected)
	}

}
