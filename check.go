package nscript

import "fmt"

type inCheck struct {
	name   string
	params []interface{}
}

func inNewCheck(name string, params []interface{}) *inCheck {
	var c = &inCheck{}
	c.name = name
	c.params = params
	return c
}

func (this *inCheck) exec(script *Script, ctx Context) (bool, error) {
	var cmd = getCheckCommand(this.name)
	if cmd == nil {
		return false, fmt.Errorf("not found check command %s", this.name)
	}
	return cmd(script, ctx, this.params...)
}
