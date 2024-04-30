package main

import (
	"github.com/nevermarine/gloco/parse"
	"github.com/nevermarine/gloco/cli"
	"fmt"
)

func main() {
	cli.ArgsInit()

	project, err := parse.LoadFile(cli.ComposeFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parse.WriteIni(project, cli.IniFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
}
