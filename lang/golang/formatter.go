package golang

import (
	"fmt"
	"github.com/koyeo/snippet/logger"
	"go/format"
	"strings"
)

func Formatter(content string) (output string, err error) {
	bytes, err := format.Source([]byte(content))
	if err != nil {
		lines := strings.Split(content, "\n")
		for k, v := range lines {
			fmt.Printf("%d: %s\n", k+1, v)
		}
		logger.Fatal(fmt.Sprintf("Foramt error:\n%s", content), err)
		return
	}
	output = string(bytes)

	return
}
