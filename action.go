package nscript

import "fmt"

type Action struct {
	name   string
	params []interface{}
}

func NewAction(name string, params []interface{}) *Action {
	var a = &Action{}
	a.name = name
	a.params = params
	return a
}

func (this *Action) exec(ctx Context) error {
	var cmd = GetActionCommand(this.name)
	if cmd == nil {
		return fmt.Errorf("not found action command %s", this.name)
	}
	return cmd(this.name, ctx, this.params...)
}
