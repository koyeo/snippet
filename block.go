package snippet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/koyeo/snippet/logger"
	"io"
	"log"
	"strings"
)

func NewBlock(filter Filter, code string, data interface{}) *Block {
	b := &Block{
		filter: filter,
	}
	b.SetCode(code, data)
	return b
}

func NewStaticBlock(filter Filter, code string) *Block {
	return NewBlock(filter, code, nil)
}

func NewDocument(document string) *Block {
	b := &Block{
		filter: nil,
	}
	b.SetCode(document, nil)
	return b
}

type Block struct {
	filter   Filter
	code     string
	packages *Packages
	exist    bool
}

func (p *Block) md5() string {
	h := md5.New()
	_, _ = io.WriteString(h, p.code)
	return hex.EncodeToString(h.Sum(nil))
}

func (p *Block) SetExist(exist bool) {
	p.exist = exist
}

func (p *Block) Exist() bool {
	return p.exist
}

func (p *Block) initPackages() {
	if p.packages == nil {
		p.packages = NewPackages()
	}
}

func (p *Block) UsePackage(packages ...*Package) {
	p.initPackages()
	p.packages.Add(packages...)
}

func (p *Block) Packages() []*Package {
	p.initPackages()
	return p.packages.All()
}

func (p *Block) SetFilter(filter Filter) {
	p.filter = filter
}

func (p *Block) Filter() (filter Filter) {
	return p.filter
}

func (p *Block) SetCode(code string, data interface{}) {

	_code, err := Render(nil, code, data)
	if err != nil {
		log.Println("render code:")
		items := strings.Split(code, "\n")
		for i, v := range items {
			fmt.Printf("%d: %s\n", i+1, v)
		}
		logger.Fatal("render content error: ", err)
	}
	p.code = _code
	return
}

func (p *Block) Code() string {
	return p.code
}
