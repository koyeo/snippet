package snippet

import (
	"github.com/flosch/pongo2"
	"path"
)

type RenderFunc func(ctx pongo2.Context, snippet *Snippet) (content string, err error)
type FormatterFunc func(content string) (string, error)

func NewSnippet(suffix string) *Snippet {
	return &Snippet{suffix: suffix, trimSpace: true}
}

type Snippet struct {
	cache        *Cache
	todo         bool
	path         string
	name         string
	suffix       string
	makePrefix   string
	makeSuffix   string
	namespace    string
	tags         []string
	tmpBlocks    []*Block
	tmpConstants []*Block
	blocks       *Blocks
	constants    *Blocks
	render       RenderFunc
	formatter    FormatterFunc
	ignore       bool
	trimSpace    bool
}

func (p *Snippet) SetTrimSpace(trimSpace bool) {
	p.trimSpace = trimSpace
}

func (p *Snippet) SetIgnore(ignore bool) {
	p.ignore = ignore
}

func (p *Snippet) initConstants() {
	if p.constants == nil {
		p.constants = NewBlocks()
	}
}

func (p *Snippet) Constants() []*Block {
	p.initConstants()
	return p.constants.All()
}

func (p *Snippet) initBlocks() {
	if p.blocks == nil {
		p.blocks = NewBlocks()
	}
}

func (p *Snippet) Blocks() []*Block {
	p.initBlocks()
	return p.blocks.All()
}

func (p *Snippet) SetMakeSuffix(suffix string) {
	p.makeSuffix = suffix
}

func (p *Snippet) SetMakePrefix(prefix string) {
	p.makePrefix = prefix
}

func (p *Snippet) initCache() {
	if p.cache == nil {
		p.cache = NewCache()
	}
}

func (p *Snippet) Commit() {

	filePath := path.Join(p.path, p.name+p.suffix)
	p.Cache().Add(filePath)

	p.initConstants()
	p.initBlocks()
	for _, v := range p.tmpConstants {
		if v.filter != nil && p.Cache().Match(filePath, v.filter.Rule()) {
			continue
		}
		p.constants.Add(v)
	}
	p.tmpConstants = make([]*Block, 0)

	for _, v := range p.tmpBlocks {
		if v.filter != nil && p.Cache().Match(filePath, v.filter.Rule()) {
			continue
		}
		p.blocks.Add(v)
	}
	p.tmpBlocks = make([]*Block, 0)
}

func (p *Snippet) Cache() *Cache {
	if p.cache == nil {
		p.cache = new(Cache)
	}
	return p.cache
}

func (p *Snippet) AddTag(tags ...string) {
	p.tags = append(p.tags, tags...)
}

func (p *Snippet) GetTags() []string {
	return p.tags
}

func (p *Snippet) AddConstant(constants ...*Block) {
	for _, v := range constants {
		p.tmpConstants = append(p.tmpConstants, v)
	}
}

func (p *Snippet) SetTodo() {
	p.todo = true
}

func (p *Snippet) SetName(name string) {
	p.name = name
}

func (p *Snippet) Name() string {
	return p.name
}

func (p *Snippet) SetSuffix(suffix string) {
	p.suffix = suffix
}

func (p *Snippet) Suffix() string {
	return p.suffix
}

func (p *Snippet) SetDir(path string) {

	p.path = path
}

func (p *Snippet) getDir() string {
	return p.path

}

func (p *Snippet) Merge(snippets ...*Snippet) {
	for _, v := range snippets {
		p.AddConstant(v.Constants()...)
		p.AddBlock(v.Blocks()...)
		p.Commit()
	}
}

func (p *Snippet) SetNamespace(namespace string) {
	p.namespace = namespace
}

func (p *Snippet) Namespace() string {
	return p.namespace
}

func (p *Snippet) Packages() (packages []*Package) {
	p.initBlocks()
	for _, v := range p.blocks.All() {
		if v.exist {
			continue
		}
		packages = append(packages, v.Packages()...)
	}
	return
}

func (p *Snippet) AddBlock(blocks ...*Block) {
	for _, v := range blocks {
		p.tmpBlocks = append(p.tmpBlocks, v)
	}
}

func (p *Snippet) SetRender(render RenderFunc) {
	p.render = render

}

func (p *Snippet) SetFormatter(formatter FormatterFunc) {
	p.formatter = formatter
}
