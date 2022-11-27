package main

import (
	"fmt"
	"os"

	"github.com/abhisek/sane/internal/sane"
	"github.com/spf13/cobra"
)

type validateCommand struct {
	repoPath  string
	rulesPath string
}

func newValidateCommand() *cobra.Command {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	validate := &validateCommand{}
	cmd := &cobra.Command{
		Use: "validate [OPTIONS]",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := validate.run(args); err != nil {
				sane.LoggerWithError(err).Errorf("Validation failed")
				os.Exit(1)
			}

			return err
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&validate.repoPath, "path", "p", currDir,
		"Git repository path (Default: $PWD)")
	fs.StringVarP(&validate.rulesPath, "rules", "r",
		fmt.Sprintf("%s/%s", currDir, ".sane"),
		"Sane rules path (Default: $PWD/.sane)")

	return cmd
}

func (v *validateCommand) run(args []string) error {
	sane.SetLogLevel(verbose, debug)
	return sane.Execute(sane.Config{
		RepositoryPath: v.repoPath,
		RulesPath:      v.rulesPath,
		RulesType:      sane.RULES_FORMAT_GITIGNORE,
	})
}
