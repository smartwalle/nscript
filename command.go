package nscript

import (
	"github.com/smartwalle/nscript/internal"
)

type CheckCommand func(name string, ctx Context, params ...string) (bool, error)
type ActionCommand func(name string, ctx Context, params ...string) error

var checkCommands = make(map[string]CheckCommand)
var actionCommands = make(map[string]ActionCommand)

func RegisterCheckCommand(name string, f CheckCommand) {
	name = internal.ToUpper(name)

	if name == "" || f == nil {
		return
	}
	checkCommands[name] = f
}

func GetCheckCommand(name string) CheckCommand {
	name = internal.ToUpper(name)
	return checkCommands[name]
}

func RegisterActionCommand(name string, f ActionCommand) {
	name = internal.ToUpper(name)

	if name == "" || f == nil {
		return
	}
	actionCommands[name] = f
}

func GetActionCommand(name string) ActionCommand {
	name = internal.ToUpper(name)
	return actionCommands[name]
}
