package snippet

func NewBlock(filter Filter, code string, data interface{}) *Block {
	b := &Block{
		filter: filter,
	}
	b.RenderCode(code, data)
	return b
}

func NewStaticBlock(filter Filter, code string) *Block {
	return NewBlock(filter, code, nil)
}

func NewDocument(document string) *Block {
	b := &Block{
		filter: nil,
	}
	b.RenderCode(document, nil)
	return b
}

type Block struct {
	filter   Filter
	code     string
	packages *Packages
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

func (p *Block) GetPackages() []*Package {
	p.initPackages()
	return p.packages.All()
}

func (p *Block) SetFilter(filter Filter) {
	p.filter = filter
}

func (p *Block) GetFilter() (filter Filter) {
	return p.filter
}

func (p *Block) RenderCode(code string, data interface{}) {

	code, err := Render(code, data)
	if err != nil {
		Fatal("Render code error: ", err)
	}
	p.code = code
	return
}

func (p *Block) GetCode() string {
	return p.code
}
