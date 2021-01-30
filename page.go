package nscript

import "fmt"

type Page struct {
	key      string
	segments []*Segment
}

func NewPage(key string) *Page {
	var p = &Page{}
	p.key = key
	return p
}

func (this *Page) parse(lines []string) error {
	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}
