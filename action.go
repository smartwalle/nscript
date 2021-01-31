package nscript

import "fmt"

type Action struct {
	name   string
	params []string
}

func NewAction(name string, params []string) *Action {
	var a = &Action{}
	a.name = name
	a.params = params
	return a
}

func (this *Action) exec(ctx Context) error {
	var cmd = GetActionCommand(this.name)
	if cmd == nil {
		return fmt.Errorf("not found command %s", this.name)
	}
	return cmd(this.name, ctx, this.params...)
}
