package nscript

import (
	"errors"
	"fmt"
	"github.com/smartwalle/nscript/internal"
)

type Segment struct {
	checks      []*Check
	actions     []*Action
	elseActions []*Action

	says     []string
	elseSays []string
}

func NewSegment() *Segment {
	var s = &Segment{}
	return s
}

func (this *Segment) parse(lines []string) error {
	var keyword string

	for _, line := range lines {
		if line[0] == internal.KeyPrefix {
			keyword = internal.ToUpper(line)
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
			return fmt.Errorf("not found keyword %s", keyword)
		}

	}
	return nil
}

func (this *Segment) parseCheck(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var name = internal.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var check = NewCheck(name, params)
	this.checks = append(this.checks, check)
	return nil
}

func (this *Segment) parseAction(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var name = internal.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var action = NewAction(name, params)
	this.actions = append(this.actions, action)
	return nil
}
func (this *Segment) parseElseAction(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var name = internal.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var action = NewAction(name, params)
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

func (this *Segment) check(ctx Context) (bool, error) {
	for _, check := range this.checks {
		ok, err := check.exec(ctx)
		if err != nil {
			// 若有错误，返回错误
			return false, err
		}
		if ok == false {
			// 若有失败，返回失败
			return false, nil
		}
	}
	// 全部成功，返回成功
	return true, nil
}

func (this *Segment) hasElse() bool {
	return len(this.elseActions) > 0 || len(this.elseSays) > 0
}

func (this *Segment) execAction(ctx Context) ([]string, string, error) {
	return this._execAction(ctx, this.actions, this.says)
}

func (this *Segment) execElseAction(ctx Context) ([]string, string, error) {
	return this._execAction(ctx, this.elseActions, this.elseSays)
}

func (this *Segment) _execAction(ctx Context, actions []*Action, says []string) ([]string, string, error) {
	for _, action := range actions {
		switch action.name {
		case internal.CmdGoto:
			if len(action.params) < 1 {
				return nil, "", errors.New("syntax error: invalid args for GOTO")
			}
			return nil, action.params[0], nil
		case internal.CmdBreak:
			break
		}

		if err := action.exec(ctx); err != nil {
			return nil, "", err
		}
	}
	return says, "", nil
}
