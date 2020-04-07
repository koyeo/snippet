package snippet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
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

func NewWriter() (c *Writer) {
	c = new(Writer)
	c.initRenderFiles()
	return
}

type Writer struct {
	renderFiles map[string]*RenderFile
	renderDirs  map[string]*RenderDir
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

func (p *Writer) loadLocalRenderFiles(path string, prefix []string, suffix []string) (err error) {

	if len(prefix) == 0 && len(suffix) == 0 {
		return
	}

	if !PathExist(path) {
		return
	}

	files, err := ReadFiles(path, prefix, suffix)
	if err != nil {
		return
	}
	for _, file := range files {

		var content string
		content, err = ReadFile(file)
		if err != nil {
			return
		}
		p.addLocalRenderFile(file, content)
	}

	return
}

func (p *Writer) loadLocalRenderDirs(path string, prefix []string, suffix []string) (err error) {

	if len(prefix) == 0 && len(suffix) == 0 {
		return
	}

	if !PathExist(path) {
		return
	}

	dirs, err := ReadDirs(path, prefix, suffix)
	if err != nil {
		return
	}

	for _, dir := range dirs {
		p.addLocalRenderDir(dir)
	}

	return
}

func (p *Writer) compare(makePath, customPath, makeContent string, makeSnippet bool) string {

	if strings.TrimSpace(makePath) == "" {
		return ""
	}

	var err error
	var compareContent string

	if makeSnippet {

		if PathExist(customPath) {
			compareContent, err = ReadFile(customPath)
			if err != nil {
				Error(fmt.Sprintf("Read %s error:", customPath), err)
				return ""
			}
		}

		makeContent = p.compareContent(makeContent, compareContent)

	} else {
		if PathExist(customPath) {
			return ""
		}
	}

	return makeContent
}

func (p *Writer) clean() {

	for filePath, renderFile := range p.renderFiles {
		if renderFile.MakeContent == "" {
			err := Remove(filePath)
			if err != nil {
				Error(fmt.Sprintf("Remove file error:"), err)
				return
			}
			CleanFileSuccess(filePath)
		}
	}

	for dirPath, renderDir := range p.renderDirs {

		if renderDir.MakeFiles == 0 || PathExist(renderDir.CustomPath) {
			err := Remove(dirPath)
			if err != nil {
				Error(fmt.Sprintf("Remove Dir error:"), err)
				return
			}
			CleanDirSuccess(dirPath)
		}
	}
}

func (p *Writer) compareContent(makeContent, compareContent string) string {

	blockRegex := regexp.MustCompile(`<block\s+rule\s*=\s*"(\S*)">([\s\S]*?)</block>`)
	res := blockRegex.FindAllStringSubmatch(makeContent, -1)

	if len(res) > 0 {
		count := 0
		for _, v := range res {
			if p.matchSegment(v[1], compareContent) {
				makeContent = strings.Replace(makeContent, v[0], "", -1)
			} else {
				count++

				makeContent = strings.Replace(makeContent, v[0], v[2], -1)

			}
		}

		if count == 0 {
			return ""
		}
	}

	newLineRegex := regexp.MustCompile(`<\s*\\n\s*>`)
	spaceRegex := regexp.MustCompile(`<\s*\\s\s*>`)
	segments := strings.Split(makeContent, "\n")
	results := make([]string, 0)

	for _, v := range segments {

		v = strings.TrimSpace(v)

		if newLineRegex.MatchString(v) {
			results = append(results, "")
			v = newLineRegex.ReplaceAllString(v, "")
		}
		if spaceRegex.MatchString(v) {
			v = spaceRegex.ReplaceAllString(v, " ")
		}

		if v != "" {
			results = append(results, v)
		}
	}

	return strings.Join(results, "\n")

}

func (p *Writer) matchSegment(rule string, content string) (match bool) {

	reg, err := regexp.Compile(rule)
	if err != nil {
		Fatal(fmt.Sprintf(`Compile regexp %s error:`, rule), err)
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
				if PathExist(makePath) {
					var err error
					renderFile.LocalContent, err = ReadFile(makePath)
					if err != nil {
						Error("CheckDiff read file error: ", err)
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
				MakeDir(Abs(renderFile.MakePath))
			}

			err := WriteFile(makePath, []byte(renderFile.MakeContent))
			if err != nil {
				return
			}

			renderFile.HasWrite = true
			MakeFileSuccess(makePath)
		}
	}
}
