package html

import (
	"bytes"
	"os/exec"
	"strings"
)

func Formatter(content string) (output string, err error) {
	outBuffer := new(bytes.Buffer)
	c2 := exec.Command("prettier", "--parser", "html")
	c2.Stdin = strings.NewReader(content)
	c2.Stdout = outBuffer
	err = c2.Start()
	if err != nil {
		return
	}
	err = c2.Wait()
	if err != nil {
		return
	}
	output = outBuffer.String()
	return
}
