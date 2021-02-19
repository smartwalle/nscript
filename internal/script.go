package internal

type Script struct {
	Sections map[string]*Section
}

func NewScript() *Script {
	var s = &Script{}
	s.Sections = make(map[string]*Section)
	return s
}

func (this *Script) Add(p *Section) {
	if p == nil {
		return
	}
	this.Sections[p.Name] = p
}
