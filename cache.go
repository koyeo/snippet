package snippet

import (
	"regexp"
)

func NewCache() *Cache {
	return &Cache{}
}

type Cache struct {
	fileMap map[string]string
}

func (p *Cache) Add(path string) {

	if p.fileMap == nil {
		p.fileMap = make(map[string]string)
	}

	if _, ok := p.fileMap[path]; ok {
		return
	}

	p.fileMap[path], _ = ReadFile(path)

	return
}

func (p *Cache) Read(path string) string {
	if _, ok := p.fileMap[path]; !ok {
		return ""
	}
	return p.fileMap[path]
}

func (p *Cache) Match(path, rule string) bool {

	content := p.Read(path)
	if content == "" {
		return false
	}

	reg := regexp.MustCompile(rule)
	if !reg.MatchString(content) {
		return false
	}

	return true
}
