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
	stringDataMap := make(map[string]func() string)

	stringDataMap["address"] = randomdata.Address
	stringDataMap["city"] = randomdata.City
	stringDataMap["email"] = randomdata.Email
	stringDataMap["lastname"] = randomdata.LastName
	stringDataMap["paragraph"] = randomdata.Paragraph
	stringDataMap["street"] = randomdata.Street
	stringDataMap["firstname"] = func() string { return randomdata.FirstName(randomdata.RandomGender) }
	stringDataMap["name"] = func() string { return randomdata.FullName(randomdata.RandomGender) }
	stringDataMap["country"] = func() string { return randomdata.Country(0) }
	stringDataMap["postalcode"] = func() string { return randomdata.PostalCode("US") }
	stringDataMap["state"] = func() string { return randomdata.State(1) }

	tags := strings.Split(structTag, ",")
	// TODO(tvi): Design struct tags ordering.
	if val.Kind() == reflect.String {
		gen, ok := stringDataMap[tags[0]]
		if ok {
			val.SetString(gen())
		}
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

		if val.Kind() == reflect.Int ||
			val.Kind() == reflect.Int8 ||
			val.Kind() == reflect.Int16 ||
			val.Kind() == reflect.Int32 ||
			val.Kind() == reflect.Int64 {
			val.SetInt(rand.Int63n(100))
			return
		} else if val.Kind() == reflect.Uint ||
			val.Kind() == reflect.Uint8 ||
			val.Kind() == reflect.Uint16 ||
			val.Kind() == reflect.Uint32 ||
			val.Kind() == reflect.Uint64 {
			val.SetUint(uint64(rand.Int63n(100)))
			return
		} else if val.Kind() == reflect.Float32 ||
			val.Kind() == reflect.Float64 {
			val.SetFloat(float64(rand.Float32()))
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
