package snippet

import (
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet/logger"
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
	content    string
}

func (p *File) fullPath() string {
	return filepath.Join(p.dir, p.name)
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

func (p *File) SetContent(ctx pongo2.Context, content string, data interface{}) {
	content, err := Render(ctx, content, data)
	if err != nil {
		logger.Fatal("render content error: ", err)
	}
	p.content = content
	return
}
