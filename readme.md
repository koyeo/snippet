# Snippet

用来生成代码片段的包。

* [1. 特点](#1-特点)
* [2. 安装](#2-安装)

## 1. 特点

* 生成代码片段均支持自定义覆盖；
* 支持按目录和文件定义代码片段；
* 支持代码中自定义包依赖引入；
* 支持自定义多语言代码模板渲染；
* 支持旧文件或目录自动清除；

## 2. 安装

```
$ go get -v -u github.com/koyeo/snippet
```

## 3. 快速上手

```go
package main

import (
	"github.com/koyeo/snippet"
	"github.com/koyeo/snippet/golang"
)

func main() {
	testFile := snippet.NewFile()
	testFile.SetName("test")
	testFile.SetSuffix(".md")
	testFile.SetMakeSuffix(".make")
	testFile.SetContent(`Hello world!`, nil)

	testFolder := snippet.NewFolder()
	testFolder.SetName("test")
	testFolder.SetMakePrefix("make-")
	testFolder.AddFile(testFile)

	project := snippet.NewProject()
	project.SetRoot("./make-example")
	project.AddFile(testFile)
	project.AddFolder(testFolder)
	project.Render()
}
```

运行结果：

```sh
$ tree ./make-example 
./make-example
├── make-test
│   └── test.make.md
└── test.make.md
```

```sh
$ cat ./make-example/test.make.md 
Hello world!
```

## 项目定义


```
```
