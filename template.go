package snippet

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/flosch/pongo2"
	"strings"
)

func Render(content string, data interface{}) (res string, err error) {

	content = fmt.Sprintf("{%s autoescape off %s}%s{%s endautoescape %s}", "%", "%", content, "%", "%")
	tpl, err := pongo2.FromString(strings.TrimSpace(content))
	if err != nil {
		return
	}

	res, err = tpl.Execute(makeContext(data))
	if err != nil {
		return
	}

	return
}

func makeContext(input interface{}) pongo2.Context {

	ctx := make(pongo2.Context)

	if input != nil {
		if v1, ok := input.(pongo2.Context); ok {
			ctx = v1
		} else if structs.IsStruct(input) {
			ctx = structs.Map(input)
		}
	}

	return ctx
}
