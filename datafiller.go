package datafiller

import "fmt"
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
		} else if val.Kind() == reflect.Slice {
			// TODO(tvi): Fix slice length.
			typ := val.Type()
			elem := typ.Elem()
			nw := reflect.Zero(elem)
			m := reflect.MakeSlice(typ,0,1)
			m = reflect.Append(m, nw)
			m = reflect.Append(m, nw)
			val.Set(m)
			recursiveSet(val.Index(0))
			recursiveSet(val.Index(1))
			return
		} else if val.Kind() == reflect.Map {
			typ := val.Type()
			m := reflect.MakeMap(typ)

			k := reflect.Zero(typ.Key())
			v := reflect.Zero(typ.Elem())

			recursiveSet(k)
			recursiveSet(v)
			m.SetMapIndex(k,v)
			val.Set(m)
			fmt.Println("asdf")

			// m.SetMapIndex(l,b)
			// val.Set(m)
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
