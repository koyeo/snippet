package snippet

import (
	"path/filepath"
)

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
		ident := Join(v.getDir(), v.getName())
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

func (p *Snippets) render(project *Project) {
	for _, v := range p.snippetList {

		distFile := v.makePrefix + v.name + v.makeSuffix + v.suffix
		customFile := v.name + v.suffix

		distPath := filepath.Join(project.root, v.getDir())
		makePath := filepath.Join(distPath, distFile)
		customPath := filepath.Join(distPath, customFile)

		//filter := v.makePrefix != "" || v.makeSuffix != ""
		project.writer.compareSnippet(v, customPath)
		content, err := v.render(v)
		if err != nil {
			Fatal("Render error: ", err)
		}

		if content != "" {
			content, err = v.formatter(content)
			if err != nil {
				Fatal("Format error: ", err)
			}
			project.writer.addMakeRenderFile(distPath, makePath, customPath, content, true)
		}
	}
}
