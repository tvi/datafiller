// Package datafiller implements function for randomly filling passed
// datastructures by mock sample data.
package datafiller

import (
	"math/rand"
	"reflect"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

func init() {
	packagesInit()
}

const (
	taggedStructKey = "datafiller"
)

func taggedFieldSet(val reflect.Value, structTag string) {
	tags := strings.Split(structTag, ",")
	// TODO(tvi): Design struct tags ordering.
	if tags[0] == "name" && val.Kind() == reflect.String {
		val.SetString(randomdata.FullName(randomdata.RandomGender))
	}
}

func recursiveSet(val reflect.Value) {
	if val.CanSet() {
		var fullPath string
		fullPath = val.Type().PkgPath() + "." + val.Type().Name()
		pkgVal, ok := packages[fullPath]
		if ok {
			val.Set(pkgVal)
			return
		}

		if val.Kind() == reflect.Int {
			val.SetInt(rand.Int63n(100))
			return
		} else if val.Kind() == reflect.Bool {
			val.SetBool(true)
			return
		} else if val.Kind() == reflect.String {
			val.SetString("test")
			return
		} else if val.Kind() == reflect.Struct {
			lngth := val.NumField()
			strType := val.Type()

			for i := 0; i < lngth; i++ {
				if strType.Field(i).Tag.Get(taggedStructKey) == "" {
					recursiveSet(val.Field(i))
				} else if strType.Field(i).Tag.Get(taggedStructKey) == "-" {
				} else {
					advStrTag := strType.Field(i).Tag.Get(taggedStructKey)
					taggedFieldSet(val.Field(i), advStrTag)
				}
			}
			return
		} else if val.Kind() == reflect.Ptr {
			recursiveSet(reflect.Indirect(val))
			return
		} else if val.Kind() == reflect.Slice {
			// TODO(tvi): Fix slice length.
			typ := val.Type()
			elem := typ.Elem()
			nw := reflect.Zero(elem)
			m := reflect.MakeSlice(typ, 0, 1)
			m = reflect.Append(m, nw)
			m = reflect.Append(m, nw)
			val.Set(m)
			recursiveSet(val.Index(0))
			recursiveSet(val.Index(1))
			return
		}
	}
}

// Function Fill takes a pointer to variable of any type and fills the variable
// by with sample data. It panics if the passed value is not a pointer.
func Fill(i interface{}) {
	valPtr := reflect.ValueOf(i)

	if valPtr.Kind() != reflect.Ptr && valPtr.Kind() != reflect.UnsafePointer {
		panic("Passed argument is not a pointer.")
	}

	val := reflect.Indirect(valPtr)
	recursiveSet(val)
}
