package internal

import "strings"

type Page struct {
	Key   string
	Lines []string
}

func NewPage(key string) *Page {
	return &Page{
		Key: strings.ToUpper(key),
	}
}

func (this *Page) Add(line string) {
	this.Lines = append(this.Lines, line)
}
