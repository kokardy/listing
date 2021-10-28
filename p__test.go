package listing

import (
	"testing"
)

func TestPermStringSlice(t *testing.T) {
	buf := 5
	list := StringReplacer([]string{
		"a", "b", "c", "d", "e",
	})
	n := len(list)
	m := 4
	expected := pCount(n, m)
	count := 0
	for perm := range Permutations(list, m, false, buf) {
		count++
		//t.Log(count, comb)
		_ = perm
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}

func TestPermInt(t *testing.T) {
	buf := 5
	list := []int{
		1, 2, 3, 4, 5,
	}
	n := len(list)
	m := 4
	expected := pCount(n, m)
	count := 0
	for perm := range permutations(list, m, buf) {
		count++
		//t.Log(count, comb)
		_ = perm
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}

func TestRepeatedlyPermInt(t *testing.T) {
	buf := 5
	list := []int{
		1, 2, 3, 4,
	}
	n := len(list)
	m := 2
	expected := 1
	for i := 0; i < m; i++ {
		expected *= n
	}
	count := 0
	for perm := range repeatedPermutations(list, m, buf) {
		count++
		//t.Log(count, comb)
		_ = perm
	}
	if count != expected {
		t.Fatalf("%dC%d != %d", n, m, expected)
	}
}
