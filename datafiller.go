package datafiller

import "reflect"
import "math/rand"

func recursiveSet(val reflect.Value) {
	if val.CanSet() {
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
			for i := 0; i < lngth; i++ {
				recursiveSet(val.Field(i))
			}
			return
		} else if val.Kind() == reflect.Ptr {
			recursiveSet(reflect.Indirect(val))
			return
		}
	}
}

func Fill(i interface{}) {
	valPtr := reflect.ValueOf(i)

	if valPtr.Kind() != reflect.Ptr && valPtr.Kind() != reflect.UnsafePointer {
		panic("Incorrect type.")
	}

	val := reflect.Indirect(valPtr)
	recursiveSet(val)
}
