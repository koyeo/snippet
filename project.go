package snippet

func NewProject() *Project {
	return &Project{}
}

type Project struct {
	root           string
	makeFileSuffix *Collection
	makeFilePrefix *Collection
	makeDirSuffix  *Collection
	makeDirPrefix  *Collection
	ignore         map[string]bool
	files          *Files
	folders        *Folders
	snippets       *Snippets
	writer         *Writer
}

func (p *Project) initFiles() {
	if p.files == nil {
		p.files = NewFiles()
	}
}

func (p *Project) initFolders() {
	if p.folders == nil {
		p.folders = NewFolders()
	}
}

func (p *Project) initSnippets() {
	if p.snippets == nil {
		p.snippets = NewSnippets()
	}
}

func (p *Project) initMakeFilePrefix() {
	if p.makeFilePrefix == nil {
		p.makeFilePrefix = NewCollection()
	}
}

func (p *Project) initMakeDirPrefix() {
	if p.makeDirPrefix == nil {
		p.makeDirPrefix = NewCollection()
	}
}

func (p *Project) initMakeFileSuffix() {
	if p.makeFileSuffix == nil {
		p.makeFileSuffix = NewCollection()
	}
}

func (p *Project) initMakeDirSuffix() {
	if p.makeDirSuffix == nil {
		p.makeDirSuffix = NewCollection()
	}
}

func (p *Project) initWriter() {
	if p.writer == nil {
		p.writer = NewWriter()
	}
}

func (p *Project) SetRoot(root string) {
	p.root = root
}

func (p *Project) AddFile(files ...*File) {
	p.initFiles()
	p.files.Add(files...)
}

func (p *Project) AddFolder(folders ...*Folder) {
	p.initFolders()
	p.folders.Add(folders...)
}

func (p *Project) AddSnippet(snippets ...*Snippet) {
	p.initSnippets()
	p.snippets.Add(snippets...)
}

func (p *Project) Render() {
	p.initWriter()
	p.collectMakePrefixAndSuffix()
	p.loadLocalFiles()
	p.loadLocalDirs()
	p.renderSnippets()
	p.renderFiles()
	p.renderFolders()
	p.render()
	p.clean()
}

func (p *Project) collectMakePrefixAndSuffix() {

	p.initMakeFilePrefix()
	p.initMakeFileSuffix()
	p.initMakeDirPrefix()
	p.initMakeDirSuffix()

	p.initSnippets()
	for _, v := range p.snippets.All() {
		p.makeFilePrefix.Add(v.makePrefix)
		p.makeFileSuffix.Add(v.makeSuffix + v.suffix)
	}

	p.initFiles()
	for _, v := range p.files.All() {
		p.makeFilePrefix.Add(v.makePrefix)
		p.makeFileSuffix.Add(v.makeSuffix + v.suffix)
	}

	p.initFolders()
	for _, v := range p.folders.All() {
		v.initFiles()
		for _, vv := range v.files.All() {
			p.makeFilePrefix.Add(vv.makePrefix)
			p.makeFileSuffix.Add(vv.makeSuffix + vv.suffix)
		}
		p.makeDirPrefix.Add(v.makePrefix)
		p.makeDirSuffix.Add(v.makeSuffix)
	}
}

func (p *Project) loadLocalFiles() {

	p.initMakeFilePrefix()
	p.initMakeFileSuffix()

	err := p.writer.loadLocalRenderFiles(p.root, p.makeFilePrefix.All(), p.makeFileSuffix.All())
	if err != nil {
		Fatal("Load local render files error: ", err)
	}
}

func (p *Project) loadLocalDirs() {

	p.initMakeDirPrefix()
	p.initMakeDirSuffix()

	err := p.writer.loadLocalRenderDirs(p.root, p.makeFilePrefix.All(), p.makeFileSuffix.All())
	if err != nil {
		Fatal("Load local render dirs error: ", err)
	}
}

func (p *Project) renderSnippets() {
	p.initSnippets()
	p.snippets.render(p)
}

func (p *Project) renderFiles() {
	p.initFiles()
	p.files.render(p, p.root)
}

func (p *Project) renderFolders() {
	p.initFolders()
	p.folders.render(p)
}

func (p *Project) render() {
	p.writer.render()
}

func (p *Project) clean() {
	p.writer.clean()
	MakeDone()
}
