package nscript

import (
	"fmt"
	"github.com/smartwalle/nscript/internal"
	"io"
	"strings"
)

type Script struct {
	values    map[string][]string
	functions map[string]*inFunction
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
	nScript.values = make(map[string][]string)
	nScript.functions = make(map[string]*inFunction)

	for _, section := range iScript.Values {
		if err := parseValue(nScript, section); err != nil {
			return nil, err
		}
	}

	for _, section := range iScript.Functions {
		if err := parseFunction(nScript, section); err != nil {
			return nil, err
		}
	}

	return nScript, nil
}

func parseValue(nScript *Script, section *internal.Section) error {
	nScript.values[section.Name] = section.Lines
	return nil
}

func parseFunction(nScript *Script, section *internal.Section) error {
	var nFunc = inNewFunction(section.Name)
	if err := nFunc.parse(section.Lines); err != nil {
		return err
	}
	nScript.functions[nFunc.name] = nFunc
	return nil
}

func (this *Script) Functions() []string {
	var fList = make([]string, 0, len(this.functions))
	for key := range this.functions {
		fList = append(fList, key)
	}
	return fList
}

func (this *Script) FunctionExists(name string) bool {
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

func (this *Script) ValueExists(key string) bool {
	var _, exists = this.values[key]
	return exists
}

func (this *Script) Value(key string) string {
	var values = this.values[key]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func (this *Script) Values(key string) []string {
	return this.values[key]
}

func (this *Script) Var(ctx Context, key string) string {
	if key[0] != '$' {
		return key
	}
	return this.getVar(ctx, key, key)
}

func (this *Script) getVar(ctx Context, key, dValue string) string {
	var param string
	//var matches = internal.RegexVar.FindStringSubmatch(key)
	var idx = strings.IndexByte(key, '|')
	if idx > -1 {
		param = key[idx+1:]
		key = key[0:idx]
	}

	var varCmd = getVarCommand(key)
	if varCmd == nil {
		return dValue
	}
	return varCmd(this, ctx, param)
}
