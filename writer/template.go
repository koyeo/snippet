package writer

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/flosch/pongo2"
	"regexp"
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

	if data != nil {
		res = filterTags(res)
	}

	return
}

func filterTags(content string) string {

	newLineRegex := regexp.MustCompile(`<\s*\\n\s*>`)
	spaceRegex := regexp.MustCompile(`<\s*\\s\s*>`)
	segments := strings.Split(content, "\n")
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
