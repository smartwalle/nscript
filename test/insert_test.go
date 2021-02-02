package test

import (
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_Insert(t *testing.T) {
	script, err := nscript.NewScript("./insert.txt")
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()

	var testTbl = []struct {
		gender Gender
		expect string
	}{
		{Female, "女士"},
		{Male, "先生"},
	}

	for _, test := range testTbl {
		ctx.User.Gender = test.gender

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
