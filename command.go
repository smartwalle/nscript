package nscript

import (
	"github.com/smartwalle/nscript/internal"
)

type CommandParser func(params ...string) ([]interface{}, error)
type CheckCommand func(script *Script, ctx Context, params ...interface{}) (bool, error)
type ActionCommand func(script *Script, ctx Context, params ...interface{}) error
type ValueCommand func(script *Script, ctx Context, param string) string

// 解析指令：解析脚本的过程中，用于对各指令进行解析，比如判断参数个数，转换参数类型。
var checkCommandParsers = make(map[string]CommandParser)
var actionCommandParsers = make(map[string]CommandParser)

// 判断指令：用于在 #IF 语句块中进行逻辑判断，当其返回的 error 不为空时，该 error 将会返回给调用者。
var checkCommands = make(map[string]CheckCommand)

// 操作指令：用于在 #ACT 和 #ELSEACT 语句块中执行具体的操作，当其返回的 error 不为空时，该 error 将会返回给调用者。
var actionCommands = make(map[string]ActionCommand)

// 值指令：用于定义在 #SAY 和 #ELSESAY 语句块输出的动态内容。
var valueCommands = make(map[string]ValueCommand)

var defaultCommandParser = func(params ...string) ([]interface{}, error) {
	var nParams = make([]interface{}, 0, len(params))
	for _, param := range params {
		nParams = append(nParams, param)
	}
	return nParams, nil
}

func getCheckCommandParser(name string) CommandParser {
	name = internal.ToUpper(name)
	var f = checkCommandParsers[name]
	if f == nil {
		f = defaultCommandParser
	}
	return f
}

func getActionCommandParser(name string) CommandParser {
	name = internal.ToUpper(name)
	var f = actionCommandParsers[name]
	if f == nil {
		f = defaultCommandParser
	}
	return f
}

func RegisterCheckCommand(name string, parser CommandParser, check CheckCommand) {
	name = internal.ToUpper(name)

	if name == "" || check == nil {
		return
	}
	if parser != nil {
		checkCommandParsers[name] = parser
	}
	checkCommands[name] = check
}

func getCheckCommand(name string) CheckCommand {
	name = internal.ToUpper(name)
	return checkCommands[name]
}

func RegisterActionCommand(name string, parser CommandParser, action ActionCommand) {
	name = internal.ToUpper(name)

	if name == "" || action == nil {
		return
	}
	if parser != nil {
		actionCommandParsers[name] = parser
	}
	actionCommands[name] = action
}

func getActionCommand(name string) ActionCommand {
	name = internal.ToUpper(name)
	return actionCommands[name]
}

func RegisterValueCommand(name string, f ValueCommand) {
	name = internal.ToUpper(name)

	if name == "" || f == nil {
		return
	}
	valueCommands[name] = f
}

func getValueCommand(name string) ValueCommand {
	name = internal.ToUpper(name)
	return valueCommands[name]
}
