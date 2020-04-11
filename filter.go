package snippet

type Filter interface {
	SetRule(rule string)
	Rule() (rule string)
}
