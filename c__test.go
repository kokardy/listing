package listing

import (
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
	expected := C(n, m)
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
	expected := C(n, m)
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
	expected := H(n, m)
	count := 0
	for comb := range repeated_combinations(list, m, buf) {
		count++
		_ = comb
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}
