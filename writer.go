package snippet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type RenderFile struct {
	Local string
	Make  string
	Path  string
	Write bool
	Diff  bool
}

func NewWriter() (c *Writer) {
	c = new(Writer)
	c.initRenderFiles()
	return
}

type Writer struct {
	renderFiles map[string]*RenderFile
}

func (p *Writer) initRenderFiles() {
	if p.renderFiles == nil {
		p.renderFiles = make(map[string]*RenderFile)
	}
}

func (p *Writer) initRenderFile(filePath string) {
	p.initRenderFiles()
	if p.renderFiles[filePath] == nil {
		p.renderFiles[filePath] = new(RenderFile)
	}
}

func (p *Writer) addLocalRenderFile(filePath, content string) {
	p.initRenderFile(filePath)
	p.renderFiles[filePath].Local = content
}

func (p *Writer) AddMakeRenderFile(distPath, distFile, makeContent string, makeDiff bool) {
	filePath := filepath.Join(distPath, distFile)
	p.initRenderFile(filePath)
	p.renderFiles[filePath].Path = distPath
	p.renderFiles[filePath].Diff = makeDiff
	p.renderFiles[filePath].Make = makeContent
}

func (p *Writer) LoadLocalRenderFiles(path string, prefix []string, suffix []string) (err error) {

	files, err := ReadFiles(path, prefix, suffix)
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

func (p *Writer) Compare(distPath, distFile, makeContent string, makeDiff bool) string {

	if strings.TrimSpace(distPath) == "" {
		return ""
	}

	MakeDir(Abs(distPath))

	var err error
	var compareContent string

	customPath := CustomPath(Abs(distPath, distFile))
	if PathExists(customPath) {
		compareContent, err = ReadFile(customPath)
		if err != nil {
			Error(fmt.Sprintf("Read %s error:", customPath), err)
			return ""
		}
	}

	return p.renderContent(makeContent, compareContent)
}

func (p *Writer) Make() {

	for filePath, renderFile := range p.renderFiles {

		render := true

		if renderFile.Make == "" {
			render = false
		}

		if renderFile.Diff {

			if renderFile.Local == "" {
				if PathExists(filePath) {
					var err error
					renderFile.Local, err = ReadFile(filePath)
					if err != nil {
						Error("Diff read file error: ", err)
						os.Exit(1)
					}
				}
			}

			m1 := md5.New()
			m2 := md5.New()

			m1.Write([]byte(renderFile.Local))
			m2.Write([]byte(renderFile.Make))

			hash1 := hex.EncodeToString(m1.Sum(nil))
			hash2 := hex.EncodeToString(m2.Sum(nil))

			if hash1 == hash2 {
				render = false
			}
		}

		if render && !renderFile.Write {

			if renderFile.Path != "" {
				MakeDir(Abs(renderFile.Path))
			}

			err := WriteFile(filePath, []byte(renderFile.Make))
			if err != nil {
				return
			}

			renderFile.Write = true
			MakeFileSuccess(filePath)
		}
	}
}

func (p *Writer) Clean() {

	for filePath, renderFile := range p.renderFiles {

		if renderFile.Make == "" {
			err := Remove(filePath)
			if err != nil {
				Error(fmt.Sprintf("Remove file error:"), err)
				return
			}
			CleanFileSuccess(filePath)
		}
	}
}

func (p *Writer) renderContent(makeContent, compareContent string) string {

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
