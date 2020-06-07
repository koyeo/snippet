package snippet

func NewBlocks() *Blocks {
	return &Blocks{}
}

type Blocks struct {
	blockList []*Block
	blockMap  map[string]*Block
}

func (p *Blocks) Add(blocks ...*Block) {

	if p.blockMap == nil {
		p.blockMap = make(map[string]*Block)
	}

	for _, v := range blocks {
		md5 := v.md5()
		if _, ok := p.blockMap[md5]; !ok {
			p.blockList = append(p.blockList, v)
			p.blockMap[md5] = v
		}
	}
}

func (p *Blocks) All() (blocks []*Block) {
	for _, v := range p.blockList {
		blocks = append(blocks, v)
	}
	return
}
