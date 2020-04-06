package snippet

type Project struct {
	root     string
	suffix   []string
	prefix   []string
	ignore   map[string]bool
	files    *Files
	folders  *Folders
	snippets *Snippets
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

func (p *Project) SetRoot(root string) {
	p.root = root
}

func (p *Project) SetSuffix(suffix ...string) {
	p.prefix = suffix
}

func (p *Project) SetPrefix(prefix ...string) {
	p.prefix = prefix
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
