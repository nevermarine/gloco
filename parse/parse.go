package parse

import (
	"context"
	// "io/fs"
	"os"
	"text/template"
	// "fmt"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
)

func LoadFile(composeFilePath string) (*types.Project, error) {
	options, err := cli.NewProjectOptions([]string{composeFilePath})
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	project, err := cli.ProjectFromOptions(ctx, options)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func WriteIni(project *types.Project, iniFilePath string) error {
	tmplFilePath := "tmpl/systemd.ini"	
	tmplFile, err := os.ReadFile(tmplFilePath)
	if err != nil {
		return err
	}
	tmpl, err := template.New("ini").Parse(string(tmplFile))
	if err != nil {
		return err
	}
	os.Remove(iniFilePath)
	f, err := os.Create(iniFilePath)
	if err != nil {
		return err
	}
	err = tmpl.Execute(f, project.Services["db"])
	f.Close()

	return err
}

// func CreateTemplateMap(project *types.Project) (map[string]string, error) {
// 	tmplMap := make(map[string]string)
// }