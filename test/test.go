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
	s.SetDir("snippets2")
	s.AddTag("build dev")
	s.SetMakeSuffix(".mix")
	s.AddBlock(mainBlock)
	s.AddConstant(golang.NewConstant("ok", `"ok123"`))
	s.SetRender(golang.Render)
	s.SetFormatter(golang.Formatter)
	s.Commit()

	// 定义工作空间
	workspace1 := snippet.NewWorkspace()
	workspace1.SetRoot("./example", true)
	workspace1.AddFile(testFile)
	workspace1.AddFolder(testFolder)
	workspace1.AddIgnorePath("vendor")
	workspace1.AddSnippet(s)

	workspace2 := snippet.NewWorkspace()
	workspace2.SetRoot("./example2", true)
	workspace2.AddFile(testFile)
	workspace2.AddFolder(testFolder)
	workspace2.AddIgnorePath("vendor")
	workspace2.AddSnippet(s)

	// 定义项目
	project := snippet.NewProject()
	project.AddWorkspace(workspace1, workspace2)
	project.Render()
}

const mainTpl = `
<\n>
func main(){
	fmt.Println("{{ Hi }}")
}
`
