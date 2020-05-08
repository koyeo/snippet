package snippet

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet/logger"
	"path/filepath"
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
	ignoreMakeDirs *Collection
	makeDirs       *Collection
	workspaces     []*Workspace
	writer         *Writer
	debug          bool
	ctx            pongo2.Context
	hideMakeDone   bool
}

func (p *Project) initMakeDirs() {
	if p.makeDirs == nil {
		p.makeDirs = NewCollection()
	}
}

func (p *Project) initIgnoreMakeDirs() {
	if p.ignoreMakeDirs == nil {
		p.ignoreMakeDirs = NewCollection()
	}
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
	p.loadLocalDirs()
	p.render()
	p.clean()
}

func (p *Project) loadLocalFiles() {

	p.initFilterPaths()
	p.initIgnorePaths()
	p.initMakeFilePrefix()
	p.initMakeFileSuffix()

	ignorePaths := p.ignorePaths.All()
	ignorePaths = append(ignorePaths, p.ignoreMakeDirs.All()...)
	err := p.writer.loadLocalRenderFiles(
		p.debug,
		p.filterPaths.All(),
		ignorePaths,
		p.makeFilePrefix.All(),
		p.makeFileSuffix.All(),
	)
	if err != nil {
		logger.Fatal("Load local render files error: ", err)
	}
}

func (p *Project) loadLocalDirs() {

	p.initFilterPaths()
	p.initIgnorePaths()
	p.initMakeDirPrefix()
	p.initMakeDirSuffix()

	err := p.writer.loadLocalRenderDirs(
		p.debug,
		p.filterPaths.All(),
		p.ignorePaths.All(),
		p.makeDirPrefix.All(),
		p.makeDirSuffix.All(),
	)
	if err != nil {
		logger.Fatal("Load local render dirs error: ", err)
	}
}

func (p *Project) collectMakePrefixAndSuffix() {

	p.initFilterPaths()
	p.initIgnorePaths()
	p.initMakeDirPrefix()
	p.initMakeDirSuffix()
	p.initMakeFilePrefix()
	p.initMakeFileSuffix()
	p.initMakeDirs()
	p.initIgnoreMakeDirs()

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
		if v.folders != nil {
			for _, vv := range v.folders.All() {
				p.ignoreMakeDirs.Add(filepath.Join(vv.dir, fmt.Sprintf("%s%s%s", vv.makePrefix, vv.name, vv.makeSuffix)))
				p.ignoreMakeDirs.Add(filepath.Join(vv.dir, vv.name))
				p.makeDirs.Add(filepath.Join(v.root, vv.dir, fmt.Sprintf("%s%s%s", vv.makePrefix, vv.name, vv.makeSuffix)))
				p.makeDirs.Add(filepath.Join(v.root, vv.dir, vv.name))
			}
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
	p.writer.clean(p.makeDirs)
	if !p.hideMakeDone {
		logger.MakeDone()
	}
}
