package snippet

func NewFolder() *Folder {
	return &Folder{}
}

type Folder struct {
	name       string
	makePrefix string
	makeSuffix string
	files      *Files
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
