package parse

import (
	"embed"
	"github.com/nevermarine/gloco/errdefs"
	"github.com/nevermarine/gloco/tmpl"
	"io/fs"
	"os"
	"testing"
)

const (
	testEmptyFilePath    = "testfiles/empty.file"
	testMultiSvcFilePath = "testfiles/multi-svc.file"
)

//go:embed testfiles
var testfiles embed.FS

func TestMultiSvc(t *testing.T) {
	content, err := fs.ReadFile(testfiles, testMultiSvcFilePath)
	if err != nil {
		t.Fatalf("failed to read embedded test file: %v", err)
	}

	// create a temporary file to hold the content
	tmpFile, err := os.CreateTemp("", "multi-svc-*.file")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		tmpFile.Close()
		t.Fatalf("failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	project, err := LoadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load project from temp file: %v", err)
	}

	tmplFile, err := fs.ReadFile(tmpl.TmplFile, tmpl.TmplFilePath)
	if err != nil {
		t.Fatalf("failed to read template file: %v", err)
	}

	_, err = CreateIni(project, tmplFile)
	if !errdefs.ContainsMultipleServices(err) {
		t.Errorf("expected multiple services error, got %v", err)
	}
}
