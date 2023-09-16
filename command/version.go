package command

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	//go:embed version.txt
	version string
)

// newVersionCommand prints the version information to the console.
func newVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:          "version",
		Short:        "Print the version",
		Example:      `  deus version`,
		SilenceUsage: false,
	}

	command.Run = func(cmd *cobra.Command, args []string) {
		if len(version) == 0 {
			fmt.Println("dev")
		} else {
			fmt.Println(version)
		}
	}

	return command
}
