package snippet

import (
	"path/filepath"
)

type File struct {
	path string
	name string
}

func (p *File) fullPath() string {
	return filepath.Join(p.path, p.name)
}
