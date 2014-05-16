listing
=======

package for listing combinations and permutations


sample code:
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

	select_num := 3
	repeatable := false
	buf := 5

	fmt.Println("Permutations")
	for perm := range listing.Permutations(ss, select_num, repeatable, buf) {
		fmt.Println(perm)
	}

	fmt.Println("Combinations")
	for comb := range listing.Combinations(ss, select_num, repeatable, buf) {
		fmt.Println(comb)
	}

}

```

output:
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
