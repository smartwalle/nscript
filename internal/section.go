package internal

type Section struct {
	Name     string
	Function bool
	Lines    []string
}

func NewSection(name string) *Section {
	return &Section{
		Name:     ToUpper(name),
		Function: name[0] == '@',
	}
}

func (this *Section) Add(line string) {
	this.Lines = append(this.Lines, line)
}
