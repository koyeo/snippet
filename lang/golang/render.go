package golang

import (
	"github.com/flosch/pongo2"
	"github.com/koyeo/snippet"
)

func Render(ctx pongo2.Context, item *snippet.Snippet) (string, error) {
	return snippet.Render(ctx, fileTpl, prepareRenderGolangData(item))
}

type renderGolangData struct {
	Tags      []string
	Namespace string
	Packages  []*renderGolangPackage
	Constants []*renderGolangBlock
	Blocks    []*renderGolangBlock
}

type renderGolangPackage struct {
	Name string
	Path string
}

type renderGolangBlock struct {
	Code string
}

func prepareRenderGolangData(snippet *snippet.Snippet) *renderGolangData {

	data := new(renderGolangData)

	data.Tags = snippet.GetTags()
	data.Namespace = snippet.Namespace()

	for _, v := range snippet.Packages() {
		data.Packages = append(data.Packages, &renderGolangPackage{
			Name: v.Name(),
			Path: v.Path(),
		})
	}

	for _, v := range snippet.Constants() {
		if !v.Exist() {
			data.Constants = append(data.Constants, &renderGolangBlock{
				Code: v.Code(),
			})
		}
	}

	for _, v := range snippet.Blocks() {
		if !v.Exist() {
			data.Blocks = append(data.Blocks, &renderGolangBlock{
				Code: v.Code(),
			})
		}
	}

	return data
}

var fileTpl = `
{% for tag in Tags %}
// {{ tag }}
<\n>
{% endfor %}
package {{ Namespace }}
{% if Packages %}
<\n>
import (
	{% for pkg in Packages %}
		{% if pkg.Name != "" %}
			{{ pkg.Name }} "{{ pkg.Path }}"
		{% else %}
			"{{ pkg.Path }}"
        {% endif %}
	{% endfor %}
)
{% endif %}

{% if Constants %}
<\n>
const (
	{% for block in Constants %}
		{{ block.Code }}
	{% endfor %}
)
{% endif %}

<\n>
{% for block in Blocks %}
{{ block.Code }}
{% endfor %}
`
