package nscript

import "strings"

type CheckFunction func(key string, ctx Context, params ...string) (bool, error)

type ActionFunction func(key string, ctx Context, params ...string) error

var checkFunctions = make(map[string]CheckFunction)
var actionFunctions = make(map[string]ActionFunction)

func RegisterCheckFunction(key string, f CheckFunction) {
	key = strings.TrimSpace(key)
	key = strings.ToUpper(key)

	if key == "" || f == nil {
		return
	}
	checkFunctions[key] = f
}

func GetCheckFunction(key string) CheckFunction {
	key = strings.TrimSpace(key)
	key = strings.ToUpper(key)
	return checkFunctions[key]
}

func RegisterActionFunction(key string, f ActionFunction) {
	key = strings.TrimSpace(key)
	key = strings.ToUpper(key)

	if key == "" || f == nil {
		return
	}
	actionFunctions[key] = f
}

func GetActionFunction(key string) ActionFunction {
	key = strings.TrimSpace(key)
	key = strings.ToUpper(key)
	return actionFunctions[key]
}
