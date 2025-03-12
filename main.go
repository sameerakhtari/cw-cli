package main

import (
	"github.com/sameerakhtari/cw-cli/cmd"
	"github.com/sameerakhtari/cw-cli/pkg/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()
	cmd.Execute()
}
