package datafiller

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSimpleInt(t *testing.T) {
	i := 1
	Fill(&i)
	fmt.Println(i)
}

type S struct {
	A string
	B []struct {
		Q []struct {
			W int
		}
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

type A struct {
	T time.Time
	Q string
}

func TestSimpleTimeStruct(t *testing.T) {
	i := A{}
	Fill(&i)

	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fmt.Println(i)
}
