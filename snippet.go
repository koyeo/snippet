package snippet

import (
	"path"
)

type RenderFunc func(snippet *Snippet) (content, suffix string, err error)
type FormatterFunc func(content string) (string, error)

func NewSnippet(suffix string) *Snippet {
	return &Snippet{suffix: suffix}
}

type Snippet struct {
	cache        *Cache
	todo         bool
	path         string
	name         string
	suffix       string
	namespace    string
	tag          string
	tmpBlocks    []*Block
	tmpConstants []*Block
	blocks       []*Block
	constants    []*Block
	render       RenderFunc
	formatter    FormatterFunc
}

func (p *Snippet) initCache() {
	if p.cache == nil {
		p.cache = NewCache()
	}
}

func (p *Snippet) Commit() {

	filePath := path.Join(p.path, p.name+p.suffix)
	p.getFileCache().Add(filePath)

	for _, v := range p.tmpConstants {
		if v.filter != nil && p.getFileCache().Match(filePath, v.filter.GetRule()) {
			continue
		}
		p.constants = append(p.constants, v)
	}
	p.tmpConstants = make([]*Block, 0)

	for _, v := range p.tmpBlocks {
		if v.filter != nil && p.getFileCache().Match(filePath, v.filter.GetRule()) {
			continue
		}
		p.blocks = append(p.blocks, v)
	}
	p.tmpBlocks = make([]*Block, 0)
}

func (p *Snippet) getFileCache() *Cache {
	if p.cache == nil {
		p.cache = new(Cache)
	}
	return p.cache
}

func (p *Snippet) SetTag(tag string) {
	p.tag = tag
}

func (p *Snippet) GetTag() string {
	return p.tag
}

func (p *Snippet) AddConst(constants ...*Block) {
	for _, v := range constants {
		p.tmpConstants = append(p.tmpConstants, v)
	}
}

func (p *Snippet) GetConsts() (constants []*Block) {
	for _, v := range p.constants {
		constants = append(constants, v)
	}
	return
}

func (p *Snippet) SetTodo() {
	p.todo = true
}

func (p *Snippet) SetFileName(name string) {
	p.name = name
}

func (p *Snippet) getFileName() string {
	return p.name
}

func (p *Snippet) SetSuffix(suffix string) {
	p.suffix = suffix
}

func (p *Snippet) getSuffix() string {
	return p.suffix
}

func (p *Snippet) SetFilePath(path string) {

	p.path = path
}

func (p *Snippet) getFilePath() string {
	return p.path

}

func (p *Snippet) Merge(snippets ...*Snippet) {
	for _, v := range snippets {
		p.AddConst(v.GetConsts()...)
		p.AddBlock(v.GetBlocks()...)
		p.Commit()
	}
}

func (p *Snippet) SetNamespace(namespace string) {
	p.namespace = namespace
}

func (p *Snippet) GetNamespace() string {
	return p.namespace
}

func (p *Snippet) GetPackages() (packages []*Package) {
	for _, v := range p.blocks {
		packages = append(packages, v.GetPackages()...)
	}
	return
}

func (p *Snippet) AddBlock(blocks ...*Block) {
	for _, v := range blocks {
		p.tmpBlocks = append(p.tmpBlocks, v)
	}
}

func (p *Snippet) GetBlocks() (snippetBlocks []*Block) {
	for _, v := range p.blocks {
		snippetBlocks = append(snippetBlocks, v)
	}
	return
}

func (p *Snippet) SetRender(render RenderFunc, formatter FormatterFunc) {
	p.render = render
	p.formatter = formatter
}
