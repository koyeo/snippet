package snippet

import (
	"path/filepath"
)

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

func (p *Files) render(project *Project, root string) {

	for _, v := range p.fileList {
		project.initMakeFilePrefix()
		project.makeFilePrefix.Add(v.makePrefix)

		project.initMakeFileSuffix()
		project.makeFileSuffix.Add(v.makeSuffix)

		distFile := v.makePrefix + v.name + v.makeSuffix + v.suffix
		customFile := v.name + v.suffix

		distPath := filepath.Join(root, v.dir)
		makePath := filepath.Join(distPath, distFile)
		customPath := filepath.Join(distPath, customFile)

		content := project.writer.compare(makePath, customPath, v.code, false, false, )
		if content != "" && !PathExist(customPath) {
			project.writer.addMakeRenderFile(distPath, makePath, customPath, content, true)
		}
	}
}
