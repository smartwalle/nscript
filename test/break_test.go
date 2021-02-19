package test

import (
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_Break(t *testing.T) {
	script, err := nscript.NewScript("./break.txt")
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()
	ctx.User.Gender = Female

	var testTbl = []struct {
		gold   int64
		level  int64
		age    int64
		nAge   int64
		expect string
	}{
		{999, 10, 18, 18, "金币不能低于 1000"},
		{1000, 9, 18, 18, "等级不能低于 10 级"},
		{1000, 10, 17, 17, "不要太贪心"},
		{1000, 10, 18, 16, "永葆青春"},
	}

	for _, test := range testTbl {
		ctx.User.Gold = test.gold
		ctx.User.Level = test.level
		ctx.User.Age = test.age

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

		if ctx.User.Age != test.nAge {
			t.Fatal("脚本结果不符合预期，期望结果:", test.nAge, "实际结果:", ctx.User.Age)
		}
	}
}
