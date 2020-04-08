# 1. Snippet

用来生成代码片段的包。

<!-- TOC depthFrom:2 --><!-- /TOC -->
* [1. 特点](#1-特点)
* [2. 安装](#2-安装)
* [3. 快速上手](#3-快速上手). [4. 项目定义](#4-项目定义)
  * [4.1 新项目](#41-新项目)
  * [4.2 设置项目路径](#42-设置项目路径)
  * [4.3 设置遍历忽略路径](#43-设置遍历忽略路径)
  * [4.4 开启调试模式](#44-开启调试模式)
  * [4.5 添加文件](#45-添加文件)
  * [4.6 添加目录](#46-添加目录)
  * [4.7 添加代码段文件](#47-添加代码段文件)
  * [4.8 执行项目生成](#48-执行项目生成)
* [5. 生成文件](#5-生成文件)
  * [5.1 新文件](#51-新文件)
  * [5.2 设置文件名](#52-设置文件名)
  * [5.3 设置文件路径](#53-设置文件路径)
  * [5.4 设置生成后缀](#54-设置生成后缀)
  * [5.5 设置生成前缀](#55-设置生成前缀)
  * [5.6 设置生成内容](#56-设置生成内容)
* [6. 生成目录](#6-生成目录)
  * [6.1 新目录](#61-新目录)
  * [6.2 设置目录名称](#62-设置目录名称)
  * [6.3 设置目录路径](#63-设置目录路径)
  * [6.4 设置生成后缀](#64-设置生成后缀)
  * [6.5 设置生成前缀](#65-设置生成前缀)
  * [6.6 添加文件](#66-添加文件)
* [7 生成代码块文件](#7-生成代码块文件)
  * [7.1 新文件](#71-新文件)
  * [7.2 设置文件名](#72-设置文件名)
  * [7.3 设置文件路径](#73-设置文件路径)
  * [7.4 设置生成后缀](#74-设置生成后缀)
  * [7.5 设置生成前缀](#75-设置生成前缀)
  * [7.6 设置命名空间](#76-设置命名空间)
  * [7.7 添加标签](#77-添加标签)
  * [7.8 设置渲染函数](#78-设置渲染函数)
  * [7.9 设置格式化函数](#79-设置格式化函数)
  * [7.10 添加常量](#710-添加常量)
  * [7.11 添加代码块](#711-添加代码块)
  * [7.12 文件合并](#712-文件合并)
  * [7.13 提交文件](#713-提交文件)
* [8. 代码块定义](#8-代码块定义)
  * [8.1 新代码块](#81-新代码块)
  * [8.2 定义依赖包](#82-定义依赖包)
  * [8.3 使用依赖包](#83-使用依赖包)
* [9. 模板渲染](#9-模板渲染)
  * [9.1 渲染数据](#91-渲染数据)
  * [9.2 If 语句](#92-if-语句)
  * [9.3 For 循环](#93-for-循环)
  * [9.4 模板函数](#94-模板函数)
  * [9.5 空格和换行](#95-空格和换行)
* [10. 去重机制](#10-去重机制)
* [11. 完整示例](#11-完整示例)
* [12. 问题反馈](#12-问题反馈)


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

## 4. 项目定义


### 4.1 新项目

```go
project := snippet.NewProject()
```

### 4.2 设置项目路径

```go
project.SetRoot("./make-example")
```

用来存放生成的所有文件和目录，同时也是自动去除冗余文件的遍历根目录。

### 4.3 设置遍历忽略路径

```go
project.SetIgnore(
  "vendor/*",
  "node_modules"
)
```
可以设置第三方库的目录，提高遍历性能。

### 4.4 开启调试模式

```go
project.SetDebug()
```
打开调试模式，可以打印运行过程中的一些关键日志，如观察设置的遍历忽略路径是否生效。

### 4.5 添加文件

```go
project.AddFile(file1, file2, ...)
```
添加需要生成的[文件](#5-生成文件)。

### 4.6 添加目录

```go
project.AddFolder(folder1, folder2, ...)
```
添加需要生成的[目录](#6-生成目录)。

### 4.7 添加代码段文件

```go
project.AddSnippet(snippet1, snippet2, ...)
```
添加需要生成的[代码段文件](#7-生成代码块文件)。

### 4.8 执行项目生成

```
project.Render()
```
执行项目生成操作。


## 5. 生成文件

生成文件是以文件为去重单位，当同位置有不带后缀和前缀的同名文件存在时，便不会执行文件的生成操作；同时原来生成的文件将会被清除。

### 5.1 新文件

```go
file := snippet.NewFile()
```

### 5.2 设置文件名

```go
file.SetName("path/to/test")
```

### 5.3 设置文件路径

```go
file.SetDir("static/files")
```

### 5.4 设置生成后缀
```go
file.SetMakeSuffix(".make")
```

可以通过[生成后缀](#10-去重机制)来区分生成文件和自定义文件。

### 5.5 设置生成前缀

```go
file.SetMakePrefix("make-")
```

可以通过[生成前缀](#10-去重机制)来区分生成文件和自定义文件。

### 5.6 设置生成内容

```
file.SetContent(template, data)
```

生成内容支持[模板渲染](#9-模板渲染)。


## 6. 生成目录

生成目录是以目录为去重单位，当同位置有不带后缀和前缀的同名目录存在时，便不会执行目录及其内部文件的生成操作；同时原来生成的目录将会被清除。

### 6.1 新目录

```go
folder := snippet.NewFolder()
```

### 6.2 设置目录名称

```go
folder.SetName("test")
```

### 6.3 设置目录路径

```go
folder.SetDir("test")
```

### 6.4 设置生成后缀
```go
folder.SetMakeSuffix(".make")
```

可以通过[生成后缀](#10-去重机制)来区分生成目录和自定义目录。

### 6.5 设置生成前缀

```go
folder.SetMakePrefix("make-")
```

可以通过[生成前缀](#10-去重机制)来区分生成目录和自定义目录。

### 6.6 添加文件

```go
folder.AddFile(file1, file2, ...)
```

## 7 生成代码块文件

生成代码块是以代码块为去重单位，当同位置有不带后缀和前缀的同名文件里存在同名的代码块时，便不会执行代码块的生成操作。

### 7.1 新文件

```go 
s := snippet.NewSnippet(snippet.SuffixGo)
```

### 7.2 设置文件名

```go
s.SetName("main")
```

### 7.3 设置文件路径

```go
s.SetDir("main")
```

### 7.4 设置生成后缀
```go
s.SetMakeSuffix(".make")
```

可以通过[生成后缀](#10-去重机制)来区分生成文件和自定义文件。

### 7.5 设置生成前缀

```go
s.SetMakePrefix("make-")
```
可以通过[生成前缀](#10-去重机制)来区分生成文件和自定义文件。

### 7.6 设置命名空间
```go
s.SetNamespace("main")
```

### 7.7 添加标签

```go
s.AddTag("build dev")
```

### 7.8 设置渲染函数

```go
s.SetRender(goalng.Render)
```

使用[预定义渲染函数](https://github.com/koyeo/snippet/blob/master/golang/render.go)。

### 7.9 设置格式化函数

```go
s.SetFormatter(golang.Formatter)
```
使用[预定义格式化函数](https://github.com/koyeo/snippet/blob/master/golang/render.go)。

### 7.10 添加常量

```go
const1 := golang.NewConstant("Name","foo")

s.AddConstant(const1, const2, ...)
```

使用 [golang.NewConstant](https://github.com/koyeo/snippet/blob/master/golang/types.go) 快速创建常量代码块。

### 7.11 添加代码块

```go
block1 := golang.NewFunc("main", mainTpl, renderData)
block2 := golang.NewStruct("User", userTpl, renderData)

s.AddBlock(block1, block2)
```


查看[代码块定义](#8-代码块定义)详情。

### 7.12 文件合并

```go
s.Merge(s2, s3, ...)
```

文件合并以文件路径为索引。通过合并可能可以在不同的地方定义代码块，最终合并生成一个文件。

### 7.13 提交文件

> 注：在代码块文件定义完后必须执行此操作

```go
s.Commit()
```

## 8. 代码块定义

### 8.1 新代码块

```go
block := snippet.NewBlock(filter, template, data)
```

`filter` 是用来进行代码块去重匹配的正则表达式，通过 `github.com/koyeo/snippet/golang` 包可以快速构建 Golang 相关的代码块：

```go
golang.NewConstant(name, value string) 
golang.NewVar(name, value string)  
golang.NewType(name, code string) 
golang.NewFunc(name, code string, data interface{}) 
golang.NewStruct(name, code string, data interface{}) 
golang.NewStructFunc(structName, funcName, code string, data interface{})
```



### 8.2 定义依赖包

```
packageFmt := snippet.NewPackage("fmt", "fmt")
packageMySql := snippet.NewPackage("_", "github.com/go-sql-driver/mysql")
```

### 8.3 使用依赖包

```
block.UsePackage(packageFmt, packageMySql, ...)
```

依赖包与代码块绑定，当代码块生成时（没有与自定义代码块重复），才会添加至依赖包列表。


## 9. 模板渲染

`File.SetContent()` 方法和  `Block.SetCode()` 方法均支持模板渲染，渲染引擎采用了 [pongo2](https://github.com/flosch/pongo2)。

### 9.1 渲染数据

要求传入数据类型是 `struct`。

### 9.2 If 语句

```
{% if Tag == 1 %}
...
{% elif Tag == 2 %}
...
{% endif %}
```

### 9.3 For 循环

```
{% for route in Routes %}
   {{forloop.Counter}} {{ route.Name }}
{% endfor %}
```

### 9.4 模板函数

Pongo2 支持模板函数，在此项目中暂时没有支持，希望所有的渲染数据提前预处理。

### 9.5 空格和换行

pongo2 渲染后的数据对于首空格和换行的控制有点混乱，此项目在使用支持模板渲染的方法时，传入的数据不为 `nil`，需要使用 `<\n>` 和 `<\s>` 标签来控制生成内容的换行和空格数量。如果是代码的话，可以使用代码格式化函数来完成，如 [golang.Formatter](https://github.com/koyeo/snippet/blob/master/golang/render.go)。

## 10. 去重机制

去重的核心规则是：生成文件或目录的同位置是否有**不包含生成前缀后生成后缀的同名文件或目录**。如生成文件时（File）如果生成后缀是 `.make`，文件 'example/test.md' 已存在，则 `example.make.md` 便不会再生成。

生成代码块文件(Snippet) 时会进入同名文件内部，检查是否有同名代码块。如 `example/main.go` 中已存在 `main` 函数，则生成的 `example.make.go` 文件中，则不会出现 `main` 函数及其依赖的包。如果一个生成代码块文件，所有的代码块均呗自定义，则已生成的文件会被清除。

## 11. 完整示例

```go
package main

import (
	"github.com/koyeo/snippet"
	"github.com/koyeo/snippet/golang"
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
```

## 12. 问题反馈

如果你对本项目感兴趣或有什么问题，欢迎联系：

![](qrcode.png)
