package main

import (
	"github.com/spf13/cobra"
)

func newRootCmd(args []string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "got",
		Short: "Got is a git-client written in go",
		Long: `Testing`,
		Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		},
	}

	out := cmd.OutOrStdout()

	cmd.AddCommand(
		newInitCmd(out),
		newAddCmd(out),
	)

	return cmd
}