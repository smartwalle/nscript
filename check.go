package nscript

import "fmt"

type Check struct {
	key    string
	params []string
}

func NewCheck(key string, params []string) *Check {
	var c = &Check{}
	c.key = key
	c.params = params
	return c
}

func (this *Check) exec(ctx Context) (bool, error) {
	var f = GetCheckCommand(this.key)
	if f == nil {
		return false, fmt.Errorf("%s not found", this.key)
	}
	return f(this.key, ctx, this.params...)
}
