package nscript

import (
	"fmt"
	"github.com/smartwalle/nscript/internal"
	"strings"
)

type Segment struct {
	page *Page

	checks      []*Check
	actions     []*Action
	elseActions []*Action

	says     []string
	elseSays []string
}

func NewSegment(page *Page) *Segment {
	var s = &Segment{}
	s.page = page
	return s
}

func (this *Segment) parse(lines []string) error {
	var keyword string

	for _, line := range lines {
		if line[0] == internal.KeyPrefix {
			keyword = strings.ToUpper(line)
			continue
		}

		line = internal.TrimRightSpace(line)

		switch keyword {
		case internal.KeyIf:
			if err := this.parseCheck(line); err != nil {
				return err
			}
		case internal.KeyAct:
			if err := this.parseAction(line); err != nil {
				return err
			}
			continue
		case internal.KeyElseAct:
			if err := this.parseElseAction(line); err != nil {
				return err
			}
			continue
		case internal.KeySay:
			if err := this.parseSay(line); err != nil {
				return err
			}
			continue
		case internal.KeyElseSay:
			if err := this.parseElseSay(line); err != nil {
				return err
			}
			continue
		default:
			if keyword == "" {
				keyword = internal.KeySay
				continue
			}
			return fmt.Errorf("unknown keyword %s", keyword)
		}

	}
	return nil
}

func (this *Segment) parseCheck(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var key = strings.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var check = NewCheck(key, params)
	this.checks = append(this.checks, check)
	return nil
}

func (this *Segment) parseAction(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var key = strings.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var action = NewAction(key, params)
	this.actions = append(this.actions, action)
	return nil
}
func (this *Segment) parseElseAction(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var key = strings.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var action = NewAction(key, params)
	this.elseActions = append(this.elseActions, action)
	return nil
}

func (this *Segment) parseSay(line string) error {
	this.says = append(this.says, line)
	return nil
}

func (this *Segment) parseElseSay(line string) error {
	this.elseSays = append(this.elseSays, line)
	return nil
}

func (this *Segment) Check() bool {
	for _, check := range this.checks {
		if check.Exec() == false {
			return false
		}
	}
	return true
}
