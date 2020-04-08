package snippet

import (
	"github.com/koyeo/snippet/logger"
	"github.com/koyeo/snippet/writer"
)

func NewBlock(filter Filter, code string, data interface{}) *Block {
	b := &Block{
		filter: filter,
	}
	b.SetCode(code, data)
	return b
}

func NewStaticBlock(filter Filter, code string) *Block {
	return NewBlock(filter, code, nil)
}

func NewDocument(document string) *Block {
	b := &Block{
		filter: nil,
	}
	b.SetCode(document, nil)
	return b
}

type Block struct {
	filter   Filter
	code     string
	packages *Packages
	exist    bool
}

func (p *Block) SetExist(exist bool) {
	p.exist = exist
}

func (p *Block) Exist() bool {
	return p.exist
}

func (p *Block) initPackages() {
	if p.packages == nil {
		p.packages = NewPackages()
	}
}

func (p *Block) UsePackage(packages ...*Package) {
	p.initPackages()
	p.packages.Add(packages...)
}

func (p *Block) Packages() []*Package {
	p.initPackages()
	return p.packages.All()
}

func (p *Block) SetFilter(filter Filter) {
	p.filter = filter
}

func (p *Block) Filter() (filter Filter) {
	return p.filter
}

func (p *Block) SetCode(code string, data interface{}) {

	code, err := writer.Render(code, data)
	if err != nil {
		logger.Fatal("Render content error: ", err)
	}
	p.code = code
	return
}

func (p *Block) Code() string {
	return p.code
}
