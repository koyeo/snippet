package golang

import (
	"fmt"
	"github.com/koyeo/snippet"
)

func NewConstant(name, value string) *snippet.Block {
	return snippet.NewBlock(ConstFilter(name), fmt.Sprintf("%s=%s", name, value), nil)
}

func NewVar(name, value string) *snippet.Block {
	return snippet.NewBlock(VarFilter(name), fmt.Sprintf("%s=%s", name, value), nil)
}

func NewType(name, code string) *snippet.Block {
	return snippet.NewBlock(TypeFilter(name), code, nil)
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
