package snippet

import (
	"fmt"
	"github.com/koyeo/snippet/logger"
	"github.com/koyeo/snippet/storage"
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
		ident := storage.Join(v.getDir(), v.Name())
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
		project.writer.compareSnippet(v, customPath)
		if v.ignore {
			continue
		}

		if v.render == nil {
			logger.Fatal("Render snippet error: ", fmt.Errorf("\"%s\" not set render func", distPath))
		}

		content, err := v.render(project.ctx, v)
		if err != nil {
			logger.Fatal("Render snippet error: ", err)
		}
		if v.trimSpace {
			content = TrimSpace(content)
		}

		if content != "" {
			if v.formatter != nil {
				content, err = v.formatter(content)
				if err != nil {
					logger.Fatal("Format snippet error: ", err)
				}
			}
			project.writer.addMakeRenderFile(distPath, makePath, customPath, content, true)
		}
	}
}
