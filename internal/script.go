package internal

import "strings"

type Script struct {
	Pages map[string]*Page
}

func NewScript() *Script {
	var s = &Script{}
	s.Pages = make(map[string]*Page)
	return s
}

func (this *Script) Add(p *Page) {
	if p == nil {
		return
	}
	this.Pages[strings.ToUpper(p.Key)] = p
}

func (this *Script) Take(key string) *Page {
	key = strings.ToUpper(key)
	var page = this.Pages[key]
	delete(this.Pages, key)
	return page
}
