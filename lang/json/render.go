package json

//
//type renderData struct {
//	Blocks []*renderBlock
//}
//
//type renderBlock struct {
//	Code string
//}
//
//func prepareRenderData(snippet core.Snippet) *renderData {
//
//	data := new(renderData)
//
//	for _, v := range snippet.GetBlocks() {
//
//		data.Blocks = append(data.Blocks, &renderBlock{
//			Code: v.GetCode(),
//		})
//
//	}
//
//	return data
//}
//
//func Render(snippet core.Snippet) (content string, err error) {
//
//	content, err = template.Render(fileTpl, prepareRenderData(snippet))
//	if err != nil {
//		return
//	}
//
//	return
//}
//
//var fileTpl = `
//{% for block in Blocks %}
//{{ block.Code }}
//{% endfor %}
//`
