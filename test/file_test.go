package test

import (
	"github.com/smartwalle/nscript"
	"os"
	"testing"
)

func Test_LoadFromFile(t *testing.T) {
	script, err := nscript.LoadFromFile("./file.txt")
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

func Test_LoadFromReader(t *testing.T) {
	reader, err := os.Open("./file.txt")
	if err != nil {
		t.Fatal("读取文件发生错误:", err)
	}
	script, err := nscript.LoadFromReader(reader)
	if err != nil {
		t.Fatal("加载脚本发生错误:", err)
	}
	reader.Close()

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
