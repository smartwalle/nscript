package test

import (
	"errors"
	"fmt"
	"github.com/smartwalle/conv4go"
	"github.com/smartwalle/nscript"
)

type Gender int8

const (
	Male   Gender = 1
	Female Gender = 2
)

type User struct {
	Name   string
	Age    int
	Gender Gender
	Gold   int64
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
	return c
}

func init() {
	// check
	nscript.RegisterCheckCommand("CHECKGOLD", func(name string, ctx nscript.Context, params ...string) (bool, error) {
		var op = params[0]
		var value = params[1]
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64(op, nCtx.User.Gold, value), nil
	})
	nscript.RegisterCheckCommand("CHECKGENDER", func(name string, ctx nscript.Context, params ...string) (bool, error) {
		var value = params[0]
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64("=", nCtx.User.Gender, value), nil
	})
	nscript.RegisterCheckCommand("CHECKAGE", func(name string, ctx nscript.Context, params ...string) (bool, error) {
		var op = params[0]
		var value = params[1]
		var nCtx = ctx.(*Context)
		return nscript.CompareInt64(op, nCtx.User.Age, value), nil
	})

	// action
	nscript.RegisterActionCommand("TAKEGOLD", func(name string, ctx nscript.Context, params ...string) error {
		var nCtx = ctx.(*Context)
		var value = params[0]
		var gold = conv4go.Int64(value)
		if gold <= 0 || gold > nCtx.User.Gold {
			return errors.New("没有足够的金币")
		}
		nCtx.User.Gold -= gold
		return nil
	})

	// format
	nscript.RegisterFormatCommand("$USERNAME", func(name string, ctx nscript.Context) string {
		var nCtx = ctx.(*Context)
		return nCtx.User.Name
	})
	nscript.RegisterFormatCommand("$GOLD", func(name string, ctx nscript.Context) string {
		var nCtx = ctx.(*Context)
		return fmt.Sprintf("%d", nCtx.User.Gold)
	})
}
