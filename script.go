package nscript

import (
	"fmt"
	"github.com/smartwalle/nscript/internal"
)

type Script struct {
	pages map[string]*Page
}

func NewScript(file string) (*Script, error) {
	var iScript, err = internal.LoadFile(file)
	if err != nil {
		return nil, err
	}

	var nScript = &Script{}
	nScript.pages = make(map[string]*Page)

	for _, iPage := range iScript.Pages {
		var nPage = NewPage(iPage.Key)

		fmt.Println(iPage.Key)

		if err = nPage.parse(iPage.Lines); err != nil {
			return nil, err
		}

		nScript.pages[nPage.Name] = nPage
	}

	return nScript, nil
}
