package json

import (
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet"
)

type renderData struct {
	Blocks []*renderBlock
}

type renderBlock struct {
	Code string
}

func prepareRenderData(snippet *snippet.Snippet) *renderData {

	data := new(renderData)

	for _, v := range snippet.Blocks() {

		data.Blocks = append(data.Blocks, &renderBlock{
			Code: v.Code(),
		})

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
