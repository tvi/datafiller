package datafiller

import (
	"time"
	"reflect"
	"math/rand"
)

var packages = make(map[string]reflect.Value)

func packages_init() {
	packages["time.Time"] = reflect.ValueOf(time.Unix(rand.Int63n(2000000000),0))
}
