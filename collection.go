package snippet

type Collection struct {
	_map map[string]bool
	list []string
}

func NewCollection() *Collection {
	return &Collection{}
}

func (p *Collection) initMap() {
	if p._map == nil {
		p._map = make(map[string]bool)
	}
}

func (p *Collection) Add(v string) {
	p.initMap()
	if _, ok := p._map[v]; ok {
		return
	}

	p._map[v] = true
	p.list = append(p.list, v)
}

func (p *Collection) All() []string {
	return p.list
}
