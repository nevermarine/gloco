package tmpl

import (
	"embed"
)

//go:embed systemd.ini
var TmplFile embed.FS

const TmplFilePath = "systemd.ini"
