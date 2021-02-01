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

	var s, err = nscript.NewScript("./c.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s.Exec("@Main", nil))
}
