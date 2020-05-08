package snippet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/koyeo/snippet/logger"
	"github.com/koyeo/snippet/storage"
	"os"
	"regexp"
)

type RenderFile struct {
	LocalContent string
	MakeContent  string
	MakePath     string
	CustomPath   string
	HasWrite     bool
	CheckDiff    bool
}

type RenderDir struct {
	MakePath   string
	CustomPath string
	MakeFiles  int
}

type TemplateFunc func(data interface{}) string

func NewWriter() (c *Writer) {
	c = new(Writer)
	c.initRenderFiles()
	return
}

type Writer struct {
	renderFiles map[string]*RenderFile
	renderDirs  map[string]*RenderDir
	funcs       map[string]TemplateFunc
}

func (p *Writer) initRenderFiles() {
	if p.renderFiles == nil {
		p.renderFiles = make(map[string]*RenderFile)
	}
}

func (p *Writer) initRenderDirs() {
	if p.renderDirs == nil {
		p.renderDirs = make(map[string]*RenderDir)
	}
}

func (p *Writer) initFuncs() {
	if p.funcs == nil {
		p.funcs = make(map[string]TemplateFunc)
	}
}

func (p *Writer) initRenderFile(makePath string) {
	p.initRenderFiles()
	if p.renderFiles[makePath] == nil {
		p.renderFiles[makePath] = &RenderFile{
			MakePath: makePath,
		}
	}
}

func (p *Writer) initRenderDir(makePath string) {
	p.initRenderDirs()
	if p.renderDirs[makePath] == nil {
		p.renderDirs[makePath] = &RenderDir{
			MakePath: makePath,
		}
	}
}

func (p *Writer) addLocalRenderFile(filePath, content string) {
	p.initRenderFile(filePath)
	p.renderFiles[filePath].LocalContent = content
}

func (p *Writer) addLocalRenderDir(dirPath string) {
	p.initRenderDirs()
	p.renderDirs[dirPath] = &RenderDir{
		MakePath: dirPath,
	}
}

func (p *Writer) addMakeRenderFile(distPath, makePath, customPath, makeContent string, makeDiff bool) {
	p.initRenderFile(makePath)
	p.renderFiles[makePath].MakePath = distPath
	p.renderFiles[makePath].CustomPath = customPath
	p.renderFiles[makePath].MakeContent = makeContent
	p.renderFiles[makePath].CheckDiff = makeDiff
}

func (p *Writer) addMakeRenderDir(distPath, customPath string, makeFiles int) {
	p.initRenderDir(distPath)
	p.renderDirs[distPath].CustomPath = customPath
	p.renderDirs[distPath].MakeFiles = makeFiles
}

func (p *Writer) addFunc(name string, _func TemplateFunc) {
	p.initFuncs()
	p.funcs[name] = _func
}

func (p *Writer) loadLocalRenderFiles(debug bool, paths []string, ignore, prefix, suffix []string) (err error) {

	if len(prefix) == 0 && len(suffix) == 0 {
		return
	}

	for _, path := range paths {

		if !storage.PathExist(path) {
			return
		}

		files, err := storage.ReadFiles(debug, path, ignore, prefix, suffix)
		if err != nil {
			return err
		}
		for _, file := range files {
			var content string
			content, err = storage.ReadFile(file)
			if err != nil {
				return err
			}
			p.addLocalRenderFile(file, content)
		}
	}

	return
}

func (p *Writer) loadLocalRenderDirs(debug bool, paths []string, ignore, prefix, suffix []string) (err error) {

	if len(prefix) == 0 && len(suffix) == 0 {
		return
	}

	for _, path := range paths {
		if !storage.PathExist(path) {
			return
		}

		dirs, err := storage.ReadDirs(debug, path, ignore, prefix, suffix)
		if err != nil {
			return err
		}

		for _, dir := range dirs {
			p.addLocalRenderDir(dir)
		}
	}

	return
}

func (p *Writer) compareSnippet(snippet *Snippet, customPath string) {

	var err error
	var compareContent string

	if storage.PathExist(customPath) {
		compareContent, err = storage.ReadFile(customPath)
		if err != nil {
			logger.Fatal(fmt.Sprintf("Read %s error:", customPath), err)
		}
	}

	items := append(snippet.Constants(), snippet.Blocks()...)
	i := 0
	for _, v := range items {
		if v.Filter() != nil {
			if p.matchSegment(v.Filter().Rule(), compareContent) {
				v.SetExist(true)
				i++
				continue
			}
		}
	}
	if i == len(items) {
		snippet.SetIgnore(true)
	}
}

func (p *Writer) clean(makeDirs *Collection) {

	for filePath, renderFile := range p.renderFiles {
		if !makeDirs.Has(renderFile.MakePath) {
			if renderFile.MakeContent == "" && storage.PathExist(filePath) && storage.PathExist(renderFile.CustomPath) {
				err := storage.Remove(filePath)
				if err != nil {
					logger.Error(fmt.Sprintf("Remove file error:"), err)
					return
				}
				logger.CleanFileSuccess(filePath)
			}
		}
	}

	for dirPath, renderDir := range p.renderDirs {
		if makeDirs.Has(dirPath) && storage.PathExist(renderDir.CustomPath) && storage.PathExist(dirPath) {
			err := storage.Remove(dirPath)
			if err != nil {
				logger.Error(fmt.Sprintf("Remove Dir error:"), err)
				return
			}
			logger.CleanDirSuccess(dirPath)
		}
	}
}

func (p *Writer) matchSegment(rule string, content string) (match bool) {

	reg, err := regexp.Compile(rule)
	if err != nil {
		logger.Fatal(fmt.Sprintf(`Compile regexp %s error:`, rule), err)
	}

	match = reg.MatchString(content)

	return
}

func (p *Writer) render() {

	for makePath, renderFile := range p.renderFiles {
		render := true

		if renderFile.MakeContent == "" {
			render = false
		}

		if renderFile.CheckDiff {

			if renderFile.LocalContent == "" {
				if storage.PathExist(makePath) {
					var err error
					renderFile.LocalContent, err = storage.ReadFile(makePath)
					if err != nil {
						logger.Error("CheckDiff read file error: ", err)
						os.Exit(1)
					}
				}
			}

			m1 := md5.New()
			m2 := md5.New()

			m1.Write([]byte(renderFile.LocalContent))
			m2.Write([]byte(renderFile.MakeContent))

			hash1 := hex.EncodeToString(m1.Sum(nil))
			hash2 := hex.EncodeToString(m2.Sum(nil))

			if hash1 == hash2 {
				render = false
			}
		}

		if render && !renderFile.HasWrite {

			if renderFile.MakePath != "" {
				storage.MakeDir(storage.Abs(renderFile.MakePath))
			}

			err := storage.WriteFile(makePath, []byte(renderFile.MakeContent))
			if err != nil {
				return
			}

			renderFile.HasWrite = true
			logger.MakeFileSuccess(makePath)
		}
	}
}
