package nscript

type Segment struct {
	page *Page

	checks      []*Check
	actions     []*Action
	elseActions []*Action

	says     []string
	elseSays []string
}
