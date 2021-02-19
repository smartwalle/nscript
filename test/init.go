package test

import (
	"errors"
	"fmt"
	"github.com/smartwalle/nscript"
	"strconv"
)

type Gender int8

const (
	Male   Gender = 1
	Female Gender = 2
)

type User struct {
	Name   string
	Age    int64
	Gender Gender
	Gold   int64
	Level  int64
}

type Context struct {
	User *User
}

func NewContext() *Context {
	var c = &Context{}
	c.User = &User{}
	c.User.Name = "Test User"
	c.User.Age = 18
	c.User.Gender = Male
	c.User.Gold = 10000
	c.User.Level = 1
	return c
}

func parseInt64(v string) int64 {
	nv, _ := strconv.ParseInt(v, 10, 64)
	return nv
}

func init() {
	// 回调
	nscript.OnLoadFunction(func(script *nscript.Script, name, arg string) {
		fmt.Println("Load Function:", name, arg)
	})

	// 解析器
	nscript.RegisterCommandParser("CHECKGOLD", func(params ...string) ([]interface{}, error) {
		if len(params) != 2 {
			return nil, errors.New("CHECKGOLD 指令参数异常")
		}
		var nParams = make([]interface{}, len(params))
		nParams[0] = params[0]
		nParams[1] = parseInt64(params[1])
		return nParams, nil
	})

	nscript.RegisterCommandParser("CHECKGENDER", func(params ...string) ([]interface{}, error) {
		if len(params) != 1 {
			return nil, errors.New("CHECKGENDER 指令参数异常")
		}
		var nParams = make([]interface{}, len(params))
		nParams[0] = parseInt64(params[0])
		return nParams, nil
	})

	nscript.RegisterCommandParser("CHECKAGE", func(params ...string) ([]interface{}, error) {
		if len(params) != 2 {
			return nil, errors.New("CHECKAGE 指令参数异常")
		}
		var nParams = make([]interface{}, len(params))
		nParams[0] = params[0]
		nParams[1] = parseInt64(params[1])
		return nParams, nil
	})

	nscript.RegisterCommandParser("CHECKLEVEL", func(params ...string) ([]interface{}, error) {
		if len(params) != 2 {
			return nil, errors.New("CHECKLEVEL 指令参数异常")
		}
		var nParams = make([]interface{}, len(params))
		nParams[0] = params[0]
		nParams[1] = parseInt64(params[1])
		return nParams, nil
	})

	nscript.RegisterCommandParser("TAKEGOLD", func(params ...string) ([]interface{}, error) {
		if len(params) != 1 {
			return nil, errors.New("TAKEGOLD 指令参数异常")
		}
		var nParams = make([]interface{}, len(params))
		nParams[0] = parseInt64(params[0])
		return nParams, nil
	})

	nscript.RegisterCommandParser("SETAGE", func(params ...string) ([]interface{}, error) {
		if len(params) != 1 {
			return nil, errors.New("SETAGE 指令参数异常")
		}
		var nParams = make([]interface{}, len(params))
		nParams[0] = parseInt64(params[0])
		return nParams, nil
	})

	// 判断条件
	nscript.RegisterCheckCommand("CHECKGOLD", func(script *nscript.Script, ctx nscript.Context, params ...interface{}) (bool, error) {
		var op = params[0].(string)
		var value = params[1].(int64)
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64(op, nCtx.User.Gold, value), nil
	})
	nscript.RegisterCheckCommand("CHECKGENDER", func(script *nscript.Script, ctx nscript.Context, params ...interface{}) (bool, error) {
		var value = params[0].(int64)
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64("=", int64(nCtx.User.Gender), value), nil
	})
	nscript.RegisterCheckCommand("CHECKAGE", func(script *nscript.Script, ctx nscript.Context, params ...interface{}) (bool, error) {
		var op = params[0].(string)
		var value = params[1].(int64)
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64(op, nCtx.User.Age, value), nil
	})
	nscript.RegisterCheckCommand("CHECKLEVEL", func(script *nscript.Script, ctx nscript.Context, params ...interface{}) (bool, error) {
		var op = params[0].(string)
		var value = params[1].(int64)
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64(op, nCtx.User.Level, value), nil
	})

	// 操作
	nscript.RegisterActionCommand("TAKEGOLD", func(script *nscript.Script, ctx nscript.Context, params ...interface{}) error {
		var nCtx = ctx.(*Context)
		var gold = params[0].(int64)
		if gold <= 0 || gold > nCtx.User.Gold {
			return errors.New("没有足够的金币")
		}
		nCtx.User.Gold -= gold
		return nil
	})
	nscript.RegisterActionCommand("SETAGE", func(script *nscript.Script, ctx nscript.Context, params ...interface{}) error {
		var nCtx = ctx.(*Context)
		var age = params[0].(int64)
		nCtx.User.Age = age
		return nil
	})

	// 格式化
	nscript.RegisterFormatCommand("$USERNAME", func(script *nscript.Script, ctx nscript.Context) string {
		var nCtx = ctx.(*Context)
		return nCtx.User.Name
	})
	nscript.RegisterFormatCommand("$GOLD", func(script *nscript.Script, ctx nscript.Context) string {
		var nCtx = ctx.(*Context)
		return fmt.Sprintf("%d", nCtx.User.Gold)
	})
}
