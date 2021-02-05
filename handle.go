package nscript

type LoadFunction func(script *Script, name, arg string)

// 初始化方法回调: 方法加载完成触发。
var onLoadFunction LoadFunction

func OnLoadFunction(f LoadFunction) {
	onLoadFunction = f
}
