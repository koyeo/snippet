# 1. Snippet

用来生成代码片段的包。

<!-- TOC depthFrom:2 -->autoauto1. [1.1. 特点](#11-特点)auto2. [1.2. 安装](#12-安装)auto3. [1.3. 快速上手](#13-快速上手)auto4. [1.4. 项目定义](#14-项目定义)auto    1. [1.4.1. 新项目](#141-新项目)auto    2. [1.4.2. 项目路径](#142-项目路径)auto    3. [1.4.3. 遍历忽略路径](#143-遍历忽略路径)auto    4. [1.4.4. 调试模式](#144-调试模式)auto    5. [1.4.5. 添加文件](#145-添加文件)auto    6. [1.4.6. 添加目录](#146-添加目录)auto    7. [1.4.7. 添加代码段文件](#147-添加代码段文件)auto    8. [1.4.8. 运行](#148-运行)autoauto<!-- /TOC -->

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

### 4.2 项目路径

```go
project.SetRoot("./make-example")
```

用来存放生成的所有文件和目录，同时也是自动去除冗余文件的遍历根目录。

### 4.3 遍历忽略路径

```go
project.SetIgnore(
  "vendor/*",
  "node_modules"
)
```
可以设置第三方库的目录，提高遍历性能。

### 4.4 调试模式

```go
project.SetDebug()
```
打开调试模式，可以打印运行过程中的一些关键日志，如观察设置的遍历忽略路径是否生效。

### 4.5 添加文件

```go
project.AddFile(file1, file2, ...)
```
添加需要生成的[文件]()。

### 4.6 添加目录

```go
project.AddFolder(folder1, folder2, ...)
```
添加需要生成的[目录]()。

### 4.7 添加代码段文件

```go
project.AddSnippet(snippet1, snippet2, ...)
```
添加需要生成的[代码段文件]()。

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

可以通过[生成后缀]()来区分生成文件和自定义文件。

### 5.5 设置生成前缀

```go
file.SetMakePrefix("make-")
```

可以通过[生成前缀]()来区分生成文件和自定义文件。

### 5.6 设置生成内容

```
file.SetContent(template, data)
```

生成内容支持[模板渲染]()。


## 6. 生成目录

生成目录是以目录为去重单位，当同位置有不带后缀和前缀的同名目录存在时，便不会执行目录及其内部文件的生成操作；同时原来生成的目录将会被清除。

### 新目录

```go
folder := snippet.NewFolder()
```

### 设置目录名称

```go
folder.SetName("test")
```

### 设置目录路径

```go
folder.SetDir("test")
```

### 5.4 设置生成后缀
```go
folder.SetMakeSuffix(".make")
```

可以通过[生成后缀]()来区分生成目录和自定义目录。

### 5.5 设置生成前缀

```go
folder.SetMakePrefix("make-")
```

可以通过[生成前缀]()来区分生成目录和自定义目录。

### 添加文件

```go
folder.AddFile(file1, file2, ...)
```

## 生成代码块

生成代码块是以代码块为去重单位，当同位置有不带后缀和前缀的同名文件里存在同名的代码块时，便不会执行代码块的生成操作。

### 新文件

```go 
s := snippet.NewSnippet(snippet.SuffixGo)
```

### 设置文件名

```go
s.SetName("main")
```

### 设置文件路径

```go
s.SetDir("main")
```

### 设置生成后缀
```go
s.SetMakeSuffix(".make")
```

可以通过[生成后缀]()来区分生成文件和自定义文件。

### 设置生成前缀

```go
s.SetMakePrefix("make-")
```

### 设置命名空间
```go
s.SetNamespace("main")
```

### 添加标签

```go
s.AddTag("build dev")
```

### 设置渲染函数

```go
s.SetRender(goalng.Render)
```

使用[预定义渲染函数]()。

### 设置格式化函数

```go
s.SetFormatter(golang.Formatter)
```
使用[预定义格式化函数]()。

### 添加常量

```go
const1 := golang.NewConstant("Name","foo")

s.AddConstant(const1, const2, ...)
```

使用 [golang.NewConstant]() 快速创建常量代码块。

### 添加代码块

```go
block1 := golang.NewFunc("main", mainTpl, renderData)
block2 := golang.NewStruct("User", userTpl, renderData)

s.AddBlock(block1, block2)
```


查看[代码块定义]()详情。

### 文件合并

```go
s.Merge(s2, s3, ...)
```

文件合并以文件路径为索引。通过合并可能可以在不同的地方定义代码块，最终合并生成一个文件。

### 提交文件

> 注：在代码块文件定义完后必须执行此操作

```go
s.Commit()
```

## 代码块定义

### 新代码块

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



### 定义依赖包

```
packageFmt := snippet.NewPackage("fmt", "fmt")
packageMySql := snippet.NewPackage("_", "github.com/go-sql-driver/mysql")
```

### 使用依赖包

```
block.UsePackage(packageFmt, packageMySql, ...)
```

依赖包与代码块绑定，当代码块生成时（没有与自定义代码块重复），才会添加至依赖包列表。


## 模板渲染

## 去重机制

## 