package datafiller

import (
	"math/rand"
	"reflect"
	"time"

	null "gopkg.in/guregu/null.v3"
)

var packages = make(map[string]reflect.Value)

func packagesInit() {
	randSeed := rand.New(rand.NewSource(-1))
	packages["time.Time"] = reflect.ValueOf(time.Unix(rand.Int63n(2000000000), 0))
	// guregu/null structs have Valid field that is a bool.
	// this field was occasionally randomly set to false, meaning whatever value was passsed
	// would be interpreted as null. Always want some value, so always set valid = true for null packages
	packages["gopkg.in/guregu/null.v3.String"] = reflect.ValueOf(null.NewString("test", true))
	packages["gopkg.in/guregu/null.v3.Int"] = reflect.ValueOf(null.NewInt(rand.Int63n(100), true))
	if randSeed.Int63n(2) == 0 {
		packages["gopkg.in/guregu/null.v3.Bool"] = reflect.ValueOf(null.NewBool(false, true))
	} else {
		packages["gopkg.in/guregu/null.v3.Bool"] = reflect.ValueOf(null.NewBool(true, true))
	}
	packages["gopkg.in/guregu/null.v3.Float"] = reflect.ValueOf(null.NewFloat(float64(randSeed.Float32()), true))
	packages["gopkg.in/guregu/null.v3.Time"] = reflect.ValueOf(null.NewTime(time.Unix(rand.Int63n(2000000000), 0), true))
}
