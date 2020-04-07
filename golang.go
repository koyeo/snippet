package snippet

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
	Rule string
	Code string
}

func prepareRenderGolangData(snippet *Snippet) *renderGolangData {

	data := new(renderGolangData)

	data.Tags = snippet.GetTags()
	data.Namespace = snippet.GetNamespace()

	for _, v := range snippet.GetPackages() {
		data.Packages = append(data.Packages, &renderGolangPackage{
			Name: v.GetName(),
			Path: v.GetPath(),
		})
	}

	for _, v := range snippet.GetConstants() {
		if v.GetFilter() != nil {
			data.Constants = append(data.Constants, &renderGolangBlock{
				Rule: v.GetFilter().GetRule(),
				Code: v.GetCode(),
			})
		} else {
			data.Constants = append(data.Constants, &renderGolangBlock{
				Code: v.GetCode(),
			})
		}
	}

	for _, v := range snippet.GetBlocks() {
		if v.GetFilter() != nil {
			data.Blocks = append(data.Blocks, &renderGolangBlock{
				Rule: v.GetFilter().GetRule(),
				Code: v.GetCode(),
			})
		} else {
			data.Blocks = append(data.Blocks, &renderGolangBlock{
				Code: v.GetCode(),
			})
		}
	}

	return data
}

func RenderGolang(snippet *Snippet) (content string, err error) {

	content, err = Render(fileTpl, prepareRenderGolangData(snippet))
	if err != nil {
		return
	}

	return
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
			{{ pkg.Name }} "{{ pkg.MakePath }}"
		{% else %}
			"{{ pkg.MakePath }}"
        {% endif %}
	{% endfor %}
)
{% endif %}

{% if Constants %}
<\n>
const (
	{% for block in Constants %}
	<block rule="{{ block.Rule }}">
		{{ block.Code }}
	</block>
	{% endfor %}
)
{% endif %}

<\n>
{% for block in Blocks %}
<block rule="{{ block.Rule }}">
{{ block.Code }}
</block>
{% endfor %}
`
