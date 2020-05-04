package typescript

import (
	"bytes"
	"os/exec"
	"strings"
)

func Formatter(content string) (output string, err error) {
	outBuffer := new(bytes.Buffer)
	errBuffer := new(bytes.Buffer)
	c := exec.Command("prettier", "--parser", "typescript", "--single-quote")
	c.Stdin = strings.NewReader(content)
	c.Stdout = outBuffer
	c.Stderr = errBuffer
	err = c.Start()
	if err != nil {
		return
	}
	err = c.Wait()
	if err != nil {
		return
	}
	output = outBuffer.String()
	return
}
