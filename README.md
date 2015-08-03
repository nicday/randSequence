# randSequence

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/nicday/randSequence)

This repository holds a random sequence generator.

By default a new `randSequence` will generate a sequence from lower and upper case alphanumeric characters, however a `randSequence` can be created with any source characters.


## Example

```go
package main

import (
	"fmt"

	"github.com/nicday/randSequence"
)

func main() {
	generator := randSequence.New()
	sequence := generator.Generate(10)

	fmt.Println(sequence)
}
```
