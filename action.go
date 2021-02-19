package nscript

import "fmt"

type _Action struct {
	name   string
	params []interface{}
}

func _NewAction(name string, params []interface{}) *_Action {
	var a = &_Action{}
	a.name = name
	a.params = params
	return a
}

func (this *_Action) exec(script *Script, ctx Context) error {
	var cmd = getActionCommand(this.name)
	if cmd == nil {
		return fmt.Errorf("not found action command %s", this.name)
	}
	return cmd(script, ctx, this.params...)
}
