package cli

import (
	"flag"
)

var (
	ComposeFilePath string
	IniFilePath     string
)

func ArgsInit() {
	flag.StringVar(&ComposeFilePath, "c", "docker-compose.yml", "Path to Compose file")
	flag.StringVar(&IniFilePath, "i", "result.service", "Path to resulting service file")
	flag.Parse()
}
