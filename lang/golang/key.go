package golang

import (
	"github.com/koyeo/snippet/helper"
	"strings"
)

var Types = func() []string {
	return []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"string",
		"bool",
		"complex64",
		"complex128",
		"ch",
		"chan",
		"byte",
		"rune",
		"map",
		"struct",
	}
}()

func KeyName(param string) string {
	param = strings.TrimSpace(param)
	if helper.InArray(append([]string{
		"break",
		"case",
		"chan",
		"const",
		"continue",
		"default",
		"defer",
		"else",
		"fallthrough",
		"for",
		"func",
		"go",
		"goto",
		"if",
		"import",
		"interface",
		"map",
		"package",
		"range",
		"return",
		"select",
		"struct",
		"switch",
		"type",
		"var",
	}, Types...), param) {
		return "_" + param
	}

	return param
}
