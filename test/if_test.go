package test

import (
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_IF(t *testing.T) {
	script, err := nscript.NewScript("./if.txt")
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()

	var testTbl = []struct {
		gold   int64
		expect string
	}{
		{90000, "太有钱了"},
		{20001, "太有钱了"},
		{20000, "太有钱了"},
		{10001, "有一定实力"},
		{10000, "有一定实力"},
		{5001, "好好努力"},
		{5000, "好好努力"},
		{4999, "就这样吧"},
		{1, "就这样吧"},
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
	}
}
