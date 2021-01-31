package nscript

import (
	"github.com/smartwalle/nscript/internal"
	"strings"
)

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
	var sLines []string
	for idx, line := range lines {
		if idx != 0 && strings.HasPrefix(line, internal.KeyIf) {
			// 如果不是第一行，并且又发现了 #IF
			// 则表示一个代码片断结束了，需要开启新的代码片断
			var nSegment = NewSegment(this)
			if err := nSegment.parse(sLines); err != nil {
				return err
			}
			this.segments = append(this.segments, nSegment)
			sLines = nil
		}
		sLines = append(sLines, line)
	}
	if len(sLines) > 0 {
		var nSegment = NewSegment(this)
		if err := nSegment.parse(sLines); err != nil {
			return err
		}
		this.segments = append(this.segments, nSegment)
		sLines = nil
	}
	return nil
}

func (this *Page) Exec(...interface{}) ([]string, error) {
	for _, seg := range this.segments {
		if seg.Check() {
			// 执行 act  say
		} else {
			// 执行 else act    else say
		}
	}
	return nil, nil
}
