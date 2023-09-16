package command

import (
	"github.com/spf13/cobra"
)

// NewRootCommand constructs a new base command for the terminal.
//
//go:generate sh -c "printf %s $(git describe --tags --abbrev=0) > version.txt"
func NewRootCommand() *cobra.Command {
	var rootCommand = &cobra.Command{
		Use: "kubernetes",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCommand.AddCommand(
		newVersionCommand(),
	)
	return rootCommand
}
