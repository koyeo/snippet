package main

import (
	"github.com/koyeo/snippet"
	"github.com/koyeo/snippet/lang/golang"
	"github.com/koyeo/snippet/suffix"
)

type RenderData struct {
	Hi string
}

func main() {
	// 定义生成文件
	testFile := snippet.NewFile()
	testFile.SetName("test")
	testFile.SetDir("file-dir")
	testFile.SetSuffix(".md")
	testFile.SetMakeSuffix(".make")
	testFile.SetContent(`Hello world!`, nil)

	// 定义生成目录
	testFolder := snippet.NewFolder()
	testFolder.SetName("test")
	testFolder.SetDir("folder-dir")
	testFolder.SetMakePrefix("make-")
	testFolder.AddFile(testFile)

	// 定义生成代码块文件
	data := new(RenderData)
	data.Hi = "Hello world"

	packageFmt := snippet.NewPackage("", "fmt")
	mainBlock := golang.NewFunc("main", mainTpl, data)
	mainBlock.UsePackage(packageFmt)

	s := snippet.NewSnippet(suffix.Go)
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

	// 定义项目
	project := snippet.NewProject()
	project.SetRoot("./example")
	project.AddFile(testFile)
	project.AddFolder(testFolder)
	project.SetIgnore("vendor")
	project.AddSnippet(s)
	project.Render()
}

const mainTpl = `
<\n>
func main(){
	fmt.Println("{{ Hi }}")
}
`