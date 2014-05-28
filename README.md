datafiller
==========
[![Build Status](https://drone.io/github.com/erggo/datafiller/status.png)](https://drone.io/github.com/erggo/datafiller/latest)

A Golang package for filling structs by random data.

# Installation

`go get github.com/erggo/datafiller`

# Sample Usage
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/erggo/datafiller"
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