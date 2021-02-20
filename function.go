package nscript

import (
	"github.com/smartwalle/nscript/internal"
	"strings"
)

type inFunction struct {
	name     string
	segments []*inSegment
}

func inNewFunction(name string) *inFunction {
	var f = &inFunction{}
	f.name = name
	return f
}

func (this *inFunction) parse(lines []string) error {
	var sLines []string
	for idx, line := range lines {
		if idx != 0 && strings.HasPrefix(line, internal.KeyIf) {
			// 如果不是第一行，并且又发现了 #IF
			// 则表示一个代码片断结束了，需要开启新的代码片断
			var nSegment = inNewSegment()
			if err := nSegment.parse(sLines); err != nil {
				return err
			}

			// 判断代码块是否有执行分支
			// #IF
			// CMD1
			// #IF
			// CMD2
			// #ACT
			// CMD3
			// 上述代码会把 CMD1 和 CMD2 作为并列条件存在，因为第一个 IF 语句没有任何有效的分支
			if nSegment.hasMainBranch() || nSegment.hasElseBranch() {
				this.segments = append(this.segments, nSegment)
				sLines = nil
			}
		}
		sLines = append(sLines, line)
	}
	if len(sLines) > 0 {
		var nSegment = inNewSegment()
		if err := nSegment.parse(sLines); err != nil {
			return err
		}
		this.segments = append(this.segments, nSegment)
		sLines = nil
	}
	return nil
}

func (this *inFunction) exec(script *Script, ctx Context) ([]string, string, error) {
	var nBreak bool    // 是否 break
	var nSays []string // 输出内容
	var nGoto string   // 是否需要跳转到其它方法
	var ok bool        // 条件判断结果
	var err error

	for _, seg := range this.segments {
		ok, err = seg.check(script, ctx)
		if err != nil {
			// 若有错误，返回错误
			return nil, "", err
		}

		if ok {
			// 执行 Action
			nBreak, nSays, nGoto, err = seg.execAction(script, ctx)
		} else if seg.hasElseBranch() {
			// 执行 ElseAction
			nBreak, nSays, nGoto, err = seg.execElseAction(script, ctx)
		}

		if err != nil || nGoto != "" || nBreak {
			return nSays, nGoto, err
		}
	}
	return nSays, "", nil
}
