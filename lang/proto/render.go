package proto

import (
	"github.com/koyeo/snippet"
	"mix/template"
)

type renderData struct {
	Namespace string
	Packages  []*renderPackage
	Blocks    []*renderBlock
}

type renderPackage struct {
	Name string
	Path string
}

type renderBlock struct {
	Rule string
	Code string
}

func prepareRenderData(snippet *snippet.Snippet) *renderData {

	data := new(renderData)

	data.Namespace = snippet.GetNamespace()

	for _, v := range snippet.Packages() {
		data.Packages = append(data.Packages, &renderPackage{
			Name: v.GetName(),
			Path: v.GetPath(),
		})
	}
	for _, v := range snippet.Blocks() {
		if v.GetFilter() != nil {
			data.Blocks = append(data.Blocks, &renderBlock{
				Rule: v.GetFilter().GetRule(),
				Code: v.GetCode(),
			})
		} else {
			data.Blocks = append(data.Blocks, &renderBlock{
				Code: v.GetCode(),
			})
		}
	}

	return data
}

func Render(snippet core.Snippet) (content string, err error) {

	content, err = template.Render(fileTpl, prepareRenderData(snippet))
	if err != nil {
		return
	}

	return
}

var fileTpl = `
{% for block in Blocks %}
	{{ block.Code }}
{% endfor %}
`
