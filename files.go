package snippet

func NewFiles() *Files {
	return &Files{}
}

type Files struct {
	fileMap  map[string]bool
	fileList []*File
}

func (p *Files) Add(files ...*File) {

	if p.fileMap == nil {
		p.fileMap = make(map[string]bool)
	}

	for _, v := range files {
		if _, ok := p.fileMap[v.fullPath()]; !ok {
			p.fileList = append(p.fileList, v)
			p.fileMap[v.fullPath()] = true
		}
	}
}

func (p *Files) All() []*File {
	return p.fileList
}
