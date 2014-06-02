package datafiller

import (
	"math/rand"
	"reflect"
	"time"
)

var packages = make(map[string]reflect.Value)

func packagesInit() {
	packages["time.Time"] = reflect.ValueOf(time.Unix(rand.Int63n(2000000000), 0))
}
