package snippet

import (
	"path/filepath"
)

func NewFile() *File {
	return &File{
		trimSpace: true,
	}
}

type File struct {
	dir        string
	name       string
	suffix     string
	makePrefix string
	makeSuffix string
	content    string
	formatter  FormatterFunc
	data       interface{}
	trimSpace  bool
}

func (p *File) Formatter() FormatterFunc {
	return p.formatter
}

func (p *File) SetFormatter(formatter FormatterFunc) {
	p.formatter = formatter
}

func (p *File) SetTrimSpace(trimSpace bool) {
	p.trimSpace = trimSpace
}

func (p *File) fullPath() string {
	return filepath.Join(p.dir, p.name, p.suffix)
}

func (p *File) SetDir(dir string) {
	p.dir = dir
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

func (p *File) SetContent(content string, data interface{}) {
	p.content, p.data = content, data
}
