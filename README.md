# listing

Go package for enumerating combinations and permutations (with and without replacement).

## Module Layout

- **v2 (Recommended, Go 1.23+)**: `github.com/kokardy/listing/v2` — Generic `iter.Seq` API, zero goroutine leaks, no `Replacer` boilerplate.
- **v1 (Legacy / Backward-compatible, Go 1.17+)**: `github.com/kokardy/listing` — Channel-based `Replacer` API.

---

## v2 Usage (Recommended, Go 1.23+)

### Install

```bash
go get github.com/kokardy/listing/v2
```

Import `"github.com/kokardy/listing/v2"` (package name `listing`).

### Functions

- `listing.Combinations(items, k)` — Generates `k`-length combinations without replacement in lexicographic order.
- `listing.CombinationsWithReplacement(items, k)` — Generates `k`-length combinations with replacement in lexicographic order.
- `listing.Permutations(items, k)` — Generates `k`-length permutations without replacement in lexicographic order.
- `listing.PermutationsWithReplacement(items, k)` — Generates `k`-length permutations with replacement in lexicographic order.

### Example

```go
package main

import (
	"fmt"

	"github.com/kokardy/listing/v2"
)

func main() {
	items := []string{"A", "B", "C"}
	k := 2

	fmt.Println("Permutations")
	for perm := range listing.Permutations(items, k) {
		fmt.Println(perm)
	}

	fmt.Println("PermutationsWithReplacement")
	for perm := range listing.PermutationsWithReplacement(items, k) {
		fmt.Println(perm)
	}

	fmt.Println("Combinations")
	for comb := range listing.Combinations(items, k) {
		fmt.Println(comb)
	}

	fmt.Println("CombinationsWithReplacement")
	for comb := range listing.CombinationsWithReplacement(items, k) {
		fmt.Println(comb)
	}
}
```

Output:

```
Permutations
[A B]
[A C]
[B A]
[B C]
[C A]
[C B]
PermutationsWithReplacement
[A A]
[A B]
[A C]
[B A]
[B B]
[B C]
[C A]
[C B]
[C C]
Combinations
[A B]
[A C]
[B C]
CombinationsWithReplacement
[A A]
[A B]
[A C]
[B B]
[B C]
[C C]
```

### Custom Structs, `slices.Collect`, and Early `break`

With generics, you can pass any slice `[]T` directly without implementing `Replacer`. Because v2 uses Go 1.23 `iter.Seq` iterators without background goroutines, breaking out of a `for ... range` loop early is completely safe and causes zero goroutine leaks.

```go
package main

import (
	"fmt"
	"slices"

	"github.com/kokardy/listing/v2"
)

type CustomStruct struct {
	Name string
}

func main() {
	items := []CustomStruct{
		{Name: "custom1"},
		{Name: "custom2"},
		{Name: "custom3"},
		{Name: "custom4"},
	}

	// Collect all 3-combinations into a slice
	allCombs := slices.Collect(listing.Combinations(items, 3))
	fmt.Printf("Total combinations: %d\n", len(allCombs))

	// Early break is safe with zero goroutine leaks
	for comb := range listing.Combinations(items, 3) {
		fmt.Println(comb)
		break
	}
}
```

Output:

```
Total combinations: 4
[{custom1} {custom2} {custom3}]
```

---

## v1 Usage (Legacy, Go 1.17+)

### Install

```bash
go get github.com/kokardy/listing
```

Import `"github.com/kokardy/listing"`.

### Basic Example

```go
package main

import (
	"fmt"

	"github.com/kokardy/listing"
)

func main() {
	ss := listing.StringReplacer([]string{
		"A", "B", "C", "D",
	})
	// You can use any type that implements listing.Replacer (Len() int, Replace([]int) Replacer).

	selectNum := 3
	repeatable := false
	buf := 5

	fmt.Println("Permutations")
	for perm := range listing.Permutations(ss, selectNum, repeatable, buf) {
		fmt.Println(perm)
	}

	fmt.Println("Combinations")
	for comb := range listing.Combinations(ss, selectNum, repeatable, buf) {
		fmt.Println(comb)
	}
}
```

Output:

```
Permutations
[A B C]
[A C B]
[B A C]
[B C A]
[C A B]
[C B A]
[A B D]
[A D B]
[B A D]
[B D A]
[D A B]
[D B A]
[A C D]
[A D C]
[C A D]
[C D A]
[D A C]
[D C A]
[B C D]
[B D C]
[C B D]
[C D B]
[D B C]
[D C B]
Combinations
[A B C]
[A B D]
[A C D]
[B C D]
```

### Custom Struct Example (v1)

If your slice type implements `Len() int` and `Replace([]int) listing.Replacer`, you can use custom structs with v1. Implementing `fmt.Stringer` (`String() string`) formats the output nicely.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/kokardy/listing"
)

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

func main() {
	selectNum := 5
	repeatable := false
	buf := 3
	custom := CustomReplacer{
		{Name: "custom1"},
		{Name: "custom2"},
		{Name: "custom3"},
		{Name: "custom4"},
		{Name: "custom5"},
		{Name: "custom6"},
	}

	for comb := range listing.Combinations(custom, selectNum, repeatable, buf) {
		fmt.Println(comb)
	}
}
```

Output:

```
[custom1, custom2, custom3, custom4, custom5]
[custom1, custom2, custom3, custom4, custom6]
[custom1, custom2, custom3, custom5, custom6]
[custom1, custom2, custom4, custom5, custom6]
[custom1, custom3, custom4, custom5, custom6]
[custom2, custom3, custom4, custom5, custom6]
```

---

## Migration Guide (v1 -> v2)

| v1 (`github.com/kokardy/listing`) | v2 (`github.com/kokardy/listing/v2`) | Notes |
| --- | --- | --- |
| `StringReplacer` / `Replacer` | `[]T` (any slice type) | Generic; no wrapper interface required |
| `Combinations(list, k, false, buf)` | `Combinations(items, k)` | Buffer argument removed |
| `Combinations(list, k, true, buf)` | `CombinationsWithReplacement(items, k)` | Explicit function for replacement |
| `Permutations(list, k, false, buf)` | `Permutations(items, k)` | Lexicographic order in v2 |
| `Permutations(list, k, true, buf)` | `PermutationsWithReplacement(items, k)` | Explicit function for replacement |
| `chan Replacer` | `iter.Seq[S]` | Compatible with `for ... range` and `slices.Collect` |

### Behavioral Differences to Note

- **Ordering in `Permutations`**: v1 generates permutations grouped by combination first (e.g., all permutations of `[A B C]`, then `[A B D]`), whereas v2 yields permutations in standard lexicographic index order.
- **`selectNum == 0` in `Permutations`**: v1 yields `0` items when `selectNum == 0`, whereas v2 follows standard combinatorics (`P(n, 0) = 1`) and yields `1` empty slice (`[]`).
- **Early termination (`break`)**: v1 uses channels backed by goroutines, so breaking early without draining the channel can leak goroutines; v2 uses `iter.Seq`, so early `break` is always safe with zero goroutine leaks.
