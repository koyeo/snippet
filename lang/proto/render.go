package proto

import (
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet"
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

	data.Namespace = snippet.Namespace()

	for _, v := range snippet.Packages() {
		data.Packages = append(data.Packages, &renderPackage{
			Name: v.Name(),
			Path: v.Path(),
		})
	}
	for _, v := range snippet.Blocks() {
		if v.Filter() != nil {
			data.Blocks = append(data.Blocks, &renderBlock{
				Rule: v.Filter().Rule(),
				Code: v.Code(),
			})
		} else {
			data.Blocks = append(data.Blocks, &renderBlock{
				Code: v.Code(),
			})
		}
	}

	return data
}

func Render(ctx pongo2.Context, item *snippet.Snippet) (string, error) {
	return snippet.Render(ctx, fileTpl, prepareRenderData(item))
}

var fileTpl = `
{% for block in Blocks %}
	{{ block.Code }}
{% endfor %}
`
