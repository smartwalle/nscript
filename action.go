package nscript

import "fmt"

type inAction struct {
	name   string
	params []interface{}
}

func inNewAction(name string, params []interface{}) *inAction {
	var a = &inAction{}
	a.name = name
	a.params = params
	return a
}

func (this *inAction) exec(script *Script, ctx Context) error {
	var cmd = getActionCommand(this.name)
	if cmd == nil {
		return fmt.Errorf("not found action command %s", this.name)
	}
	return cmd(script, ctx, this.params...)
}
