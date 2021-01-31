package nscript

import (
	"github.com/smartwalle/nscript/internal"
)

type CheckFunction func(key string, ctx Context, params ...string) (bool, error)

type ActionFunction func(key string, ctx Context, params ...string) error

var checkFunctions = make(map[string]CheckFunction)
var actionFunctions = make(map[string]ActionFunction)

func RegisterCheckFunction(key string, f CheckFunction) {
	key = internal.ToUpper(key)

	if key == "" || f == nil {
		return
	}
	checkFunctions[key] = f
}

func GetCheckFunction(key string) CheckFunction {
	key = internal.ToUpper(key)
	return checkFunctions[key]
}

func RegisterActionFunction(key string, f ActionFunction) {
	key = internal.ToUpper(key)

	if key == "" || f == nil {
		return
	}
	actionFunctions[key] = f
}

func GetActionFunction(key string) ActionFunction {
	key = internal.ToUpper(key)
	return actionFunctions[key]
}
