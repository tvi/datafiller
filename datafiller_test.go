package datafiller

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSimpleInt(t *testing.T) {
	i := 1
	Fill(&i)
	fmt.Println(i)
}

type S struct {
	A string
	B struct {
		C string
		D string
		E int
	}
}

func TestSimpleStruct(t *testing.T) {
	i := S{}
	Fill(&i)

	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fmt.Println(i)
}
