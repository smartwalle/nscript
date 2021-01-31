package nscript

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
