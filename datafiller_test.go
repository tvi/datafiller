package datafiller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

// Actual tests

func TestUnassignedInt(t *testing.T) {
	var i int
	Fill(&i)
	if reflect.ValueOf(i).Kind() != reflect.Int {
		t.Errorf("Type error: %v, want %v", reflect.ValueOf(i).Kind(), reflect.Int)
	}
	fmt.Println(i)
}

func TestSimpleTypes(t *testing.T) {
	// var InvalidV       invalid
	var BoolV bool
	var IntV int
	var Int8V int8
	var Int16V int16
	var Int32V int32
	var Int64V int64
	var UintV uint
	var Uint8V uint8
	var Uint16V uint16
	var Uint32V uint32
	var Uint64V uint64
	var UintptrV uintptr
	var Float32V float32
	var Float64V float64
	var Complex64V complex64
	var Complex128V complex128

	// var ArrayV         array
	// var ChanV          chan
	// var FuncV          func
	// var InterfaceV     interface
	// var MapV           map
	// var PtrV           ptr
	// var SliceV         slice
	// var StringV        string
	// var StructV        struct
	// var UnsafePointerV unsafepointer

	tests := []struct {
		value        interface{}
		expectedType reflect.Kind
	}{
		{&BoolV, reflect.Bool},
		{&IntV, reflect.Int},
		{&Int8V, reflect.Int8},
		{&Int16V, reflect.Int16},
		{&Int32V, reflect.Int32},
		{&Int64V, reflect.Int64},
		{&UintV, reflect.Uint},
		{&Uint8V, reflect.Uint8},
		{&Uint16V, reflect.Uint16},
		{&Uint32V, reflect.Uint32},
		{&Uint64V, reflect.Uint64},
		{&UintptrV, reflect.Uintptr},
		{&Float32V, reflect.Float32},
		{&Float64V, reflect.Float64},
		{&Complex64V, reflect.Complex64},
		{&Complex128V, reflect.Complex128},
	}

	for _, test := range tests {
		// Type level checking
		testValue := reflect.Indirect(reflect.ValueOf(test.value))
		if testValue.Kind() != test.expectedType {
			t.Errorf("Type error: %v, want %v", testValue.Kind(), test.expectedType)
		}

		Fill(test.value)

		testValue = reflect.Indirect(reflect.ValueOf(test.value))
		if testValue.Kind() != test.expectedType {
			t.Errorf("Type error: %v, want %v", testValue.Kind(), test.expectedType)
		}

		// Value level checking
		ifc := testValue.Interface()
		fmt.Printf("Type: %v; \t\t Value: %v \n", testValue.Kind(), ifc)
		// TODO(tvi): Figure out mock testing.

		zero := reflect.Zero(testValue.Type())
		if reflect.DeepEqual(ifc, zero.Interface()) {
			t.Errorf("Changed value is zero-value (type: %v): value %v, do not want %v", testValue.Type(), zero.Interface(), ifc)
		}
	}
}

// Tests for debugging

func TestDebugSimpleInt(t *testing.T) {
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

func TestDebugSimpleStruct(t *testing.T) {
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

func TestDebugSimpleTimeStruct(t *testing.T) {
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

type D struct {
	Q string `json:"-"`
	A string `json:"myName,omitempty"`
	B string `datafiller:"-"`
	C string `datafiller:"name,omitempty"`
}

func TestDebugSimpleTaggedStruct(t *testing.T) {
	i := D{}
	Fill(&i)

	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fmt.Println(i)
}
