package prepareDB

import (
	"fmt"
	"github.com/spf13/cobra"
	"postgres-reforger/internal"
)

// Cmd returns the "serve" command of the application.
// This command is responsible for initializing and
func Cmd(app *internal.App) *cobra.Command {
	return &cobra.Command{
		Use:   "prepareDB",
		Short: "preparations For DataBase",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := app.Init(); err != nil {
				return fmt.Errorf("application initialisation: %w", err)
			}
			if err := app.PrepareDb(); err != nil {
				return fmt.Errorf("database preparation has failed because of: %w", err)
			}
			return nil
		},
	}
}
