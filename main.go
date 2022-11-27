package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	debug   bool
)

func main() {
	cmd := &cobra.Command{
		Use:              "sane [OPTIONS] COMMAND [ARG...]",
		Short:            "Git repository structure validator",
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			return fmt.Errorf("sane: %s is not a valid command", args[0])
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose logs")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Show debug logs")

	cmd.AddCommand(newValidateCommand())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
