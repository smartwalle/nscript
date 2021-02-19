package nscript

import "fmt"

type _Check struct {
	name   string
	params []interface{}
}

func _NewCheck(name string, params []interface{}) *_Check {
	var c = &_Check{}
	c.name = name
	c.params = params
	return c
}

func (this *_Check) exec(script *Script, ctx Context) (bool, error) {
	var cmd = getCheckCommand(this.name)
	if cmd == nil {
		return false, fmt.Errorf("not found check command %s", this.name)
	}
	return cmd(script, ctx, this.params...)
}
