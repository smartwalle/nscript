package nscript

import (
	"fmt"
	"github.com/smartwalle/nscript/internal"
	"io"
	"strings"
)

type Script struct {
	pages map[string]*Page
}

func NewScript(file string) (*Script, error) {
	var iScript, err = internal.LoadFile(file)
	if err != nil {
		return nil, err
	}
	return parseScript(iScript)
}

func LoadFromFile(file string) (*Script, error) {
	return NewScript(file)
}

func LoadFromReader(r io.Reader) (*Script, error) {
	var iScript, err = internal.Load(r)
	if err != nil {
		return nil, err
	}
	return parseScript(iScript)
}

func LoadFromText(text string) (*Script, error) {
	var r = strings.NewReader(text)
	return LoadFromReader(r)
}

func parseScript(iScript *internal.Script) (*Script, error) {
	var pages = make(map[string]*Page)
	for _, iPage := range iScript.Pages {
		var nPage = NewPage(iPage.Key)
		if err := nPage.parse(iPage.Lines); err != nil {
			return nil, err
		}
		pages[nPage.key] = nPage
	}
	var nScript = &Script{}
	nScript.pages = pages
	return nScript, nil
}

func (this *Script) Exec(key string, ctx Context) ([]string, error) {
	key = strings.ToUpper(key)
	var page = this.pages[key]
	if page == nil {
		return nil, fmt.Errorf("%s not found", key)
	}

	var says, gotoKey, err = page.exec(ctx)
	if err != nil {
		return nil, err
	}

	if gotoKey != "" {
		return this.Exec(gotoKey, ctx)
	}

	return says, nil
}
