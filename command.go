package nscript

import (
	"github.com/smartwalle/nscript/internal"
)

type CheckCommand func(name string, ctx Context, params ...string) (bool, error)
type ActionCommand func(name string, ctx Context, params ...string) error
type FormatCommand func(name string, ctx Context) string

var checkCommands = make(map[string]CheckCommand)
var actionCommands = make(map[string]ActionCommand)
var formatCommands = make(map[string]FormatCommand)

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

func RegisterFormatCommand(name string, f FormatCommand) {
	name = internal.ToUpper(name)

	if name == "" || f == nil {
		return
	}
	formatCommands[name] = f
}

func GetFormatCommand(name string) FormatCommand {
	name = internal.ToUpper(name)
	return formatCommands[name]
}
