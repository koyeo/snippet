package main

import (
	"snippet"
	"snippet/golang"
)

func main() {

	testFile := snippet.NewFile()
	testFile.SetName("test2")
	testFile.SetSuffix(".md")
	testFile.SetMakeSuffix(".mix")
	testFile.SetContent(`你好世界`, nil)

	testFolder := snippet.NewFolder()
	testFolder.SetName("components")
	testFolder.SetMakePrefix("mix-")
	testFolder.AddFile(testFile)

	project := snippet.NewProject()
	project.SetRoot("./snippet-test")
	project.AddFile(testFile)
	project.AddFolder(testFolder)
	addSnippets(project)
	project.Render()

}

type RenderData struct {
	Hi string
}

func addSnippets(project *snippet.Project) {
	data := new(RenderData)
	data.Hi = "Hello world"

	packageFmt := snippet.NewPackage("", "fmt")

	mainBlock := snippet.NewBlock(golang.FuncFilter("main"), mainTpl, data)
	mainBlock.UsePackage(packageFmt)

	s := snippet.NewSnippet(snippet.SuffixGo)
	s.SetName("test")
	s.SetNamespace("main")
	s.SetDir("snippets")
	s.AddTag("build dev")
	//s.SetMakeSuffix(".mix")
	s.AddBlock(mainBlock)
	s.AddConstant(golang.NewConstant("ok", `"ok123"`))
	s.SetRender(golang.Render, golang.Formatter)
	s.Commit()

	project.AddSnippet(s)
}

const mainTpl = `
<\n>
func main(){
	fmt.Println("{{ Hi }}")
}
`
