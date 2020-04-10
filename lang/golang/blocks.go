package golang

import (
	"fmt"
	"github.com/koyeo/snippet"
)

func NewConstant(name, value string) *snippet.Block {
	return snippet.NewBlock(ConstFilter(name), fmt.Sprintf("%s=%s", name, value), nil)
}

func NewVar(name, code string, data interface{}) *snippet.Block {
	return snippet.NewBlock(VarFilter(name), code, data)
}

func NewType(name, code string, data interface{}) *snippet.Block {
	return snippet.NewBlock(TypeFilter(name), code, data)
}

func NewFunc(name, code string, data interface{}) *snippet.Block {
	return snippet.NewBlock(FuncFilter(name), code, data)
}

func NewStruct(name, code string, data interface{}) *snippet.Block {
	return snippet.NewBlock(StructFilter(name), code, data)
}

func NewStructFunc(structName, funcName, code string, data interface{}) *snippet.Block {
	return snippet.NewBlock(StructFuncFilter(structName, funcName), code, data)
}
