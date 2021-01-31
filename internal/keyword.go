package internal

// 默认函数
const (
	// 退出脚本
	FuncExit = "@EXIT"

	// 退出脚本
	FuncClose = "@CLOSE"
)

// 默认指令
const (
	CmdGoto  = "GOTO"
	CmdBreak = "BREAK"
)

// 关键字相关
const (
	// 注释
	KeyComment = ';'

	KeyPrefix = '#'

	// IF 语句，执行判断
	KeyIf = "#IF"

	// SAY 语句，输出内容
	KeySay = "#SAY"

	// ELSESAY 语句，输出内容
	KeyElseSay = "#ELSESAY"

	// ACT 语句，执行操作
	KeyAct = "#ACT"

	// ELSEACT 语句，执行操作
	KeyElseAct = "#ELSEACT"

	// TODO 将指定脚本文件中的所有内容引入到当前脚本中
	// 示例：
	// #INSERT [dir1/dir2/file.txt]
	KeyInsert = "#INSERT"

	// TODO 将指定脚本文件中的特定脚本片断(函数)引入到当前脚本中
	// 示例：
	// #INCLUDE [dir1/dir2/file.txt] @SECTION_1
	KeyInclude = "#INCLUDE"
)
