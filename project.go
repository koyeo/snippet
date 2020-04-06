package snippet

type Project struct {
	root     string
	suffix   []string
	prefix   []string
	ignore   map[string]bool
	files    *Files
	folders  *Folders
	snippets *Snippets
	writer   *Writer
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

func (p *Project) initWriter() {
	if p.writer == nil {
		p.writer = NewWriter()
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

func (p *Project) Render() {

	p.initWriter()
	p.initSnippets()

	err := p.writer.LoadLocalRenderFiles(p.root, p.prefix, p.suffix)
	if err != nil {
		Fatal("Load local make files error: ", err)
	}

	for _, v := range p.snippets.All() {
		content, suffix, err := v.render(v)
		if err != nil {
			Error("Render error: ", err)
		}
		content, err = v.formatter(content)
		if err != nil {
			Fatal("Format error: ", err)
		}

		distPath := Join(v.getFilePath())
		distFile := MakeFileName(v.getFileName(), suffix, v.getSuffix())

		content = p.writer.Compare(distPath, distFile, content, true)
		if content != "" {
			p.writer.AddMakeRenderFile(distPath, distFile, content, true)
		}
	}

	p.writer.Clean()
}
