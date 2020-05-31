package logger

import (
	"encoding/json"
	"fmt"
	"github.com/ttacon/chalk"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

func Done() {
	log.Println(chalk.Green.Color(chalk.Bold.TextStyle("Command execute done")))
}

func MakeDone() {
	log.Println(chalk.Green.Color(chalk.Bold.TextStyle("Make done")))
}

func Success(msg string) {
	log.Println(chalk.Green.Color(chalk.Bold.TextStyle(msg)))
}

func ReadFileSuccess(path string) {
	wd, _ := os.Getwd()
	path = strings.TrimPrefix(path, wd+"/")
	log.Println(chalk.Green.Color(chalk.Bold.TextStyle("Read file:")), chalk.Green.Color(path))
}

func IgnoreReadPath(msg, path string) {
	log.Println(chalk.Yellow.Color(chalk.Bold.TextStyle(msg)), chalk.Yellow.Color(path))
}

func MakeDirSuccess(path string) {
	wd, _ := os.Getwd()
	path = strings.TrimPrefix(path, wd+"/")
	log.Println(chalk.Green.Color(chalk.Bold.TextStyle("Make dir:  ")), chalk.Green.Color(path))
}

func MakeFileSuccess(path string) {
	wd, _ := os.Getwd()
	path = strings.TrimPrefix(path, wd+"/")
	log.Println(chalk.Green.Color(chalk.Bold.TextStyle("Make file: ")), chalk.Green.Color(path))
}

func CleanFileSuccess(path string) {
	wd, _ := os.Getwd()
	path = strings.TrimPrefix(path, wd+"/")
	log.Println(chalk.Cyan.Color(chalk.Bold.TextStyle("clean file:")), chalk.Cyan.Color(path))
}

func CleanDirSuccess(path string) {
	wd, _ := os.Getwd()
	path = strings.TrimPrefix(path, wd+"/")
	log.Println(chalk.Cyan.Color(chalk.Bold.TextStyle("clean dir:")), chalk.Cyan.Color(path))
}

func Warn(msg string, err ...error) {
	if len(err) > 0 && err[0] != nil {
		fmt.Println(chalk.Yellow, chalk.Bold.TextStyle(msg), err[0].Error())
	} else {
		fmt.Println(chalk.Yellow, chalk.Bold.TextStyle(msg))
	}
}

func Error(msg string, err ...error) {
	if len(err) > 0 && err[0] != nil {
		fmt.Println(chalk.Red, chalk.Bold.TextStyle(msg), err[0].Error())
	} else {
		fmt.Println(chalk.Red, chalk.Bold.TextStyle(msg))
	}
}

func Fatal(msg string, err ...error) {

	if len(err) > 0 && err[0] != nil {
		fmt.Println(chalk.Red, chalk.Bold.TextStyle(msg), err[0].Error())
	} else {
		fmt.Println(chalk.Red, chalk.Bold.TextStyle(msg))
	}

	stack := string(debug.Stack())
	lines := strings.Split(stack, "\n")
	for _, v := range lines {
		fmt.Println(chalk.Yellow, v)
	}

	os.Exit(1)
}

func DebugPrint(elem interface{}) {
	c, err := json.MarshalIndent(elem, "", "\t")
	if err != nil {
		fmt.Println("Call debug print error", err)
	}
	fmt.Println(string(c))
}
