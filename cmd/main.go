package main

import (
	"fmt"
	"github.com/smartwalle/nscript"
)

func main() {
	var s, err = nscript.NewScript("./npc.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = s
}
