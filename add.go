package main

import (
	"io"
	"fmt"
	"errors"

	"github.com/spf13/cobra"
)

type addCmd struct {
	out		io.Writer
	paths	[]string
}

func newAddCmd(out io.Writer) *cobra.Command {
	i := &addCmd{out: out}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "stage changes",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("Must specify something to add!")
			}

			i.paths = args[:]

			return i.run()
		},
	}

	return cmd
}

func (i *addCmd) run() error {
	fmt.Printf("add paths: %+v", i.paths)
	return nil
}