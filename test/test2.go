package main

import (
	"fmt"
	"github.com/ttacon/chalk"
	"log"
)

func main() {
	msg := "test"
	err := fmt.Errorf("ok")
	log.Println(chalk.Red.NewStyle().Style(chalk.Bold.TextStyle("[Error] "+msg)), chalk.Red.NewStyle().Style(err.Error()))
	log.Println(chalk.Red.NewStyle().Style(chalk.Bold.TextStyle("[Error] "+msg)), chalk.Red.NewStyle().Style(err.Error()))
	log.Println(chalk.Red.Color("你好吗"))
}
