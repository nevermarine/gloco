package main

import (
	"github.com/nevermarine/gloco/parse"
	"github.com/nevermarine/gloco/cli"
	"github.com/nevermarine/gloco/tmpl"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	cli.ArgsInit()

	project, err := parse.LoadFile(cli.ComposeFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	byteTmplFile, err := fs.ReadFile(tmpl.TmplFile, tmpl.TmplFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	iniFileStr, err := parse.CreateIni(project, byteTmplFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Remove(cli.IniFilePath)
	f, err := os.Create(cli.IniFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	
	_, err = f.WriteString(iniFileStr)
	if err != nil {
		fmt.Println(err)
	}

}
