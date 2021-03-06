package nscript

import (
	"errors"
	"fmt"
	"github.com/smartwalle/nscript/internal"
)

type inSegment struct {
	checks      []*inCheck
	actions     []*inAction
	elseActions []*inAction

	says     []string
	elseSays []string
}

func inNewSegment() *inSegment {
	var s = &inSegment{}
	return s
}

func (this *inSegment) parse(lines []string) error {
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
				if err := this.parseSay(line); err != nil {
					return err
				}
				continue
			}
			return fmt.Errorf("not found keyword %s", keyword)
		}

	}
	return nil
}

func (this *inSegment) parseCheck(line string) error {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil
	}
	var name = internal.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var cmdParser = getCheckCommandParser(name)
	nParams, err := cmdParser(name, params...)
	if err != nil {
		return err
	}

	var check = inNewCheck(name, nParams)
	this.checks = append(this.checks, check)
	return nil
}

func (this *inSegment) parseAction(line string) error {
	var action, err = this._parseAction(line)
	if err != nil {
		return err
	}
	if action != nil {
		this.actions = append(this.actions, action)
	}
	return nil
}

func (this *inSegment) parseElseAction(line string) error {
	var action, err = this._parseAction(line)
	if err != nil {
		return err
	}
	if action != nil {
		this.elseActions = append(this.elseActions, action)
	}
	return nil
}

func (this *inSegment) _parseAction(line string) (*inAction, error) {
	var parts = internal.Split(line, " ")
	if len(parts) == 0 {
		return nil, nil
	}
	var name = internal.ToUpper(parts[0])
	var params []string
	if len(parts) > 1 {
		params = parts[1:]
	}

	var cmdParser = getActionCommandParser(name)
	nParams, err := cmdParser(name, params...)
	if err != nil {
		return nil, err
	}
	var action = inNewAction(name, nParams)
	return action, nil
}

func (this *inSegment) parseSay(line string) error {
	this.says = append(this.says, line)
	return nil
}

func (this *inSegment) parseElseSay(line string) error {
	this.elseSays = append(this.elseSays, line)
	return nil
}

func (this *inSegment) check(script *Script, ctx Context) (bool, error) {
	for _, check := range this.checks {
		ok, err := check.exec(script, ctx)
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

func (this *inSegment) hasMainBranch() bool {
	return len(this.actions) > 0 || len(this.says) > 0
}

func (this *inSegment) hasElseBranch() bool {
	return len(this.elseActions) > 0 || len(this.elseSays) > 0
}

func (this *inSegment) execAction(script *Script, ctx Context) (bool, []string, string, error) {
	return this._execAction(script, ctx, this.actions, this.says)
}

func (this *inSegment) execElseAction(script *Script, ctx Context) (bool, []string, string, error) {
	return this._execAction(script, ctx, this.elseActions, this.elseSays)
}

func (this *inSegment) _execAction(script *Script, ctx Context, actions []*inAction, says []string) (bool, []string, string, error) {
	var nBreak bool
	for _, action := range actions {
		switch action.name {
		case internal.CmdGoto:
			if len(action.params) < 1 {
				return false, nil, "", errors.New("syntax error: invalid args for GOTO")
			}
			return false, nil, action.params[0].(string), nil
		case internal.CmdBreak:
			nBreak = true
		}

		if nBreak {
			break
		}

		if err := action.exec(script, ctx); err != nil {
			return false, nil, "", err
		}
	}
	says = this.formatSay(script, ctx, says)
	return nBreak, says, "", nil
}

func (this *inSegment) formatSay(script *Script, ctx Context, says []string) []string {
	return script.Format(ctx, says)
}
