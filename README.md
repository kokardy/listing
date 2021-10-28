# listing

## package for listing combinations and permutations for golang

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
	//You can use any types that implement listing.Replacer interface (Len(), Replace([]int))


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


### To use your custom struct:

If you implement "Len()" and "Replace([]int) Replacer",
you can use your custom struct:

```go
type CustomStruct struct {
	Name string
	// something
}

// no need
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

for comb := range Combinations(custom, selectNum, repeatable, buf) {
  fmt.Println(comb)

```

output:
```
[{custom1} {custom2} {custom3} {custom4} {custom5}]
[{custom1} {custom2} {custom3} {custom4} {custom6}]
[{custom1} {custom2} {custom3} {custom5} {custom6}]
[{custom1} {custom2} {custom4} {custom5} {custom6}]
[{custom1} {custom3} {custom4} {custom5} {custom6}]
[{custom2} {custom3} {custom4} {custom5} {custom6}]
```
