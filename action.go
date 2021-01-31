package nscript

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

func (this *Action) Exec() {
}
