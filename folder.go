package snippet

func NewFolder() *Folder {
	return &Folder{}
}

type Folder struct {
	name       string
	dir        string
	makePrefix string
	makeSuffix string
	files      *Files
	absolutePath bool
}

func (p *Folder) SetAbsolutePath(absolutePath bool) {
	p.absolutePath = absolutePath
}

func (p *Folder) SetDir(dir string) {
	p.dir = dir
}

func (p *Folder) initFiles() {
	if p.files == nil {
		p.files = NewFiles()
	}
}

func (p *Folder) SetName(name string) {
	p.name = name
}

func (p *Folder) SetMakeSuffix(suffix string) {
	p.makeSuffix = suffix
}

func (p *Folder) SetMakePrefix(prefix string) {
	p.makePrefix = prefix
}

func (p *Folder) AddFile(file *File) {
	p.initFiles()
	p.files.Add(file)
}
