package snippet

type Filter interface {
	SetRule(rule string)
	GetRule() (rule string)
}
