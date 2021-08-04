package internal

// 默认函数
const (
	// FuncExit 退出脚本
	FuncExit = "@EXIT"

	// FuncClose 退出脚本
	FuncClose = "@CLOSE"
)

// 默认指令
const (
	CmdGoto  = "GOTO"
	CmdBreak = "BREAK"
)

const (
	KeyPrefix = '#'

	FunctionPrefix = '@'
)

// 关键字
const (
	// KeyComment 注释
	KeyComment = ';'

	// KeyIf IF 语句，执行判断
	KeyIf = "#IF"

	// KeySay SAY 语句，输出内容
	KeySay = "#SAY"

	// KeyElseSay ELSESAY 语句，输出内容
	KeyElseSay = "#ELSESAY"

	// KeyAct ACT 语句，执行操作
	KeyAct = "#ACT"

	// KeyElseAct ELSEACT 语句，执行操作
	KeyElseAct = "#ELSEACT"

	// KeyInsert 将指定脚本文件中的所有内容引入到当前脚本中
	// 示例：
	// #INSERT [dir1/dir2/file.txt]
	KeyInsert = "#INSERT"

	// KeyInclude 将指定脚本文件中的特定脚本片断(函数)的内容引入到当前脚本中，不包含片断(函数)名
	// 示例：
	// #INCLUDE [dir1/dir2/file.txt] @SECTION_1
	KeyInclude = "#INCLUDE"
)
