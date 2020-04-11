package sqlmap

type renderData struct {
	Blocks    []*renderBlock
}


type renderBlock struct {
	Rule string
	Code string
}

//func prepareRenderData(snippet core.Snippet) *renderData {
//
//	data := new(renderData)
//
//	for _, v := range snippet.GetBlocks() {
//		if v.GetFilter() != nil {
//			data.Blocks = append(data.Blocks, &renderBlock{
//				Rule: v.GetFilter().Rule(),
//				Code: v.GetCode(),
//			})
//		} else {
//			data.Blocks = append(data.Blocks, &renderBlock{
//				Code: v.GetCode(),
//			})
//		}
//	}
//
//	return data
//}

//func Render(snippet core.Snippet) (content string, err error) {
//
//	//content, err = template.Render(fileTpl, prepareRenderData(snippet))
//	//if err != nil {
//	//	return
//	//}
//
//	return
//}

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
