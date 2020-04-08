package main

import (
	"snippet"
	"snippet/golang"
)

func main() {
	testFile := snippet.NewFile()
	testFile.SetName("test")
	testFile.SetDir("file-dir")
	testFile.SetSuffix(".md")
	testFile.SetMakeSuffix(".make")
	testFile.SetContent(`Hello world!`, nil)

	testFolder := snippet.NewFolder()
	testFolder.SetName("test")
	testFolder.SetDir("folder-dir")
	testFolder.SetMakePrefix("make-")
	testFolder.AddFile(testFile)

	project := snippet.NewProject()
	project.SetRoot("./make-example")
	project.AddFile(testFile)
	project.AddFolder(testFolder)
	project.Render()
}

type RenderData struct {
	Hi string
}

func addSnippets(project *snippet.Project) {
	data := new(RenderData)
	data.Hi = "Hello world"

	packageFmt := snippet.NewPackage("", "fmt")

	mainBlock := golang.NewFunc("main", mainTpl, data)
	mainBlock.UsePackage(packageFmt)

	s := snippet.NewSnippet(snippet.SuffixGo)
	s.SetName("test")
	s.SetNamespace("main")
	s.SetDir("snippets")
	s.AddTag("build dev")
	s.SetMakeSuffix(".mix")
	s.AddBlock(mainBlock)
	s.AddConstant(golang.NewConstant("ok", `"ok123"`))
	s.SetRender(golang.Render)
	s.SetFormatter(golang.Formatter)
	s.Commit()

	project.AddSnippet(s)
}

const mainTpl = `
<\n>
func main(){




	fmt.Println("{{ Hi }}")
}
`
