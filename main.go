package main

import (
	"github.com/nevermarine/gloco/parse"
	"fmt"
)

func main() {
	project, err := parse.LoadFile("docker-compose.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = parse.WriteIni(project, "postgres.service")
	if err != nil {
		fmt.Println(err)
		return
	}
}
