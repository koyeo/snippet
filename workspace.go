package snippet

func NewWorkspace() *Workspace {
	return &Workspace{}
}

type Workspace struct {
	root           string
	filterPaths    *Collection
	makeFileSuffix *Collection
	makeFilePrefix *Collection
	makeDirSuffix  *Collection
	makeDirPrefix  *Collection
	ignorePaths    *Collection
	files          *Files
	folders        *Folders
	snippets       *Snippets
}

func (p *Workspace) initFiles() {
	if p.files == nil {
		p.files = NewFiles()
	}
}

func (p *Workspace) initFolders() {
	if p.folders == nil {
		p.folders = NewFolders()
	}
}

func (p *Workspace) initSnippets() {
	if p.snippets == nil {
		p.snippets = NewSnippets()
	}
}

func (p *Workspace) initFilterPaths() {
	if p.filterPaths == nil {
		p.filterPaths = NewCollection()
	}
}

func (p *Workspace) initIgnorePaths() {
	if p.ignorePaths == nil {
		p.ignorePaths = NewCollection()
	}
}

func (p *Workspace) initMakeFilePrefix() {
	if p.makeFilePrefix == nil {
		p.makeFilePrefix = NewCollection()
	}
}

func (p *Workspace) initMakeDirPrefix() {
	if p.makeDirPrefix == nil {
		p.makeDirPrefix = NewCollection()
	}
}

func (p *Workspace) initMakeFileSuffix() {
	if p.makeFileSuffix == nil {
		p.makeFileSuffix = NewCollection()
	}
}

func (p *Workspace) initMakeDirSuffix() {
	if p.makeDirSuffix == nil {
		p.makeDirSuffix = NewCollection()
	}
}

func (p *Workspace) SetRoot(root string, filter bool) {
	p.root = root
	if filter{
		p.filterRoot()
	}
}

func (p *Workspace) AddFilterPath(paths ...string) {
	p.initFilterPaths()
	p.filterPaths.Add(paths...)
}

func (p *Workspace) AddIgnorePath(paths ...string) {
	p.initIgnorePaths()
	p.ignorePaths.Add(paths...)
}

func (p *Workspace) AddFile(files ...*File) {
	p.initFiles()
	p.files.Add(files...)
}

func (p *Workspace) AddFolder(folders ...*Folder) {
	p.initFolders()
	p.folders.Add(folders...)
}

func (p *Workspace) AddSnippet(snippets ...*Snippet) {
	p.initSnippets()
	p.snippets.Add(snippets...)
}

func (p *Workspace) render(project *Project) {
	p.collectMakePrefixAndSuffix()
	p.renderSnippets(project)
	p.renderFiles(project)
	p.renderFolders(project)
}

func (p *Workspace) collectMakePrefixAndSuffix() {

	p.initMakeFilePrefix()
	p.initMakeFileSuffix()
	p.initMakeDirPrefix()
	p.initMakeDirSuffix()

	p.initSnippets()
	for _, v := range p.snippets.All() {
		if v.makePrefix != "" {
			p.makeFilePrefix.Add(v.makePrefix)
		}
		if v.makeSuffix != "" {
			p.makeFileSuffix.Add(v.makeSuffix + v.suffix)
		}
	}

	p.initFiles()
	for _, v := range p.files.All() {
		if v.makePrefix != "" {
			p.makeFilePrefix.Add(v.makePrefix)
		}
		if v.makeSuffix != "" {
			p.makeFileSuffix.Add(v.makeSuffix + v.suffix)
		}
	}

	p.initFolders()
	for _, v := range p.folders.All() {
		v.initFiles()
		for _, vv := range v.files.All() {
			if vv.makePrefix != "" {
				p.makeFilePrefix.Add(vv.makePrefix)
			}
			if vv.makeSuffix != "" {
				p.makeFileSuffix.Add(vv.makeSuffix + vv.suffix)
			}
		}
		if v.makePrefix != "" {
			p.makeDirPrefix.Add(v.makePrefix)
		}
		if v.makeSuffix != "" {
			p.makeDirSuffix.Add(v.makeSuffix)
		}
	}
}

func (p *Workspace) filterRoot() {
	p.initFilterPaths()
	p.filterPaths.Add(p.root)
}

func (p *Workspace) renderSnippets(project *Project) {
	p.initSnippets()
	p.snippets.render(project, p)
}

func (p *Workspace) renderFiles(project *Project) {
	p.initFiles()
	p.files.render(project, p.root)
}

func (p *Workspace) renderFolders(project *Project) {
	p.initFolders()
	p.folders.render(project, p)
}
