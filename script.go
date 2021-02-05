package nscript

import (
	"fmt"
	"github.com/smartwalle/nscript/internal"
	"io"
	"strings"
)

type Script struct {
	functions map[string]*Function
}

func NewScript(file string) (*Script, error) {
	var iScript, err = internal.LoadFile(file)
	if err != nil {
		return nil, err
	}
	return parseScript(iScript)
}

func LoadFromFile(file string) (*Script, error) {
	return NewScript(file)
}

func LoadFromReader(r io.Reader) (*Script, error) {
	var iScript, err = internal.Load(r)
	if err != nil {
		return nil, err
	}
	return parseScript(iScript)
}

func LoadFromText(text string) (*Script, error) {
	var r = strings.NewReader(text)
	return LoadFromReader(r)
}

func parseScript(iScript *internal.Script) (*Script, error) {
	var nScript = &Script{}
	nScript.functions = make(map[string]*Function)
	for _, iFunc := range iScript.Functions {
		var nFunc = NewFunction(iFunc.Name)
		if err := nFunc.parse(iFunc.Lines); err != nil {
			return nil, err
		}
		nScript.functions[nFunc.name] = nFunc

		if onLoadFunction != nil {
			var matches = internal.RegexFunctionParam.FindStringSubmatch(nFunc.name)
			var args string
			if len(matches) > 1 {
				args = matches[1]
			}
			onLoadFunction(nScript, nFunc.name, args)
		}
	}

	return nScript, nil
}

func (this *Script) Exists(name string) bool {
	var _, exists = this.functions[name]
	return exists
}

func (this *Script) Exec(name string, ctx Context) ([]string, error) {
	name = internal.ToUpper(name)

	// 处理默认方法
	switch name {
	case internal.FuncExit:
		return nil, nil
	case internal.FuncClose:
		return nil, nil
	}

	var function = this.functions[name]
	if function == nil {
		return nil, fmt.Errorf("not found function %s", name)
	}

	var says, nFunc, err = function.exec(this, ctx)
	if err != nil {
		return nil, err
	}

	if nFunc != "" {
		// 执行其它函数
		return this.Exec(nFunc, ctx)
	}

	return says, nil
}
