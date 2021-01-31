package main

import (
	"fmt"
	"github.com/smartwalle/nscript"
)

func main() {
	nscript.RegisterCheckFunction("CHECKGOLD", func(key string, ctx nscript.Context, params ...string) (bool, error) {
		fmt.Println(key, ctx, params)
		return true, nil
	})
	nscript.RegisterActionFunction("TAKEGOLD", func(key string, ctx nscript.Context, params ...string) error {
		fmt.Println(key, ctx, params)
		return nil
	})

	var s, err = nscript.NewScript("./npc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s.Exec("@Main", nil))
}
