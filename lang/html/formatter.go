package html

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Formatter(content string) (output string, err error) {
	outBuffer := new(bytes.Buffer)
	errBuffer := new(bytes.Buffer)
	c := exec.Command("prettier", "--parser", "html")
	c.Stdin = strings.NewReader(content)
	c.Stdout = outBuffer
	c.Stderr = errBuffer
	defer func() {
		if errBuffer.String() != "" {
			err = fmt.Errorf(errBuffer.String())
		}
	}()
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
