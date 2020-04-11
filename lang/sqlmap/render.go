package sqlmap

import (
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet"
)

type renderData struct {
	Blocks []*renderBlock
}

type renderBlock struct {
	Rule string
	Code string
}

func prepareRenderData(snippet *snippet.Snippet) *renderData {

	data := new(renderData)

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
<sqlMap>
<\n>
{% for block in Blocks %}
	<block rule="{{ block.Rule }}">
		{{ block.Code }}
		<\n>
	</block>
{% endfor %}
</sqlMap>
`
