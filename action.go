package nscript

import "fmt"

var (
	ActionGoto  = "GOTO"
	ActionBreak = "BREAK"
)

type Action struct {
	key    string
	params []string
}

func NewAction(key string, params []string) *Action {
	var a = &Action{}
	a.key = key
	a.params = params
	return a
}

func (this *Action) exec(ctx Context) error {
	var f = GetActionFunction(this.key)
	if f == nil {
		return fmt.Errorf("%s not found", this.key)
	}
	return f(this.key, ctx, this.params...)
}
