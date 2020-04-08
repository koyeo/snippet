package snippet

import (
	"path/filepath"
)

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
		if _, ok := p.folderMap[v.name]; !ok {
			p.folderList = append(p.folderList, v)
			p.folderMap[v.name] = true
		}
	}
}

func (p *Folders) All() []*Folder {
	return p.folderList
}

func (p *Folders) render(project *Project) {
	for _, v := range p.folderList {

		distPath := filepath.Join(project.root, v.dir, v.makePrefix+v.name+v.makeSuffix)
		customPath := filepath.Join(project.root, v.name)

		v.initFiles()
		v.files.render(project, distPath)

		project.writer.AddMakeRenderDir(distPath, customPath, len(v.files.All()))
	}
}
