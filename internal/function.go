package internal

type Function struct {
	Name  string
	Lines []string
}

func NewFunction(name string) *Function {
	return &Function{
		Name: ToUpper(name),
	}
}

func (this *Function) Add(line string) {
	this.Lines = append(this.Lines, line)
}
