package test

import (
	"fmt"
	"github.com/smartwalle/nscript"
	"testing"
)

func Test_Format(t *testing.T) {
	script, err := nscript.NewScript("./format.txt")
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}

	var ctx = NewContext()

	var testTbl = []struct {
		name   string
		gold   int64
		expect string
	}{
		{"名字1", 90000, ""},
		{"名字2", 20001, ""},
		{"名字3", 20000, ""},
		{"名字4", 10001, ""},
		{"名字5", 10000, ""},
		{"名字6", 5001, ""},
		{"名字7", 5000, ""},
		{"名字8", 4999, ""},
		{"名字9", 1, ""},
	}

	for _, test := range testTbl {
		ctx.User.Name = test.name
		ctx.User.Gold = test.gold
		test.expect = fmt.Sprintf("你好 {{%s/RED}}, 你有 %d 金币, {这个不匹配, 我要<Exit/@EXIT>}", test.name, test.gold)

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

func Benchmark_Format(b *testing.B) {
	script, err := nscript.NewScript("./format.txt")
	if err != nil {
		b.Fatal("加载脚本发生错误:", err)
	}

	for i := 0; i < b.N; i++ {
		var ctx = NewContext()

		var testTbl = []struct {
			name   string
			gold   int64
			expect string
		}{
			{"名字1", 90000, ""},
		}

		for _, test := range testTbl {
			ctx.User.Name = test.name
			ctx.User.Gold = test.gold
			test.expect = fmt.Sprintf("你好 {{%s/RED}}, 你有 %d 金币, {这个不匹配, 我要<Exit/@EXIT>}", test.name, test.gold)

			res, err := script.Exec("@MAIN", ctx)
			if err != nil {
				b.Fatal("执行脚本发生错误:", err)
			}

			if len(res) != 1 {
				b.Fatal("脚本结果不符合预期")
			}
			var actual = res[0]

			if actual != test.expect {
				b.Fatal("脚本结果不符合预期, 期望结果:", test.expect, "实际结果:", actual)
			}
		}
	}
}
