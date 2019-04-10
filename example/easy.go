package main

import (
	"datafiller"
	"encoding/json"
	"fmt"
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
