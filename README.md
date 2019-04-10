datafiller
==========
[![Build Status](https://travis-ci.org/tvi/datafiller.svg?branch=master)](https://travis-ci.org/tvi/datafiller)
[![GoDoc](https://godoc.org/github.com/erggo/datafiller?status.png)](https://godoc.org/github.com/erggo/datafiller)

A Golang package for filling structs by random data.

# Installation

`go get github.com/HelpfulHuman/datafiller`

# Sample Usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/HelpfulHuman/datafiller"
)

type S struct {
	A string
	B struct {
		C string
		D string
		E int
	}
}

func main() {
	i := S{}
	datafiller.Fill(&i)
	b, err := json.Marshal(i)
	if err != nil {
		return
	}
	fmt.Println(string(b))
}
```

# TODO

- [x] simple types
- [x] structs
- [x] slices
- [x] packages
- [x] simple tagged structs
- [ ] maps
- [ ] smart tagged struct generation
- [ ] guessing names by names
- [ ] add documentation
- [ ] functions
- [ ] all types
- [ ] add options for filling such as slice length
