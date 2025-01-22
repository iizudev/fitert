# fitert
fitert is a collection of functional transformers for iterators presented in go 1.23.

**fitert** stands for _"**f**unctional **iter** **t**ransformations"_

[![Build & Test](https://github.com/iizudev/fitert/actions/workflows/go.yml/badge.svg)](https://github.com/iizudev/fitert/actions/workflows/go.yml)

## Example
Here is a short example of what this library looks like in action
```go
package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/iizudev/fitert"
)

func main() {
	names := []string{"mary", "michael", "david", "lucia"}
	p := fitert.Splitted(
		fitert.Mapped(slices.Values(names), strings.Title),
		func(n string) (string, int) { return n, len(n) },
	)
	for n, l := range p {
		fmt.Printf("name %s has %d characters\n", n, l)
	}
}
```

The output would be:
```
name Mary has 4 characters
name Michael has 7 characters
name David has 5 characters
name Lucia has 5 characters
```

### Explaination

Here is the explaination of what each function in the example does

#### `Mapped`

The function `Mapped` acts just like in other languages that have
functional elements in their std; it converts an `iter.Seq` of one type
into an `iter.Seq` of another type — how an element of the first type
is _mapped_ to an element of another type is determined by the mapper
function you provide.

#### `Splitted`

`Splitted` is an interesting function that was invented simply because
for some reason go's developers decided to introduce two type (`iter.Seq`
and `iter.Seq2`) instead of one — this function coverts an `iter.Seq` into
an `iter.Seq2`.

## Notes

There is no such a function as `ForEach` because Go's built-in `for` already does the job. However
it's a decision to discuss — open an issue if you believe that `ForEach` would be helpful.

The functions that covert sequences into other sequences (such as `Mapped` and `Filtered`)
are called this way — `Mapped` and not `Map` — because sequences are _lazy_, so when you call
the function you don't actually change anything about the original sequence, but create a new one
that acts as an altered version of it.

## Contribution

If you find a bug or wish to add functionality to this library then open an issue first.
