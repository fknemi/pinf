package pinf

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pinf",
		Short: "...",
		Long:  "...",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "...")
		},
	}

	return cmd
}
