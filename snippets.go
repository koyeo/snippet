package snippet

func NewSnippets() *Snippets {
	return &Snippets{}
}

type Snippets struct {
	snippetList []*Snippet
	snippetMap  map[string]*Snippet
}

func (p *Snippets) Add(snippets ...*Snippet) {

	if p.snippetMap == nil {
		p.snippetMap = make(map[string]*Snippet)
	}

	for _, v := range snippets {
		ident := Join(v.GetFilePath(), v.GetFileName())
		if old, ok := p.snippetMap[ident]; !ok {
			p.snippetList = append(p.snippetList, v)
			p.snippetMap[ident] = v
		} else {
			old.Merge(v)
		}
	}
}

func (p *Snippets) All() (snippets []*Snippet) {
	for _, v := range p.snippetList {
		snippets = append(snippets, v)
	}
	return
}
