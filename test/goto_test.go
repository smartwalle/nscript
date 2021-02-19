package test

import (
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_GOTO(t *testing.T) {
	script, err := nscript.NewScript("./goto.txt")
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()

	var testTbl = []struct {
		gold   int64
		nGold  int64
		expect string
	}{
		{999, 999, "没有足够的金币"},
		{1000, 0, "有钱就是好"},
		{1001, 1, "有钱就是好"},
	}

	for _, test := range testTbl {
		ctx.User.Gold = test.gold

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

		if ctx.User.Gold != test.nGold {
			t.Fatal("脚本结果不符合预期，期望结果:", test.nGold, "实际结果:", ctx.User.Gold)
		}
	}
}
