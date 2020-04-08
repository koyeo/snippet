package sqlmap

import "fmt"

// attributeKey, attributeValue
func AttributeFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`%s\s*=\s*"%s"`, values...))
	return n
}

type Filter struct {
	rule string
}

func (p *Filter) SetRule(rule string) {
	p.rule = rule
}

func (p *Filter) GetRule() string {
	return p.rule
}
