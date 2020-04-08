package golang

import "fmt"

type Filter struct {
	rule string
}

func (p *Filter) SetRule(rule string) {
	p.rule = rule
}

func (p *Filter) GetRule() string {
	return p.rule
}

func TypeFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`type\s+%s\s+.+`, values...))
	return n
}

func ConstFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`%s\s*=`, values...))
	return n
}

func VarFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`var\s+%s\s+.+`, values...))
	return n
}

func FuncFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`func\s+%s\s*\(`, values...))
	return n
}

func StructFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`type\s+%s\s+struct`, values...))
	return n
}

// structName,funcName
func StructFuncFilter(values ...interface{}) *Filter {
	n := &Filter{}
	n.SetRule(fmt.Sprintf(`func\s+\(\s*.+\s+\**%s\)\s+%s\s*\(`, values...))
	return n
}
