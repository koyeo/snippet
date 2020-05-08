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

func (p *Collection) Add(items ...string) {
	p.initMap()
	for _, v := range items {
		if v == "" {
			continue
		}
		if _, ok := p._map[v]; ok {
			continue
		}
		p._map[v] = true
		p.list = append(p.list, v)
	}
}

func (p *Collection) All() []string {
	return p.list
}
