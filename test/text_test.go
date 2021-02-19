package test

import (
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_LoadFromText(t *testing.T) {
	var text = `
[@MAIN]
#IF
CHECKGENDER 1
#SAY
先生你好
#ELSESAY
女士你好
`

	script, err := nscript.LoadFromText(text)
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()
	res, err := script.Exec("@MAIN", ctx)
	if err != nil {
		t.Fatal("执行脚本发生错误:", err)
	}

	if len(res) != 1 {
		t.Fatal("脚本结果不符合预期")
	}

	var expect = "先生你好"
	var actual = res[0]

	if actual != expect {
		t.Fatal("脚本结果不符合预期, 期望结果:", expect, "实际结果:", actual)
	}
}
