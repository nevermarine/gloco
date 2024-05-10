package main

import (
	"github.com/nevermarine/gloco/parse"
	"github.com/nevermarine/gloco/cli"
	"github.com/nevermarine/gloco/tmpl"
	"fmt"
	"io/fs"
)

func main() {
	cli.ArgsInit()

	project, err := parse.LoadFile(cli.ComposeFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	byteTmplFile, err := fs.ReadFile(tmpl.TmplFile, "systemd.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parse.WriteIni(project, cli.IniFilePath, byteTmplFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}
