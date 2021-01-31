package nscript

import "fmt"

type Check struct {
	name   string
	params []string
}

func NewCheck(name string, params []string) *Check {
	var c = &Check{}
	c.name = name
	c.params = params
	return c
}

func (this *Check) exec(ctx Context) (bool, error) {
	var cmd = GetCheckCommand(this.name)
	if cmd == nil {
		return false, fmt.Errorf("%s not found", this.name)
	}
	return cmd(this.name, ctx, this.params...)
}
