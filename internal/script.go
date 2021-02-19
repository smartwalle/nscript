package internal

type Script struct {
	Values    map[string]*Section
	Functions map[string]*Section
}

func NewScript() *Script {
	var s = &Script{}
	s.Values = make(map[string]*Section)
	s.Functions = make(map[string]*Section)
	return s
}

func (this *Script) Add(p *Section) {
	if p == nil {
		return
	}
	if p.Function {
		this.Functions[p.Name] = p
	} else {
		this.Values[p.Name] = p
	}
}
