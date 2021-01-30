package nscript

import "fmt"

type Page struct {
	Name string
}

func NewPage(name string) *Page {
	var p = &Page{}
	p.Name = name
	return p
}

func (this *Page) parse(lines []string) error {
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}
