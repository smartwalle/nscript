package main

import (
	"fmt"
	"github.com/smartwalle/nscript"
)

func main() {
	nscript.RegisterCheckCommand("CHECKGOLD", func(name string, ctx nscript.Context, params ...string) (bool, error) {
		fmt.Println(name, ctx, params)
		return true, nil
	})
	nscript.RegisterActionCommand("TAKEGOLD", func(name string, ctx nscript.Context, params ...string) error {
		fmt.Println(name, ctx, params)
		return nil
	})

	nscript.RegisterFormatCommand("$USERNAME", func(name string, ctx nscript.Context) string {
		return "SmartWalle"
	})

	nscript.RegisterFormatCommand("$GOLD", func(name string, ctx nscript.Context) string {
		return "1000"
	})

	var s, err = nscript.NewScript("./npc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s.Exec("@Main", nil))
}
