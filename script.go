package nscript

import (
	"fmt"
	"github.com/smartwalle/nscript/internal"
	"io"
	"strings"
)

type Script struct {
	vars      map[string][]string
	functions map[string]*inFunction
}

func NewScript(file string) (*Script, error) {
	return LoadFromFile(file)
}

func LoadFromFile(file string) (*Script, error) {
	var iScript, err = internal.LoadFile(file)
	if err != nil {
		return nil, err
	}
	return parseScript(iScript)
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
	nScript.vars = make(map[string][]string)
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
	nScript.vars[section.Name] = section.Lines
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

func (this *Script) ReloadFromFile(file string) error {
	var nScript, err = LoadFromFile(file)
	if err != nil {
		return err
	}
	this.vars = nScript.vars
	this.functions = nScript.functions
	return nil
}

func (this *Script) ReloadFromReader(r io.Reader) error {
	var nScript, err = LoadFromReader(r)
	if err != nil {
		return err
	}
	this.vars = nScript.vars
	this.functions = nScript.functions
	return nil
}

func (this *Script) ReloadFromText(text string) error {
	var r = strings.NewReader(text)
	return this.ReloadFromReader(r)
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

func (this *Script) VarExists(key string) bool {
	var _, exists = this.vars[key]
	return exists
}

func (this *Script) Var(key string) string {
	var values = this.vars[key]
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func (this *Script) Vars(key string) []string {
	return this.vars[key]
}

func (this *Script) Value(ctx Context, key string) string {
	if key[0] != '$' {
		return key
	}
	return this.getValue(ctx, key, key)
}

func (this *Script) getValue(ctx Context, key, dValue string) string {
	var param string
	var idx = strings.IndexByte(key, '|')
	if idx > -1 {
		param = key[idx+1:]
		key = key[0:idx]
	}

	var valueCmd = getValueCommand(key)
	if valueCmd == nil {
		return dValue
	}
	return valueCmd(this, ctx, param)
}

func (this *Script) Format(ctx Context, texts []string) []string {
	var nTexts = make([]string, 0, len(texts))
	for _, say := range texts {
		var nSay = internal.RegexFormat.ReplaceAllStringFunc(say, func(s string) string {
			var key = s[1 : len(s)-1]

			return this.getValue(ctx, key, s)
		})
		nTexts = append(nTexts, nSay)
	}
	return nTexts
}
