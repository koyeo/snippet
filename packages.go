package snippet

func NewPackages() *Packages {
	return &Packages{}
}

type Packages struct {
	packageList []*Package
	packageMap  map[string]bool
}

func (p *Packages) Add(packages ...*Package) {

	if p.packageMap == nil {
		p.packageMap = make(map[string]bool)
	}

	for _, v := range packages {
		if _, ok := p.packageMap[v.Path()]; !ok {
			p.packageList = append(p.packageList, v)
			p.packageMap[v.Path()] = true
		}
	}
}

func (p *Packages) All() []*Package {
	return p.packageList
}
