package snippet

func NewFolders() *Folders {
	return &Folders{}
}

type Folders struct {
	folderMap  map[string]bool
	folderList []*Folder
}

func (p *Folders) Add(folders ...*Folder) {

	if p.folderMap == nil {
		p.folderMap = make(map[string]bool)
	}

	for _, v := range folders {
		if _, ok := p.folderMap[v.path]; !ok {
			p.folderList = append(p.folderList, v)
			p.folderMap[v.path] = true
		}
	}
}

func (p *Folders) All() []*Folder {
	return p.folderList
}
