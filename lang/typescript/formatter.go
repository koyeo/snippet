package typescript

import (
	"bytes"
	"fmt"
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
		err = fmt.Errorf(err.Error() + " " + errBuffer.String())
		return
	}
	err = c.Wait()
	if err != nil {
		err = fmt.Errorf(err.Error() + " " + errBuffer.String())
		return
	}
	output = outBuffer.String()
	return
}
