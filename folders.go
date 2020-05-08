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

func (p *Folders) render(project *Project, workspace *Workspace) {

	for _, v := range p.folderList {
		var distPath string
		var customPath string

		distPath = filepath.Join(workspace.root, v.dir, v.makePrefix+v.name+v.makeSuffix)
		customPath = filepath.Join(workspace.root, v.dir, v.name)

		v.initFiles()
		v.files.render(project, distPath)

		project.writer.addMakeRenderDir(distPath, customPath, len(v.files.All()))
	}
}
