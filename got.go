package main

import (
	"os"
	"fmt"
)

func main() {
	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}