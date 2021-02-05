package nscript

import "fmt"

type Check struct {
	name   string
	params []interface{}
}

func NewCheck(name string, params []interface{}) *Check {
	var c = &Check{}
	c.name = name
	c.params = params
	return c
}

func (this *Check) exec(script *Script, ctx Context) (bool, error) {
	var cmd = GetCheckCommand(this.name)
	if cmd == nil {
		return false, fmt.Errorf("not found check command %s", this.name)
	}
	return cmd(script, ctx, this.params...)
}
