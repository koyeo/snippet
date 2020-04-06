package snippet

type Package struct {
	name string
	path string
}

func NewPackage(name string, path string) *Package {
	return &Package{name: name, path: path}
}

func (p *Package) SetName(name string) {
	p.name = name
}

func (p *Package) GetName() (name string) {
	name = p.name
	return
}

func (p *Package) SetPath(path string) {
	p.path = path
}

func (p *Package) GetPath() (path string) {
	path = p.path
	return
}

