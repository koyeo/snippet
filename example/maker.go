package main

import "snippet"

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
	project.Render()

}
