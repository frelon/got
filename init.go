package main

import (
	"io"
	"os"
	"fmt"
	"errors"

	"path/filepath"

	"github.com/spf13/cobra"
)

type initCmd  struct {
	out     io.Writer
	path 	string
}

func newInitCmd(out io.Writer) *cobra.Command {
	i := &initCmd{out: out}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "initialise a git repository",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			i.path = "."
			if len(args) != 0 {
				i.path = args[0]
			}

			return i.run()
		},
	}

	return cmd
}

func (i *initCmd) run() error {
	_, err := os.Stat(i.path)
    if err != nil {
		return errors.New(fmt.Sprintf("%+v does not exist", i.path))
	}

	gitPath := filepath.Join(i.path, ".git")
	_, err = os.Stat(gitPath)
	if err == nil {
		return errors.New(fmt.Sprintf("%+v is already a repository", i.path))
	}

	err = os.Mkdir(gitPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}