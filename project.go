package snippet

import (
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet/logger"
)

func NewProject() *Project {
	return &Project{}
}

type Project struct {
	filterPaths    *Collection
	ignorePaths    *Collection
	makeFileSuffix *Collection
	makeFilePrefix *Collection
	makeDirSuffix  *Collection
	makeDirPrefix  *Collection
	workspaces     []*Workspace
	writer         *Writer
	debug          bool
	ctx            pongo2.Context
	hideMakeDone   bool
}

func (p *Project) initFilterPaths() {
	if p.filterPaths == nil {
		p.filterPaths = NewCollection()
	}
}

func (p *Project) initIgnorePaths() {
	if p.ignorePaths == nil {
		p.ignorePaths = NewCollection()
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

func (p *Project) AddWorkspace(workspace ...*Workspace) {
	p.workspaces = append(p.workspaces, workspace...)
}

func (p *Project) HideMakeDone() {
	p.hideMakeDone = true
}

func (p *Project) SetCtx(ctx pongo2.Context) {
	p.ctx = ctx
}

func (p *Project) SetDebug() {
	p.debug = true
}

func (p *Project) initWriter() {
	if p.writer == nil {
		p.writer = NewWriter()
	}
}

func (p *Project) Render() {
	p.initWriter()
	p.collectMakePrefixAndSuffix()
	p.loadLocalFiles()
	p.render()
	p.clean()
}

func (p *Project) loadLocalFiles() {

	p.initFilterPaths()
	p.initIgnorePaths()
	p.initMakeFilePrefix()
	p.initMakeFileSuffix()

	err := p.writer.loadLocalRenderFiles(
		p.debug,
		p.filterPaths.All(),
		p.ignorePaths.All(),
		p.makeFilePrefix.All(),
		p.makeFileSuffix.All(),
	)
	if err != nil {
		logger.Fatal("Load local render files error: ", err)
	}
}

func (p *Project) collectMakePrefixAndSuffix() {

	p.initFilterPaths()
	p.initIgnorePaths()
	p.initMakeDirPrefix()
	p.initMakeDirSuffix()
	p.initMakeFilePrefix()
	p.initMakeFileSuffix()

	for _, v := range p.workspaces {
		v.collectMakePrefixAndSuffix()
		if v.filterPaths != nil {
			p.filterPaths.Add(v.filterPaths.All()...)
		}
		if v.ignorePaths != nil {
			p.ignorePaths.Add(v.ignorePaths.All()...)
		}
		if v.makeDirPrefix != nil {
			p.makeDirPrefix.Add(v.makeDirPrefix.All()...)
		}
		if v.makeDirSuffix != nil {
			p.makeDirSuffix.Add(v.makeDirSuffix.All()...)
		}
		if v.makeFilePrefix != nil {
			p.makeFilePrefix.Add(v.makeFilePrefix.All()...)
		}
		if v.makeFileSuffix != nil {
			p.makeFileSuffix.Add(v.makeFileSuffix.All()...)
		}
	}
}

func (p *Project) render() {
	for _, v := range p.workspaces {
		v.render(p)
	}
	p.writer.render()
}

func (p *Project) clean() {
	p.writer.clean()
	if !p.hideMakeDone {
		logger.MakeDone()
	}
}
