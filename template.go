package snippet

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/flosch/pongo2"
	"regexp"
	"strings"
)

func Render(ctx pongo2.Context, content string, data interface{}) (res string, err error) {

	content = fmt.Sprintf("{%s autoescape off %s}%s{%s endautoescape %s}", "%", "%", content, "%", "%")
	tpl, err := pongo2.FromString(strings.TrimSpace(content))
	if err != nil {
		return
	}

	res, err = tpl.Execute(makeContext(ctx, data))
	if err != nil {
		return
	}

	return
}

func TrimSpace(content string) string {

	newLineRegex := regexp.MustCompile(`<\s*\\n\s*>`)
	spaceRegex := regexp.MustCompile(`<\s*\\s\s*>`)
	segments := strings.Split(content, "\n")
	results := make([]string, 0)

	for _, v := range segments {

		v = strings.TrimSpace(v)

		if newLineRegex.MatchString(v) {
			results = append(results, "")
			v = newLineRegex.ReplaceAllString(v, "")
		} else if spaceRegex.MatchString(v) {
			v = spaceRegex.ReplaceAllString(v, " ")
		}

		if v != "" {
			results = append(results, v)
		}
	}

	return strings.Join(results, "\n")
}

func wrapCtxFunc(ctx1 *pongo2.Context, ctx2 pongo2.Context) {
	for k, v := range ctx2 {
		(*ctx1)[k] = v
	}
}

func makeContext(ctx pongo2.Context, input interface{}) pongo2.Context {

	if ctx == nil {

	}

	data := make(pongo2.Context)

	if input != nil {
		if v1, ok := input.(pongo2.Context); ok {
			data = v1
		} else if structs.IsStruct(input) {
			data = structs.Map(input)
		}
	}

	if ctx != nil {
		wrapCtxFunc(&data, ctx)
	}

	return data
}
