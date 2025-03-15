package main

import (
	"os"
	"postgres-reforger/cmd/prepareDB"
	"postgres-reforger/cmd/root"
	"postgres-reforger/cmd/serve"
	"postgres-reforger/internal"

	"github.com/misnaged/scriptorium/logger"
)

func main() {
	app, err := internal.NewApplication()
	if err != nil {
		logger.Log().Errorf("An error occurred %#v", err)
		os.Exit(1)
	}

	rootCmd := root.Cmd(app)
	rootCmd.AddCommand(serve.Cmd(app))
	rootCmd.AddCommand(prepareDB.Cmd(app))

	if err = rootCmd.Execute(); err != nil {
		logger.Log().Errorf("An error occurred %#v", err)
		os.Exit(1)
	}
}
