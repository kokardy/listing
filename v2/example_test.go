package listing_test

import (
	"fmt"
	"slices"

	"github.com/kokardy/listing/v2"
)

func ExampleCombinations() {
	items := []string{"A", "B", "C", "D"}
	for comb := range listing.Combinations(items, 3) {
		fmt.Println(comb)
	}
	// Output:
	// [A B C]
	// [A B D]
	// [A C D]
	// [B C D]
}

func ExampleCombinationsWithReplacement() {
	items := []string{"A", "B", "C"}
	for comb := range listing.CombinationsWithReplacement(items, 2) {
		fmt.Println(comb)
	}
	// Output:
	// [A A]
	// [A B]
	// [A C]
	// [B B]
	// [B C]
	// [C C]
}

func ExamplePermutations() {
	items := []string{"A", "B", "C"}
	for perm := range listing.Permutations(items, 2) {
		fmt.Println(perm)
	}
	// Output:
	// [A B]
	// [A C]
	// [B A]
	// [B C]
	// [C A]
	// [C B]
}

func ExamplePermutationsWithReplacement() {
	items := []int{0, 1}
	for perm := range listing.PermutationsWithReplacement(items, 3) {
		fmt.Println(perm)
	}
	// Output:
	// [0 0 0]
	// [0 0 1]
	// [0 1 0]
	// [0 1 1]
	// [1 0 0]
	// [1 0 1]
	// [1 1 0]
	// [1 1 1]
}

func ExampleCombinations_customStruct() {
	type Item struct {
		Name string
	}
	type ItemList []Item

	items := ItemList{
		{Name: "alpha"},
		{Name: "beta"},
		{Name: "gamma"},
	}

	collected := slices.Collect(listing.Combinations(items, 2))
	for _, comb := range collected {
		fmt.Printf("%T: %v\n", comb, comb)
	}
	// Output:
	// listing_test.ItemList: [{alpha} {beta}]
	// listing_test.ItemList: [{alpha} {gamma}]
	// listing_test.ItemList: [{beta} {gamma}]
}
