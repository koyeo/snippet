package golang

import (
	"fmt"
	"github.com/koyeo/snippet"
	"github.com/koyeo/snippet/logger"
	"go/format"
	"strings"
)

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
	data.Namespace = snippet.GetNamespace()

	for _, v := range snippet.Packages() {
		data.Packages = append(data.Packages, &renderGolangPackage{
			Name: v.GetName(),
			Path: v.GetPath(),
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

func Formatter(content string) (output string, err error) {
	bytes, err := format.Source([]byte(content))
	if err != nil {
		lines := strings.Split(content, "\n")
		for k, v := range lines {
			fmt.Printf("%d: %s\n", k+1, v)
		}
		logger.Fatal(fmt.Sprintf("Foramt file %s error:", content), err)
		return
	}
	output = string(bytes)

	return
}

func Render(item *snippet.Snippet) (string, error) {
	return snippet.Render(fileTpl, prepareRenderGolangData(item))
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
