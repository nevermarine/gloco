package parse

import (
	"context"
	// "io/fs"
	"bytes"
	"text/template"
	// "fmt"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/nevermarine/gloco/errdefs"
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

func CreateIni(project *types.Project, tmplFile []byte) (string, error) {
	tmpl, err := template.New("ini").Parse(string(tmplFile))
	if err != nil {
		return "", err
	}
	switch {
	case len(project.Services) > 1:
		return "", errdefs.ErrMultipleServices
	case len(project.Services) < 1:
		return "", errdefs.ErrNoService
	}
	firstServiceName := getFirstServiceConfigName(project.Services)

	// write ini file to temporary buffer
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, project.Services[firstServiceName]); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func getFirstServiceConfigName(m map[string]types.ServiceConfig) string {
    for s := range m {
        return s
    }
    return ""
}

// func CreateTemplateMap(project *types.Project) (map[string]string, error) {
// 	tmplMap := make(map[string]string)
// }