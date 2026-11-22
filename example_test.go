package listing_test

import (
	"fmt"
	"strings"

	"github.com/kokardy/listing"
)

func ExampleCombinations() {
	ss := listing.StringReplacer{"A", "B", "C", "D"}
	for comb := range listing.Combinations(ss, 3, false, 5) {
		fmt.Println(comb)
	}
	// Output:
	// [A B C]
	// [A B D]
	// [A C D]
	// [B C D]
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

func (cr CustomReplacer) Replace(indices []int) listing.Replacer {
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

func ExampleCombinations_customStruct() {
	custom := CustomReplacer{
		{Name: "custom1"},
		{Name: "custom2"},
		{Name: "custom3"},
		{Name: "custom4"},
		{Name: "custom5"},
		{Name: "custom6"},
	}
	for comb := range listing.Combinations(custom, 5, false, 3) {
		fmt.Println(comb)
	}
	// Output:
	// [custom1, custom2, custom3, custom4, custom5]
	// [custom1, custom2, custom3, custom4, custom6]
	// [custom1, custom2, custom3, custom5, custom6]
	// [custom1, custom2, custom4, custom5, custom6]
	// [custom1, custom3, custom4, custom5, custom6]
	// [custom2, custom3, custom4, custom5, custom6]
}

func ExamplePermutations() {
	ss := listing.StringReplacer{"A", "B", "C", "D"}
	for perm := range listing.Permutations(ss, 3, false, 5) {
		fmt.Println(perm)
	}
	// Output:
	// [A B C]
	// [A C B]
	// [B A C]
	// [B C A]
	// [C A B]
	// [C B A]
	// [A B D]
	// [A D B]
	// [B A D]
	// [B D A]
	// [D A B]
	// [D B A]
	// [A C D]
	// [A D C]
	// [C A D]
	// [C D A]
	// [D A C]
	// [D C A]
	// [B C D]
	// [B D C]
	// [C B D]
	// [C D B]
	// [D B C]
	// [D C B]
}
