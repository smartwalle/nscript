package test

import (
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_Act(t *testing.T) {
	script, err := nscript.NewScript("./act.txt")
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()

	var testTbl = []struct {
		before int64
		after  int64
		expect string
	}{
		{1001, 1, "谢谢老板"},
		{1000, 0, "谢谢老板"},
		{999, 999, "你说什么?"},
	}

	for _, test := range testTbl {
		ctx.User.Gold = test.before

		res, err := script.Exec("@MAIN", ctx)
		if err != nil {
			t.Fatal("执行脚本发生错误:", err)
		}

		if len(res) != 1 {
			t.Fatal("脚本结果不符合预期")
		}
		var actual = res[0]

		if actual != test.expect {
			t.Fatal("脚本结果不符合预期, 期望结果:", test.expect, "实际结果:", actual)
		}

		if ctx.User.Gold != test.after {
			t.Fatal("剩余数量不符合预期, 期望结果:", test.after, "实际结果:", ctx.User.Gold)
		}
	}
}
