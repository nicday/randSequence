# randSequence

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/nicday/randSequence)
[![Circle CI](https://circleci.com/gh/nicday/randSequence.svg?style=svg)](https://circleci.com/gh/nicday/randSequence)

randSequence is a random sequence generator with sane defaults and the ability to customize the character set for generating sequences from.

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
