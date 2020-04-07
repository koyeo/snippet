package snippet

import (
	"path/filepath"
)

func NewFile() *File {
	return &File{}
}

type File struct {
	dir        string
	name       string
	suffix     string
	makePrefix string
	makeSuffix string
	code       string
}

func (p *File) fullPath() string {
	return filepath.Join(p.dir, p.name)
}

func (p *File) SetPath(path string) {
	p.dir = path
}

func (p *File) SetDir(path string) {
	p.dir = path
}

func (p *File) SetName(name string) {
	p.name = name
}

func (p *File) SetSuffix(suffix string) {
	p.suffix = suffix
}

func (p *File) SetMakeSuffix(makeSuffix string) {
	p.makeSuffix = makeSuffix
}

func (p *File) SetMakePrefix(makePrefix string) {
	p.makePrefix = makePrefix
}

func (p *File) SetContent(code string, data interface{}) {
	code, err := Render(code, data)
	if err != nil {
		Fatal("Render code error: ", err)
	}
	p.code = code
	return
}
